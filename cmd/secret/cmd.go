package secret

import (
	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:     "secret",
	Aliases: []string{"s"},
	Short:   "Secret operations",
	Long:    "Manage secret.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Run(listCmd, nil)
	},
}

func init() {
	// Add any flags if needed for the 'secret' command

	// Add subcommands to the 'secret' command
	secretCmd.AddCommand(createCmd)
	secretCmd.AddCommand(listCmd)
	secretCmd.AddCommand(updateCmd)
	secretCmd.AddCommand(deleteCmd)
}

func AddToRoot(rootCmd *cobra.Command) {
	rootCmd.AddCommand(secretCmd)
}
