package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var commitMessage string

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes to environment files with a message, establishing version control",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Committing changes with message: %s\n", commitMessage)
		// Code here
		fmt.Println("Changes committed successfully.")
	},
}

func init() {
	commitCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "Commit message")
	commitCmd.MarkFlagRequired("message")
}
