package cmd

import (
	"fmt"
	"github.com/ErfiDev/golang-task-manager/structure"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command {
	Use: "remove [flags]",
	Short: "for remove a task",
	Long: `   removing a task with remove command, should
	pass the [ --id ] flags for finding a task and removing
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(taskId) > 0 {
			_, err := structure.RemoveTask(taskId)
			if err != nil {
				fmt.Println("error on removing task: ", err)
				return
			}

			fmt.Println("Task removed!")
		} else {
			fmt.Println("Please set the --id flag!")
		}
	},
}
