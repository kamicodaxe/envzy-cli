package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Projects operations",
	Long:  "Manage projects.",
	Run: func(cmd *cobra.Command, args []string) {
		// Your code for the 'project' command here
		fmt.Println("Project operations...")
	},
}

func init() {
	// Add any flags if needed for the 'project' command

	// Add subcommands to the 'project' command
	projectsCmd.AddCommand(createCmd)
	projectsCmd.AddCommand(listCmd)
	// projectsCmd.AddCommand(selectCmd)
	// projectsCmd.AddCommand(updateCmd)
	// projectsCmd.AddCommand(deleteCmd)
}

func AddToRoot(rootCmd *cobra.Command) {
	rootCmd.AddCommand(projectsCmd)
}
