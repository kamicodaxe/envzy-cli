package secret

import (
	"fmt"

	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List secrets for the active project and branch",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the active project and branch
		activeProject := api.GetCurrentProject()
		activeBranch := api.GetCurrentBranch()

		if activeProject == nil || activeBranch == nil {
			fmt.Println("No active project or branch selected.")
			return
		}

		// Call your API endpoint to get secrets for the active project and branch
		secrets, err := api.GetSecretsByProjectAndBranch(activeProject.ID, activeBranch.ID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Display the list of secrets
		fmt.Println("Secrets for Project:", activeProject.Name, "Branch:", activeBranch.Name)
		for _, secret := range secrets {
			fmt.Println("- Name:", secret.Name, "Value:", secret.Value)
		}
	},
}

func init() {
	// Add flags or options here if needed
}
