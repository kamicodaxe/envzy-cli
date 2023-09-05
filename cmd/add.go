// cmd/add.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [env-file]",
	Short: "Add environment files to staging for tracking changes",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envFile := args[0]
		fmt.Printf("Adding environment file: %s\n", envFile)
		// Logic here
		fmt.Println("Environment file added to staging.")
	},
}
