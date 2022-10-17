package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addonsCmd = &cobra.Command{
	Use:   "addons",
	Short: "Check current add-ons version and which can be upgraded",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addons called")
	},
}

func init() {
	checkCmd.AddCommand(addonsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addonsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addonsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
