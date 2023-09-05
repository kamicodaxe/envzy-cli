package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch [branch name]",
	Short: "Work with different branches or configurations within the selected project",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]
		fmt.Printf("Working with branch: %s\n", branchName)
		// Implement your logic here to work with the specified branch
		fmt.Println("Branch selected successfully.")
	},
}
