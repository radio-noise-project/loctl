package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "loctl",
	Short: "A self-sufficient runtime for Last-Order",
	Long: `The control tool for Radio-Noise-Project.
This command is connected Last-Order in your registered computer.
And the command is provided to operate Sisters nodes via Last-Order.
When you operate these, you have to use the command.`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
