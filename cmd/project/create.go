package project

import (
	"fmt"
	"log"

	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [projectName]",
	Short: "Create a new project",
	Long:  "Create a new project with the specified name.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db := app.GetDB()
		if db == nil {
			log.Fatalf("Failed to load the database.")
		}

		projectName := args[0]
		projectDescription, _ := cmd.Flags().GetString("description")

		// Create a new project
		project := models.Project{
			Name:        projectName,
			Description: projectDescription,
		}

		if err := db.Create(&project).Error; err != nil {
			log.Fatalf("Error creating project: %v", err)
		}

		fmt.Println("New project added successfully.")
	},
}

func init() {
	// Add flags or options here if needed
	createCmd.Flags().StringP("description", "d", "", "Project description (optional)")
}
