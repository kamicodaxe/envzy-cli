package config

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config files operations",
	Long:  "Manage config files and secrets.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	configCmd.AddCommand(addCmd)
	configCmd.AddCommand(commitCmd)
	configCmd.AddCommand(pushCmd)
	configCmd.AddCommand(pullCmd)
}

func AddToRoot(rootCmd *cobra.Command) {
	rootCmd.AddCommand(configCmd)
}
