package main

import (
	"fmt"
)

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  string
}

type ContactList struct {
	list map[int]Contact
}

var id int = 1

func (contacts *ContactList) Create(newContact Contact) {
	if contacts.list == nil {
		contacts.list = make(map[int]Contact)
	}
	contacts.list[id] = newContact
	id++
}

func (contacts *ContactList) Get(id int) Contact {
	return contacts.list[id]
}

func (contacts *ContactList) GetAll() map[int]Contact {
	return contacts.list
}

func (contacts *ContactList) Update(ids int, newContact Contact) {
	contacts.list[ids] = newContact
}

func (contacts *ContactList) Delete(ids int) {
	delete(contacts.list, ids)
}

func (contacts *ContactList) DeleteAll() {
	for elements := range contacts.list {
		delete(contacts.list, elements)
	}
}

func main() {
	var list ContactList

	// I am creating two contacts using Create() and displaying them using Get()
	newContact := Contact{
		ID:        id,
		FirstName: "zero",
		LastName:  "white",
		Phone:     "+9989099999998",
		Email:     "zero@gmail.com",
		Position:  "Lead",
	}
	list.Create(newContact)

	newContact = Contact{
		ID:        id,
		FirstName: "mark",
		LastName:  "fake",
		Phone:     "+9989099999998",
		Email:     "mark@gmail.com",
		Position:  "QA",
	}
	list.Create(newContact)

	fmt.Println(list.Get(1))
	fmt.Println(list.Get(2))

	// I am updating first contact using Update()
	newContact = Contact{
		ID:        1,
		FirstName: "test",
		LastName:  "Update",
		Phone:     "+998905555555",
		Email:     "update@gmail.com",
		Position:  "Updated",
	}
	list.Update(1, newContact)

	// I am deleting second contact using Delete()
	list.Delete(2)
	fmt.Println(list.GetAll())

	// I am deleting all contacts using DeleteAll()
	list.DeleteAll()
	fmt.Println(list.GetAll())
}
