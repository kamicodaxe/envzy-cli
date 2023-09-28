package config

import (
	"fmt"
	"log"

	"github.com/kamicodaxe/envzy-cli/internal/api"
	"github.com/kamicodaxe/envzy-cli/internal/models"
	"github.com/kamicodaxe/envzy-cli/utils/dotenv"
	"github.com/spf13/cobra"
)

var pullProject string
var pullBranch string

var pullCmd = &cobra.Command{
	Use:   "pull [path/to/dotenv]",
	Short: "Pulls all env variables from envy and writes them to an env-file",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		var selectedProject *models.Project
		var selectedBranch *models.Branch
		envFile := args[0]
		fmt.Printf("Creating environment file: %s\n", envFile)

		if pullProject == "" {
			activeProject := api.GetCurrentProject()

			if activeProject != nil {
				pullProject = activeProject.Name
				selectedProject = activeProject
			} else {
				selectedProject = SelectProject()
				if selectedProject != nil {
					pullProject = selectedProject.Name
				} else {
					fmt.Println("No Project selected")
					return
				}
			}
		}

		if pullBranch == "" {
			activeBranch := api.GetCurrentBranch()

			if activeBranch != nil {
				pullBranch = activeBranch.Name
				selectedBranch = activeBranch
			} else {
				selectedBranch = SelectBranch(selectedProject.ID)
				if selectedBranch != nil {
					pullBranch = selectedBranch.Name
				} else {
					fmt.Println("No Branch selected")
					return
				}
			}
		}

		fmt.Printf("Project: %s\n", pullProject)
		fmt.Printf("Branch: %s\n", pullBranch)

		// Retrieve secrets for the selected project and branch from the database
		secrets, err := api.GetSecretsByProjectAndBranch(selectedProject.ID, selectedBranch.ID)
		if err != nil {
			log.Fatalln("Error retrieving secrets.", err)
			return
		}
		// For each secret, call WriteToEnvFile to write them to the env file
		for _, secret := range secrets {
			dotenv.WriteToEnvFile(envFile, secret.Name, secret.Value, secret.Comment)
		}

	},
}

func init() {
	// Add flags or additional configuration for the "pull" command here, if needed
	// TODO: Uncomment the flags below and use them to enable the flexibility to specify a project and branch
	// pullCmd.PersistentFlags().StringVarP(&pullProject, "project", "p", "", "Project name")
	// pullCmd.PersistentFlags().StringVarP(&pullBranch, "branch", "b", "", "Branch name")
}
