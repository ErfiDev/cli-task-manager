package cmd

import (
	"fmt"
	"github.com/ErfiDev/golang-task-manager/structure"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use: "complete [flags]",
	Short: "For complete tasks",
	Long:`	Command for complete a task, you should declare [ --id ] flag`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(taskId) > 0 {
			_, err := structure.CompleteTask(taskId)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Task successfully completed!")
		} else {
			fmt.Println("Please contain --id flag value! for more detail run --help flag")
		}
	},
}
