package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "task [command] [flags] [args]",
	Short: "task manager cli",
	Long: ` 	task manager go cli application powered by cobra 
	framework and viper, managing tasks, add, delete and update
	tasks is the features of this application
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please enter --help flag after command to show you all commands and flags")
	},
}