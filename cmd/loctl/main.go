package main

import (
	"github.com/radio-noise-project/loctl/cli"
	_ "github.com/radio-noise-project/loctl/cli/config"
	_ "github.com/radio-noise-project/loctl/cli/lb"
	_ "github.com/radio-noise-project/loctl/cli/node"
	_ "github.com/radio-noise-project/loctl/cli/run"
)

func main() {
	cli.Execute()
}
