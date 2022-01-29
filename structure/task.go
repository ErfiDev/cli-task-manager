package structure

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
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

func CompleteTask(id string) (bool, error) {
	files, err := ioutil.ReadDir("./db")
	if err != nil {
		return false, err
	}

	for _, val := range files {
		file, err := ioutil.ReadFile("./db/"+val.Name())
		if err != nil {
			return false, err
		}

		var unMarshalFile Task
		err = json.Unmarshal(file, &unMarshalFile)
		if err != nil {
			return false, err
		}

		if unMarshalFile.Id == id {
			// writing again this file by changing idDone = to true
			if unMarshalFile.IsDone {
				return false, errors.New("the task are done")
			}

			unMarshalFile.IsDone = true
			newData, err := json.Marshal(unMarshalFile)
			if err != nil {
				return false, err
			}
			err = ioutil.WriteFile("./db/"+ val.Name(), newData, fs.ModeAppend)
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}