package cmd

import (
	"fmt"
	"github.com/ErfiDev/golang-task-manager/structure"
	"github.com/spf13/cobra"
	"log"
)

var readCmd = &cobra.Command{
	Use: "read [flags]",
	Short: "for reading tasks",
	Long: `		we can read the tasks by using -a stand for all tasks
		-c for completed tasks and -u for uncompleted tasks, use one of them in 
		when writing command
		the priority of the flags: --all -a, --comp -c, --uncomp -u
	`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := map[string]bool{
			"all": all,
			"comp": comp,
			"uncomp": uncomp,
		}

		for key, value := range flags {
			if value {
				tasks, err := structure.ReadTasks(key)
				if err != nil {
					log.Println("Error on reading tasks: %s", err)
				}

				if len(tasks) > 0 {
					for _, val := range tasks {
						fmt.Println("------------------------------")
						fmt.Println("title: ", val.Title)
						fmt.Println("is done?: ", val.IsDone)
						fmt.Println("------------------------------")
					}
					return
				} else {
					fmt.Println("Nothing to show, let's add!")
					fmt.Println("run the command:")
					fmt.Println("add --title=your task title")
				}

				break
			}
		}
	},
}
