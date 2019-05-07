package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/seamounts/essh/pkg/config"
	"github.com/seamounts/essh/pkg/sshcli"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var ESSHCmd = &cobra.Command{
	Use:   "essh",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func Execute() {
	if err := ESSHCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var (
	templates = &promptui.SelectTemplates{
		Label:    "✨ {{ . | green}}",
		Active:   "➤ {{ .Name | cyan  }}{{if .Alias}}({{.Alias | yellow}}){{end}} {{if .Host}}{{if .User}}{{.User | faint}}{{`@` | faint}}{{end}}{{.Host | faint}}{{end}}",
		Inactive: "  {{.Name | faint}}{{if .Alias}}({{.Alias | faint}}){{end}} {{if .Host}}{{if .User}}{{.User | faint}}{{`@` | faint}}{{end}}{{.Host | faint}}{{end}}",
	}
)

func run() {

	err := config.LoadConfig([]string{".essh.yaml", "essh.yaml"})
	if err != nil {
		log.Printf("load config error", err)
		os.Exit(1)
	}
	nodes := config.GetConfig()

	node := choose(nodes)
	if node == nil {
		return
	}
	client, err := sshcli.NewClient(node)
	if err != nil {
		log.Fatal(err)
	}
	client.Login()

}

func choose(nodes []*config.Node) *config.Node {
	prompt := promptui.Select{
		Label:     "select host",
		Items:     nodes,
		Templates: templates,
		Size:      20,
		Searcher: func(input string, index int) bool {
			node := nodes[index]
			content := fmt.Sprintf("%s %s %s", node.Name, node.User, node.Host)
			if strings.Contains(input, " ") {
				for _, key := range strings.Split(input, " ") {
					key = strings.TrimSpace(key)
					if key != "" {
						if !strings.Contains(content, key) {
							return false
						}
					}
				}
				return true
			}
			if strings.Contains(content, input) {
				return true
			}
			return false
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return nil
	}

	node := nodes[index]
	return node
}
