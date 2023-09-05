package cmd

import (
	"fmt"
	"log"

	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Manage projects",
	Long:  "Commands to manage projects.",
}

var projectsAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new project",
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

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the projects you have access to, providing an overview of available projects",
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
		for _, projectName := range projectNames {
			fmt.Println("- " + projectName)
		}
	},
}

var projectsRemoveCmd = &cobra.Command{
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

func init() {
	projectsCmd.AddCommand(projectsAddCmd)
	projectsCmd.AddCommand(projectsListCmd)
	projectsAddCmd.Flags().StringP("description", "d", "", "Project description (optional)")
}
