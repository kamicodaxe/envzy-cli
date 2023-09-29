package secret

import (
	"fmt"

	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update [secretName] [newSecretValue]",
	Aliases: []string{"u"},
	Short:   "Update a secret's value",
	Args:    cobra.ExactArgs(2), // Expect two arguments (secretName and newSecretValue)
	Run: func(cmd *cobra.Command, args []string) {
		secretName := args[0]
		newSecretValue := args[1]

		// Get the active project and branch
		activeProject := api.GetCurrentProject()
		activeBranch := api.GetCurrentBranch()

		if activeProject == nil || activeBranch == nil {
			fmt.Println("No active project or branch selected.")
			return
		}

		// Call your API endpoint to update the secret's value
		err := api.UpdateSecretByName(secretName, newSecretValue, activeProject.ID, activeBranch.ID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Secret \"%s\" updated successfully for Project \"%s\", Branch \"%s\"!\n", secretName, activeProject.Name, activeBranch.Name)
	},
}

func init() {
	// Add flags or options here if needed
}
