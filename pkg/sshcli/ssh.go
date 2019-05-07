package sshcli

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"syscall"
	"time"

	"github.com/seamounts/essh/pkg/config"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	DefaultCiphers = []string{
		"aes128-ctr",
		"aes192-ctr",
		"aes256-ctr",
		"aes128-gcm@openssh.com",
		"chacha20-poly1305@openssh.com",
		"arcfour256",
		"arcfour128",
		"arcfour",
		"aes128-cbc",
		"3des-cbc",
		"blowfish-cbc",
		"cast128-cbc",
		"aes192-cbc",
		"aes256-cbc",
	}
)

type defaultClient struct {
	ClientConfig *ssh.ClientConfig
	Node         *config.Node
}

func NewClient(node *config.Node) (*defaultClient, error) {
	auth := make([]ssh.AuthMethod, 0)

	if node.KeyPath != "" {
		keybyte, err := ioutil.ReadFile(node.KeyPath)
		if err != nil {
			return nil, err
		}
		signer, err := ssh.ParsePrivateKey(keybyte)
		if err != nil {
			return nil, err
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}
	if node.Password != "" {
		auth = append(auth, ssh.Password(node.Password))
	}

	if len(auth) == 0 && node.NeedAuth {
		fmt.Printf("password:")
		b, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil, err
		}
		auth = append(auth, ssh.Password(string(b)))
	}

	sshConfig := &ssh.ClientConfig{
		User:            node.User,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 10,
	}
	sshConfig.SetDefaults()
	sshConfig.Ciphers = append(sshConfig.Ciphers, DefaultCiphers...)

	client := &defaultClient{
		ClientConfig: sshConfig,
		Node:         node,
	}

	return client, nil
}

func (c *defaultClient) Login() {

	var client *ssh.Client
	var err error

	host := c.Node.Host
	port := c.Node.Port

	if len(c.Node.Jump) > 0 {
		jclient, err := NewClient(c.Node.Jump[0])
		if err != nil {
			log.Fatal(err)
		}
		proxyClient, err := ssh.Dial("tcp", net.JoinHostPort(host, port), c.ClientConfig)
		if err != nil {
			log.Fatal(err)
		}
		conn, err := proxyClient.Dial("tcp", net.JoinHostPort(jclient.Node.Host, jclient.Node.Port))
		if err != nil {
			log.Fatal(err)
		}
		proxyconn, chans, reqs, err := ssh.NewClientConn(conn, net.JoinHostPort(jclient.Node.Host, jclient.Node.Port), jclient.ClientConfig)
		if err != nil {
			log.Fatal(err)
		}
		client = ssh.NewClient(proxyconn, chans, reqs)
	} else {
		client, err = ssh.Dial("tcp", net.JoinHostPort(host, port), c.ClientConfig)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		log.Fatal(err)
	}
	defer terminal.Restore(fd, state)

	w, h, err := terminal.GetSize(fd)
	if err != nil {
		log.Fatal(err)
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("xterm", h, w, modes)
	if err != nil {
		log.Fatal(err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	stdinPipe, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = session.Shell()
	if err != nil {
		log.Fatal(err)
	}

	// then exec cmd
	for i := range c.Node.Cmds {
		shellCmd := c.Node.Cmds[i]
		time.Sleep(shellCmd.Delay * time.Millisecond)
		stdinPipe.Write([]byte(shellCmd.Cmd + "\r"))
	}

	// change stdin to user
	go func() {
		_, err = io.Copy(stdinPipe, os.Stdin)

		log.Fatal(err)
		session.Close()
	}()

	go func() {
		var (
			ow = w
			oh = h
		)
		for {
			cw, ch, err := terminal.GetSize(fd)
			if err != nil {
				break
			}

			if cw != ow || ch != oh {
				err = session.WindowChange(ch, cw)
				if err != nil {
					break
				}
				ow = cw
				oh = ch
			}
			time.Sleep(time.Second)
		}
	}()

	// send keepalive
	go func() {
		for {
			time.Sleep(time.Second * 10)
			client.SendRequest("keepalive@openssh.com", false, nil)
		}
	}()

	err = session.Wait()
	if err != nil {
		log.Fatal(err)
	}

}
