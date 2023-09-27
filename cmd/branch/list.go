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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List branches",
	Long:  "List all branches within the active branch.",
	Run: func(cmd *cobra.Command, args []string) {
		db := app.GetDB()
		if db == nil {
			log.Fatalf("Failed to load the database.")
		}

		selectedProject := api.GetCurrentProject()
		if selectedProject == nil {
			log.Fatalf("Error fetching selected project")
			return
		}

		var branchNames []string
		if err := db.Model(&models.Branch{}).Where("project_id = ?", selectedProject.ID).Pluck("name", &branchNames).Error; err != nil {
			log.Fatalf("Error fetching branch names: %v", err)
			project.SelectProject()
			return
		}

		fmt.Println("Project: ", selectedProject.Name)
		fmt.Println("List of branches you have access to:")
		selectedBranch := api.GetCurrentBranch()
		for _, branchName := range branchNames {
			if selectedBranch != nil && (branchName == selectedBranch.Name) {
				fmt.Println("* " + branchName)
				continue
			}
			fmt.Println("- " + branchName)
		}
	},
}

func init() {
	// Add any flags if needed for the 'list' command
}
