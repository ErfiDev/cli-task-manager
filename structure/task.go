package structure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Task struct {
	Title string
	Id string
	IsDone bool
}

func (t *Task) Add() error {
	toJson, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("./db/%s.json", t.Id), toJson, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadTasks(by string) ([]*Task, error) {
	files, err := ioutil.ReadDir("./db")
	if err != nil {
		return []*Task{}, err
	}

	var readFiles []*Task
	for _, val := range files {
		file, err := ioutil.ReadFile("./db/"+ val.Name())
		if err != nil {
			return []*Task{}, err
		}

		var unMarshalTask Task
		err = json.Unmarshal(file, &unMarshalTask)
		if err != nil {
			return []*Task{}, err
		}

		if by == "all" {
			readFiles = append(readFiles, &unMarshalTask)
		} else if by == "comp" {
			if unMarshalTask.IsDone {
				readFiles = append(readFiles, &unMarshalTask)
			}
		} else if by == "uncomp" {
			if !unMarshalTask.IsDone {
				readFiles = append(readFiles, &unMarshalTask)
			}
		} else {
			continue
		}
	}

	return readFiles , nil
}
