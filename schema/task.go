package task

import (
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"encoding/json"
	"log"
)

type Task []structure

type structure struct {
	Title string
	IsDone bool
}

func AddTask(task *Task , title string , isDone bool) {
	*task = append(*task , structure{title , isDone})
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
