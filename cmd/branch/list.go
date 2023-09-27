package branch

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List branches",
	Long:  "List all branches within the active project.",
	Run: func(cmd *cobra.Command, args []string) {
		// Your code to list branches here
		fmt.Println("Listing branches...")
	},
}

func init() {
	// Add any flags if needed for the 'list' command
}
