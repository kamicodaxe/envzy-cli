package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch [project name]",
	Short: "Switch to a different project",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		fmt.Printf("Switching to project: %s\n", projectName)
	},
}

func init() {
	switchCmd.Flags().StringP("branch", "b", "", "Select a branch (optional)")
}
