package branch

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kamicodaxe/envzy-cli/cmd/project"
	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/kamicodaxe/envzy-cli/internal/app"
	"github.com/kamicodaxe/envzy-cli/internal/constants"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a branch",
	Args:  cobra.MaximumNArgs(1), // Requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		SelectBranch()
	},
}

func SelectBranch() {
	var selectedBranch models.Branch

	selectedProject := api.GetCurrentProject()
	if selectedProject == nil {
		fmt.Println("Select a project first")
		project.SelectProject()
	}

	branches, err := api.GetBranchesByProjectID(selectedProject.ID)
	if err != nil {
		fmt.Println("An error occured retrieving branch!")
		panic(err)
	}

	if len(branches) == 0 {
		fmt.Println("Please add a branch first: envzy branch add [Branch Name]")
		return
	}

	branchNames := make([]string, len(branches))
	for i, branch := range branches {
		branchNames[i] = branch.Name
	}

	prompt := &survey.Select{
		Message: "Select a branch:",
		Options: branchNames,
	}

	// Prompt the user to select a branch
	survey.AskOne(prompt, &selectedBranch.Name)

	for _, branch := range branches {
		if branch.Name == selectedBranch.Name {
			selectedBranch = branch
		}
	}

	// Update the application's state with the selected branch
	kvstore := app.GetKVStore()

	// TODO: Handle errors
	kvstore.SetString(constants.CURRENT_BRANCH_NAME, selectedBranch.Name)
	kvstore.SetUInt(constants.CURRENT_BRANCH_ID, selectedBranch.ID)

	fmt.Printf("Branch selected: %s\n", selectedBranch.Name)
}

func init() {
	selectCmd.Flags().StringP("branch", "b", "", "Select a branch (optional)")
}
