package project

import (
	"fmt"
	"log"

	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a project from tthe list of projects",
	Run: func(cmd *cobra.Command, args []string) {

		if args[0] == "" {
			log.Println("Please provide the pro")
		}

		db := app.GetDB()
		if db == nil {
			log.Fatalf("Failed to connect to the database.")
		}

		// var projectID uint
		// if err:= helpers.GetProjects()
		var projectNames []string
		if err := db.Model(&models.Project{}).Pluck("name", &projectNames).Error; err != nil {
			log.Fatalf("Error fetching project names: %v", err)
		}

		fmt.Println("List of projects you have access to:")
		for _, projectName := range projectNames {
			fmt.Println("- " + projectName)
		}
	},
}
