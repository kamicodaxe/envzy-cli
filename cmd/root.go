package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envzy",
	Short: "A CLI tool for managing environment variables and configurations",
	Long: `
    Envzy is a command-line tool for managing environment variables and configuration files. It provides a simple and intuitive way to switch between projects, manage branches or configurations within projects, and track changes to environment files.

    Envzy is designed to streamline the process of securely handling environment variables and configuration files, making it easier for developers to work on multiple projects and collaborate with team members.

    Key Features:
    - Switch between projects or environments with ease.
    - Manage branches or configurations within projects.
    - Track changes to environment files using version control.
    - Collaborate with team members by inviting them to projects.
    - Securely manage environment variables and configurations.

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
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(branchCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(dashboardCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(projectsCmd)
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(selectCmd)
	rootCmd.AddCommand(teamAddCmd)
}

func main() {
	Execute()
}
