package main

import (
	"fmt"
	task "github.com/erfidev/task-manager-golang/schema"
)

func main() {
	init := task.Task{}

	task.AddTask(&init , "learn go")
	task.AddTask(&init , "learn java")


	fmt.Println(init)

	task.SaveToFile(init)

	task.CompleteTask(init , 19727887)

	fmt.Println(init)
}
