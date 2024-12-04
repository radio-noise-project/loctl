package run

import (
	"github.com/radio-noise-project/loctl/cli"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run code on Sisters",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cli.RootCmd.AddCommand(runCmd)
}
