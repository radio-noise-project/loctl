package node

import (
	"fmt"

	"github.com/radio-noise-project/loctl/pkg/node"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new node to Last-order",
	Long:  ``,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		status := node.AddNode(args[0], args[1], args[2])
		if status != 200 {
			fmt.Println("Error")
		} else {
			fmt.Println("Node added")
		}
	},
}

func init() {
	nodeCmd.AddCommand(addCmd)
}
