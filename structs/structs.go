package main

import (
	"time"
)

type ContactList struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  string
}

// C-create,
// R- read, GET(id) & GETALL
// U - update(id),
// D - delete(id), & DELETEALL

func (jopa *ContactList) create(firstName string, lastName string) {
	jopa.FirstName = firstName
	jopa.LastName = lastName
}

type TaskList struct {
	ID        int
	Name      string
	Status    int
	Priority  int
	CreatedAt time.Time
	CreatedBy string
	DueDate   time.Time
}

func (object *TaskList) create() {

}
