package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Task []Structure

type Structure struct {
	Title string
	IsDone bool
	Id int
}

func AddTask(task *Task , title string) {
	*task = append(*task , Structure{title , false , len(*task)+1})
	fmt.Println(*task)
}

func SaveToFile(task Task) string {
	json , err := json.Marshal(task)

	if err != nil {
		return "something wrong!"
	} else {
		err := ioutil.WriteFile("tasks.json" , json , 0644)
		if err != nil {
			return "something wrong!"
		} else {
			return "success"
		}
	}
}

func CompleteTask(tasks Task,id int) {
	var index int
	for ind , value := range tasks {
		if value.Id == id {
			index = ind
			break
		}else {
			continue
		}
	}

	find := tasks[index]

	find.IsDone = true

	tasks[index] = find
}
