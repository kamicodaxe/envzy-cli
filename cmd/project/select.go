package project

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/constants"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select [project name]",
	Short: "Select a project",
	Args:  cobra.MaximumNArgs(1), // Requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve the list of available projects from your database
		var selectedProject models.Project
		var selectedProjectName string
		var branches []models.Branch
		shouldChooseBranch := false

		projects, err := api.GetProjects()

		if err != nil {
			fmt.Println("An error occured.")
		}

		if len(projects) == 0 {
			fmt.Println("No projects available.")
			return
		}

		// Define a slice of project names for the survey prompt
		projectNames := make([]string, len(projects))
		for i, project := range projects {
			projectNames[i] = project.Name
		}

		prompt := &survey.Select{
			Message: "Select a project:",
			Options: projectNames,
		}

		// Prompt the user to select a project
		survey.AskOne(prompt, &selectedProject.Name)

		for _, project := range projects {
			if project.Name == selectedProjectName {
				selectedProject = project
			}
		}

		// Update the application's state with the selected project
		kvstore := app.GetKVStore()

		// TODO: Handle errors
		kvstore.SetString(constants.PROJECT_NAME, selectedProject.Name)
		kvstore.SetUInt(constants.PROJECT_ID, selectedProject.ID)

		name := kvstore.String(constants.PROJECT_NAME)
		fmt.Printf("Project selected: %s\n", name)

		// Create a question to ask the user
		branchQuestion := &survey.Confirm{
			Message: "Do you want to choose a branch now?",
			Default: false, // You can set the default answer here
		}

		// Ask the user the question
		err = survey.AskOne(branchQuestion, &shouldChooseBranch)
		if err != nil {
			fmt.Println("Error asking the question:", err)
			return
		}

		// Check the user's answer
		if shouldChooseBranch {
			// User wants to choose a branch
			fmt.Println("You chose to select a branch.")
			branches, err = api.GetBranchesByProjectID(selectedProject.ID)
			if err != nil {
				fmt.Println("An error occured retrieving branch!")
			}

			branchNames := make([]string, len(branches))
			for i, branch := range branches {
				branchNames[i] = branch.Name
			}

			// Add your code here to handle branch selection
		} else {
			// User doesn't want to choose a branch
			fmt.Println("You chose not to select a branch.")
			// Add your code here for the other scenario
		}

	},
}

func SelectBranch() {

}

func init() {
	selectCmd.Flags().StringP("branch", "b", "", "Select a branch (optional)")
}
