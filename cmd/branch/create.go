package branch

import (
	"fmt"
	"log"

	"github.com/kamicodaxe/envzy-cli/cmd/project"
	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [branchName]",
	Short: "Create a new branch",
	Long:  "Create a new branch with the specified name.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db := app.GetDB()
		if db == nil {
			log.Fatalf("Failed to load the database.")
		}

		branchName := args[0]
		branchDescription, _ := cmd.Flags().GetString("description")

		currentProject := api.GetCurrentProject()
		if currentProject == nil {
			project.SelectProject()
			return
		}

		// Create a new branch
		branch := models.Branch{
			Name:        branchName,
			ProjectID:   uint(currentProject.ID),
			Description: branchDescription,
		}

		if err := db.Create(&branch).Error; err != nil {
			log.Fatalf("Error creating branch: %v", err)
		}

		fmt.Println("New branch added successfully.")
	},
}

func Create() {

}

func init() {
	// Add flags or options here if needed
	createCmd.Flags().StringP("description", "d", "", "Project description (optional)")
}
