package project

import (
	"fmt"
	"log"

	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all projects",
	Long:    "List the projects you have access to, providing an overview of available projects",
	Run: func(cmd *cobra.Command, args []string) {
		db := app.GetDB()
		if db == nil {
			log.Fatalf("Failed to connect to the database.")
		}

		var projectNames []string
		if err := db.Model(&models.Project{}).Pluck("name", &projectNames).Error; err != nil {
			log.Fatalf("Error fetching project names: %v", err)
		}

		fmt.Println("List of projects you have access to:")
		selectedProject := api.GetCurrentProject()
		for _, projectName := range projectNames {
			if selectedProject != nil && (projectName == selectedProject.Name) {
				fmt.Println("* " + projectName)
				continue
			}
			fmt.Println("- " + projectName)
		}
	},
}
