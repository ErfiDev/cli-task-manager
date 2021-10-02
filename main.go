package main

import (
	"fmt"
	task "github.com/erfidev/task-manager-golang/schema"
)

func main() {
	init := task.Task{}

	task.AddTask(&init , "learn go" , false);

	fmt.Println(init)

	task.SaveToFile(init)
}
