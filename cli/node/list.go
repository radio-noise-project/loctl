package node

import (
	"fmt"

	"github.com/radio-noise-project/loctl/pkg/node"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all nodes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, address, port := node.ListNode()
		fmt.Printf("Name: %s\nAddress: %s\nPort: %s\n", name, address, port)
	},
}

func init() {
	nodeCmd.AddCommand(listCmd)
}
