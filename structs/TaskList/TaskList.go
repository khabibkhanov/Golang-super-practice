package Tasklist

import (
	"time"
)

type Task struct {
	ID        int
	Name      string
	Status    int
	Priority  int
	CreatedAt time.Time
	CreatedBy string
	DueDate   time.Time
}

type TaskList struct {
	tasks map[int]Task
}

var id int = 1

func (taskList *TaskList) create(task Task) {
	if taskList.tasks == nil {
		taskList.tasks = make(map[int]Task)
	}
	taskList.tasks[id] = task
	id++
}

func (taskList *TaskList) Get(id int) Task {
	return taskList.tasks[id]
}

func (taskList *TaskList) GetAll() map[int]Task {
	return taskList.tasks
}

func (taskList *TaskList) Update(id int, newTask Task) {
	taskList.tasks[id] = newTask
}

func (taskList *TaskList) Delete(id int) {
	delete(taskList.tasks, id)
}

func (taskList *TaskList) DeleteAll() {
	for task := range taskList.tasks {
		delete(taskList.tasks, task)
	}
}
