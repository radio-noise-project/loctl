package run

import (
	"github.com/radio-noise-project/loctl/cli"
	"github.com/radio-noise-project/loctl/pkg/run"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run code on Sisters",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		run.StartContainer(args[0], args[1])
	},
}

func init() {
	cli.RootCmd.AddCommand(runCmd)
}
