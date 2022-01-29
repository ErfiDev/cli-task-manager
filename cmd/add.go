package cmd

import (
	"fmt"
	"github.com/ErfiDev/golang-task-manager/structure"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"log"
)

var addCmd = &cobra.Command{
	Use: "add [flags] [args]",
	Short: "Add task",
	Long: "Adding a new task in tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(taskTitle) > 0 {
			task := structure.Task{
				Id: uuid.NewString(),
				Title: taskTitle,
				IsDone: false,
			}

			err := task.Add()
			if err != nil {
				log.Println("Error on adding task: %s", err)
				return
			}

			log.Println("Task added!")
			return
		}

		fmt.Println("Please contain --title flag")
	},
}