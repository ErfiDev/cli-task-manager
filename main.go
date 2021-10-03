package main

import (
	"fmt"
	task "github.com/erfidev/golang-task-manager/schema"
)

func main() {
	program := true
	initTask := task.Task{}

	for program {
		fmt.Println("please select the number of item:")
		fmt.Println("-1: add Task")
		fmt.Println("-2: show tasks")
		fmt.Println("-3: complete task")
		fmt.Println("-4: print to file")
		fmt.Println("-5: quit program")

		var input int
		fmt.Scanln(&input)

		if input == 1 {
			Title := ""
			fmt.Println("enter the task title:")
			fmt.Scanln(&Title)

			task.AddTask(&initTask , Title)
			fmt.Println("successfully add!")
			continue
		}
		if input == 2 {
			for _ , value := range initTask {
				fmt.Print("id:" , value.Id)
				fmt.Print(" title:" , value.Title)
				fmt.Println(" status:" , value.IsDone)
			}
			continue
		}
		if input == 3 {
			for _ , value := range initTask {
				fmt.Print("id:" , value.Id)
				fmt.Print(" title:" , value.Title)
				fmt.Println(" status:" , value.IsDone)
			}
			fmt.Println("enter the target task for completing: ")
			forCompleting := 0
			fmt.Scanln(&forCompleting)

			if forCompleting < len(initTask) + 1 {
				task.CompleteTask(initTask , forCompleting)
				fmt.Println("successfully complete!")
				continue
			} else {continue}
		}
		if input == 4 {
			msg := task.SaveToFile(initTask)
			if msg == "something wrong!" {
				fmt.Println(msg)
				continue
			} else {
				fmt.Println("file created on path: /tasks.json")
				continue
			}
		}
		if input == 5 {
			program = false
			break
		} else {
			program = false
			break
		}
	}
}
