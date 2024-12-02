package cli

import (
	"fmt"

	"github.com/radio-noise-project/loctl/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the RNP version information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		loctl, _ := cmd.Flags().GetBool("loctl")
		lastOrder, _ := cmd.Flags().GetBool("last-order")
		sistersNum, _ := cmd.Flags().GetInt("sisters")

		if loctl {
			version, codeName, builtDate := version.GetVersionLoctl()
			fmt.Print("Control Application: loctl\n")
			fmt.Printf("Version: %s\n", version)
			fmt.Printf("CodeName: %s\n", codeName)
			fmt.Printf("Built Date: %s\n", builtDate)
		} else if lastOrder {
		} else if sistersNum > 0 {
			fmt.Printf("Sisters version %d\n", sistersNum)
		}
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolP("loctl", "", false, "Show loctl version")
	versionCmd.Flags().BoolP("last-order", "", false, "Show last-order version")
	versionCmd.Flags().IntP("sisters", "", 0, "Show sisters version")
}
