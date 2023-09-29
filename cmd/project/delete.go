package project

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Deletes a project and its associated secrets",
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve the list of available projects from your database
		projects, err := api.GetProjects()
		if err != nil {
			log.Fatalf("An error occurred: %v", err)
		}

		if len(projects) == 0 {
			fmt.Println("No projects available to delete.")
			return
		}

		// Define a slice of project names for the survey prompt
		projectNames := make([]string, len(projects))
		for i, project := range projects {
			projectNames[i] = project.Name
		}

		prompt := &survey.Select{
			Message: "Select a project to delete:",
			Options: projectNames,
		}

		var selectedProjectName string
		survey.AskOne(prompt, &selectedProjectName)

		// Confirm the deletion with the user
		confirm := false
		confirmPrompt := &survey.Confirm{
			Message: fmt.Sprintf("Are you sure you want to delete project '%s' and its associated secrets?", selectedProjectName),
			Default: false,
		}
		survey.AskOne(confirmPrompt, &confirm)

		if !confirm {
			fmt.Println("Deletion canceled.")
			return
		}

		// Find the selected project by name
		var selectedProject *models.Project
		for _, project := range projects {
			if project.Name == selectedProjectName {
				selectedProject = &project
				break
			}
		}

		if selectedProject == nil {
			fmt.Println("Selected project not found.")
			return
		}

		// Delete the project and its associated secrets from the database
		err = api.DeleteProjectByID(selectedProject.ID)
		if err != nil {
			fmt.Printf("Failed to delete project '%s': %v\n", selectedProjectName, err)
		} else {
			fmt.Printf("Project '%s' and its associated secrets have been deleted.\n", selectedProjectName)
		}
	},
}

func init() {
	// Add flags or additional configuration for the "delete" command here, if needed
}
