package command

import (
	"fmt"

	"github.com/radio-noise-project/loctl/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Radio-Noise-Project",
	Long: `Create new Radio-Noise-Project.
This command is registared node of Last-Order to thsi loctl environment.
When you execute this command, command will create config files in ~/.config/ directory.
Config file is written by toml and recorded any settings in Radio-Noise-Project environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		var addr string
		var port int
		fmt.Println("Please enter the IP address of Last-Order")
		fmt.Scan(&addr)
		fmt.Println("Please enter the port of Last-Order")
		fmt.Scan(&port)
		config.InitializeConfiguration(addr, port)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	RootCmd.Flags()
}
