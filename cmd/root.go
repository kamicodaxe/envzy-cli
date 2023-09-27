package cmd

import (
	"fmt"
	"os"

	"github.com/kamicodaxe/envzy-cli/cmd/branch"
	"github.com/kamicodaxe/envzy-cli/cmd/project"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envzy",
	Short: "A CLI tool for managing environment variables and configurations",
	Long: `
    Envzy is a command-line tool for managing environment variables and configuration files. It provides a simple and intuitive way to manage your environment variables and configurations within projects, branches, and secrets.

    Envzy is designed to streamline the process of securely handling environment variables and configuration files, making it easier for developers to work on multiple projects and collaborate with team members.

    Key Features:
    - Create, list, select, update, and delete projects.
    - Manage branches or configurations within projects.
    - Add, list, update, and delete secrets within branches.
    - Switch between projects or environments with ease.
    - Securely manage environment variables and configurations.
    - Collaborate with team members by inviting them to projects.
    - Track changes to environment files using version control.

    To get started, use the "envzy" command along with subcommands to perform various operations related to managing your environment variables and configurations.

    For help with a specific command, use "envzy [command] --help."

    Visit the Envzy website for documentation and updates: https://envzy.io
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Envzy CLI!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	project.AddToRoot(rootCmd)
	branch.AddToRoot(rootCmd)

	rootCmd.AddCommand(dashboardCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(statusCmd)
}

func main() {
	Execute()
}
