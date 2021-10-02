package task

type Task []structure

type structure struct {
	Title string
	IsDone bool
}

func AddTask(task *Task , title string , isDone bool) {
	*task = append(*task , structure{title , isDone})
}
