package secret

import (
	"fmt"

	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [secretName]",
	Aliases: []string{"d"},
	Short:   "Delete a secret",
	Args:    cobra.ExactArgs(1), // Expect one argument (secretName)
	Run: func(cmd *cobra.Command, args []string) {
		secretName := args[0]

		// Get the active project and branch
		activeProject := api.GetCurrentProject()
		activeBranch := api.GetCurrentBranch()

		if activeProject == nil || activeBranch == nil {
			fmt.Println("No active project or branch selected.")
			return
		}

		// Call your API endpoint to delete the secret
		err := api.DeleteSecretByName(secretName, activeProject.ID, activeBranch.ID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Secret \"%s\" deleted successfully for Project \"%s\", Branch \"%s\"!\n", secretName, activeProject.Name, activeBranch.Name)
	},
}

func init() {
	// Add flags or options here if needed
}
