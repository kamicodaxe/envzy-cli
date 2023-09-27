package branch

import (
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Branch operations",
	Long:  "Manage branch.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Run(listCmd, nil)
	},
}

func init() {
	// Add any flags if needed for the 'branch' command

	// Add subcommands to the 'branch' command
	branchCmd.AddCommand(createCmd)
	branchCmd.AddCommand(listCmd)
	branchCmd.AddCommand(selectCmd)
	// branchCmd.AddCommand(updateCmd)
	// branchCmd.AddCommand(deleteCmd)
}

func AddToRoot(rootCmd *cobra.Command) {
	rootCmd.AddCommand(branchCmd)
}
