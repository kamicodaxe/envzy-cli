package team

import (
	"fmt"

	"github.com/spf13/cobra"
)

var inviteCmd = &cobra.Command{
	Use:   "add [user@email.com]",
	Short: "Invite collaborators to your projects, supporting teamwork",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		userEmail := args[0]
		fmt.Printf("Inviting user: %s\n", userEmail)
		// Code here
		fmt.Println("Invitation sent successfully.")
	},
}

func init() {
	// Add flags or additional configuration for the "team add" command here, if needed
}
