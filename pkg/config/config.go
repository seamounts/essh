package config

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path"
	"time"

	"gopkg.in/yaml.v2"
)

type Node struct {
	Name     string      `yaml:"name"`
	Alias    string      `yaml:"alias"`
	Host     string      `yaml:"host"`
	User     string      `yaml:"user"`
	Port     string      `yaml:"port"`
	KeyPath  string      `yaml:"keypath"`
	Password string      `yaml:"password"`
	Jump     []*Node     `yaml:"jump"`
	NeedAuth bool        `yaml:"needauth"`
	Cmds     []*ShellCmd `yaml:"cmds`
}

type ShellCmd struct {
	Cmd   string        `yaml:"cmd"`
	Delay time.Duration `yaml:"delay"`
}

func (n *Node) String() string {
	return fmt.Sprintf("%s@%s:%s", n.User, n.Host, n.Port)
}

var (
	configs []*Node
)

func GetConfig() []*Node {
	return configs
}

func LoadConfig(names []string) error {
	b, err := LoadConfigBytes(names...)
	if err != nil {
		return err
	}
	nodes := []*Node{}
	err = yaml.Unmarshal(b, &nodes)
	if err != nil {
		return err
	}
	for i, _ := range nodes {
		nodes[i].NeedAuth = true
		if nodes[i].Port == "" {
			nodes[i].Port = "22"
		}
	}
	configs = nodes

	return nil
}

func LoadConfigBytes(names ...string) ([]byte, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	// homedir
	for i := range names {
		sshw, err := ioutil.ReadFile(path.Join(u.HomeDir, names[i]))
		if err == nil {
			return sshw, nil
		}
	}
	// relative
	for i := range names {
		sshw, err := ioutil.ReadFile(names[i])
		if err == nil {
			return sshw, nil
		}
	}
	return nil, err
}
