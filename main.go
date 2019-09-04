package main

import (
	"github.com/seamounts/essh/cmd"
)

var (
	version string
)

func main() {
	cmd.Execute(version)
}
