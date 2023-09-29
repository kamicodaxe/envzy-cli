package secret

import (
	"fmt"

	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create [secretName] [secretValue]",
	Aliases: []string{"c"},
	Short:   "Create a new secret",
	Long:    "Create a new secret with the specified name.",
	Args:    cobra.ExactArgs(2), // Expect two arguments (secretName and secretValue)
	Run: func(cmd *cobra.Command, args []string) {
		secretName := args[0]
		secretValue := args[1]

		// Get the active project and branch
		activeProject := api.GetCurrentProject()
		activeBranch := api.GetCurrentBranch()

		if activeProject == nil || activeBranch == nil {
			fmt.Println("No active project or branch selected.")
			return
		}

		// Create a models.Secret instance
		secret := models.Secret{
			Name:    secretName,
			Value:   secretValue,
			Project: *activeProject,
			Branch:  *activeBranch,
		}

		// Implement the logic to create a secret for the active project and branch
		err := api.CreateSecret(&secret)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Secret \"%s\" created successfully for Project \"%s\", Branch \"%s\"!\n", secretName, activeProject.Name, activeBranch.Name)
	},
}

func init() {
	// Add flags or options here if needed
}
