package Tasklist

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var date time.Time = time.Date(2021, 9, 18, 02, 05, 54, 15, time.Now().Location())

func TestCreateAndGet(test *testing.T) {

	// creating a task
	var tasks TaskList
	newTask := Task{1, "Writing", 2, 5, date, "Bekhzod", date}
	tasks.create(newTask)

	// actual and expected result
	expectedResult := Task{1, "Writing", 2, 5, date, "Bekhzod", date}
	actualResult := tasks.Get(1)

	assert.Equal(test, expectedResult, actualResult, "Created task should be same as expected task")
}

func TestGetAll(test *testing.T) {

	// creating two tasks
	var tasksList TaskList
	tasksList.tasks = make(map[int]Task)
	tasksList.tasks[1] = Task{1, "Writing", 2, 5, date, "Bekhzod", date}
	tasksList.tasks[2] = Task{2, "Reading", 1, 3, date, "Khulkar", date}

	// actual and expected result of number of tasks
	expectedNumberResult := len(tasksList.tasks)
	actualNumberResult := len(tasksList.GetAll())
	assert.Equal(test, expectedNumberResult, actualNumberResult, "Should return correct number of users")

	// actual and expected result of all info about task
	expectedResult := tasksList.tasks
	actualResult := tasksList.GetAll()
	assert.Equal(test, expectedResult, actualResult, "Should return correct number of tasks")
}

func TestUpdate(test *testing.T) {

	//creating one task
	var tasksList TaskList
	tasksList.tasks = make(map[int]Task)
	tasksList.tasks[1] = Task{1, "Writing", 2, 5, date, "Bekhzod", date}

	//updating the taks
	updatedTask := Task{1, "Speaking", 1, 1, date, "Mark", date}
	tasksList.Update(1, updatedTask)

	// actual and expected results
	expectedResult := Task{1, "Speaking", 1, 1, date, "Mark", date}
	actualResult := tasksList.Get(1)
	assert.Equal(test, expectedResult, actualResult, "Should correctly update taks")
}
func TestDelete(test *testing.T) {

	//creating one task
	timeDate := time.Date(0001, 01, 01, 00, 00, 00, 00, time.UTC)
	var tasksList TaskList
	tasksList.tasks = make(map[int]Task)
	tasksList.tasks[1] = Task{1, "Writing", 1, 5, date, "Bekhzod", date}

	//deleting the task
	tasksList.Delete(1)

	// actual and expected results
	expectedResult := Task{0, "", 0, 0, timeDate, "", timeDate}
	actualResult := tasksList.Get(1)

	assert.Equal(test, expectedResult, actualResult, "Should correctly update taks")
}

func TestDeleteAll(test *testing.T) {

	//creating two tasks
	var tasksList TaskList
	tasksList.tasks = make(map[int]Task)
	tasksList.tasks[1] = Task{1, "Writing", 1, 5, date, "Bekhzod", date}
	tasksList.tasks[2] = Task{2, "Reading", 1, 3, date, "Khulkar", date}

	//deleting the tasks
	tasksList.DeleteAll()

	// actual and expected results
	expectedResult := len(tasksList.tasks)
	actualResult := len(tasksList.GetAll())
	assert.Equal(test, expectedResult, actualResult, "All the tasks should be deleted")
}
