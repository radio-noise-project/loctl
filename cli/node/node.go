package node

import (
	"github.com/radio-noise-project/loctl/cli"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Operate Sisters nodes",
	Long:  ``,
}

func init() {
	cli.RootCmd.AddCommand(nodeCmd)
}
