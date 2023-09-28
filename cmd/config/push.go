package config

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/constants"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/kamicodaxe/envzy-cli/utils/dotenv"
	"github.com/spf13/cobra"
)

var pushProject string
var pushBranch string

var pushCmd = &cobra.Command{
	Use:   "push [path/to/dotenv]",
	Short: "Pushes all env files variables in env-file to envy",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		var selectedProject *models.Project
		var selectedBranch *models.Branch
		envFile := args[0]
		fmt.Printf("Adding environment file: %s\n", envFile)

		if pushProject == "" {
			activeProject := api.GetCurrentProject()

			if activeProject != nil {
				pushProject = activeProject.Name
				selectedProject = activeProject
			} else {
				selectedProject = SelectProject()
				if selectedProject != nil {
					pushProject = selectedProject.Name
				} else {
					fmt.Println("No Project selected")
					return
				}
			}
		}

		if pushBranch == "" {
			activeBranch := api.GetCurrentBranch()

			if activeBranch != nil {
				pushBranch = activeBranch.Name
				selectedBranch = activeBranch
			} else {
				selectedBranch = SelectBranch(selectedProject.ID)
				if selectedBranch != nil {
					pushBranch = selectedBranch.Name
				} else {
					fmt.Println("No Branch selected")
					return
				}
			}

		}

		fmt.Printf("Project: %s\n", pushProject)
		fmt.Printf("Branch: %s\n", pushBranch)

		// Add dotenv file to secrets
		err := dotenv.AddDotenvFileToSecrets(envFile, *selectedProject, *selectedBranch)

		if err != nil {
			fmt.Println("Error:", err)
		}

	},
}

// SelectProject prompts the user to select a project from available options.
// It returns the selected project object.
func SelectProject() *models.Project {
	// Retrieve the list of available projects from your database
	var selectedProject *models.Project
	var selectedProjectName string

	projects, err := api.GetProjects()
	if err != nil {
		fmt.Println("An error occurred.")
		return nil
	}

	if len(projects) == 0 {
		fmt.Println("No projects available.")
		return nil
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
	survey.AskOne(prompt, &selectedProjectName)

	for _, project := range projects {
		if project.Name == selectedProjectName {
			selectedProject = &project
		}
	}

	// Update the application's state with the selected project
	kvstore := app.GetKVStore()

	// TODO: Handle errors
	kvstore.SetString(constants.PROJECT_NAME, selectedProject.Name)
	kvstore.SetUInt(constants.PROJECT_ID, selectedProject.ID)

	name := kvstore.String(constants.PROJECT_NAME)
	fmt.Printf("Project selected: %s\n", name)

	return selectedProject
}

// SelectBranch prompts the user to select a branch from available options for a given project.
// It returns the selected branch object.
func SelectBranch(projectID uint) *models.Branch {
	// Retrieve the list of branches for the specified project from the API
	branches, err := api.GetBranchesByProjectID(projectID)
	if err != nil {
		fmt.Println("An error occurred while retrieving branches.")
		return nil
	}

	if len(branches) == 0 {
		fmt.Println("No branches available for the selected project.")
		return nil
	}

	// Define a slice of branch names for the survey prompt
	branchNames := make([]string, len(branches))
	for i, branch := range branches {
		branchNames[i] = branch.Name
	}

	prompt := &survey.Select{
		Message: "Select a branch:",
		Options: branchNames,
	}

	var selectedBranchName string
	survey.AskOne(prompt, &selectedBranchName)

	for _, branch := range branches {
		if branch.Name == selectedBranchName {
			return &branch
		}
	}

	// In case no branch is found (this should not happen)
	return nil
}

func init() {
	// Add flags or additional configuration for the "push" command here, if needed
	// TODO: Uncomment the flags bellow and use it to enable the flexibility to specify a project and branch
	// pushCmd.PersistentFlags().StringVarP(&pushProject, "project", "p", "", "Project name")
	// pushCmd.PersistentFlags().StringVarP(&pushBranch, "branch", "b", "", "Branch name")
}
