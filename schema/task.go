package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Task []Structure

type Structure struct {
	Title string
	IsDone bool
	Id int
}

func AddTask(task *Task , title string) {
	*task = append(*task , Structure{title , false , len(*task)+1})
}

func SaveToFile(task Task) {
	json , err := json.Marshal(task)

	if err != nil {
		log.Fatal(err);
	} else {
		err := ioutil.WriteFile("tasks.json" , json , 0644)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("file created");
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
	fmt.Println(find)

	find.IsDone = true

	tasks[index] = find
}
