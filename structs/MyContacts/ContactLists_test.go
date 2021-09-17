package MyContacts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndGet(test *testing.T) {

	// creating one user
	var contacts ContactList
	newContact := Contact{1, "zero", "white", "+9989099999998", "zero@gmail.com", "Lead"}
	contacts.Create(newContact)

	// actual and expected results
	actualResult := contacts.Get(1)
	expectedResult := Contact{1, "zero", "white", "+9989099999998", "zero@gmail.com", "Lead"}

	if actualResult != expectedResult {
		fmt.Printf("Expected %v, but got %v", expectedResult, actualResult)
		test.Fail()
	}
}

func TestGetAll(test *testing.T) {

	// creating two users
	var contacts ContactList
	contacts.list = make(map[int]Contact)
	contacts.list[1] = Contact{1, "zero", "white", "+9989099999998", "zero@gmail.com", "Lead"}
	contacts.list[2] = Contact{2, "what", "black", "+9989998945644", "mark@gmail.com", "Team"}

	// actual and expected results
	expectedResult := len(contacts.list)
	actualResult := len(contacts.GetAll())

	if actualResult != expectedResult {
		fmt.Printf("Expected %v, but got %v", expectedResult, actualResult)
		test.Fail()
	}
}

func TestUpdate(test *testing.T) {

	//creating one user
	var contacts ContactList
	contacts.list = make(map[int]Contact)
	contacts.list[1] = Contact{1, "zero", "white", "+9989099999998", "zero@gmail.com", "Lead"}

	//updating the user
	newContact := Contact{1, "what", "black", "+9989998945644", "mark@gmail.com", "Team"}
	contacts.Update(1, newContact)

	// actual and expected results
	expectedResult := Contact{1, "what", "black", "+9989998945644", "mark@gmail.com", "Team"}
	actualResult := contacts.Get(1)

	if actualResult != expectedResult {
		fmt.Printf("Expected %v, but got %v", expectedResult, actualResult)
		test.Fail()
	}
}

func TestDelete(test *testing.T) {

	//creating one user
	var contacts ContactList
	contacts.list = make(map[int]Contact)
	contacts.list[1] = Contact{1, "zero", "white", "+9989099999998", "zero@gmail.com", "Lead"}

	//deleting the user
	contacts.Delete(1)

	// actual and expected results
	expectedResult := Contact{0, "", "", "", "", ""}
	actualResult := contacts.Get(1)

	if actualResult != expectedResult {
		fmt.Printf("Expected %v, but got %v", expectedResult, actualResult)
		test.Fail()
	}
}

func TestDeleteAll(test *testing.T) {

	//creating two users
	var contacts ContactList
	contacts.list = make(map[int]Contact)
	contacts.list[1] = Contact{1, "zero", "white", "+9989099999998", "zero@gmail.com", "Lead"}
	contacts.list[2] = Contact{2, "what", "black", "+9989998945644", "mark@gmail.com", "Team"}

	//deleting the user
	contacts.DeleteAll()

	// actual and expected results
	expectedResult := len(contacts.list)
	actualResult := len(contacts.GetAll())

	// using assert package to compare actual and expected results
	assert.Equal(test, expectedResult, actualResult, "All the contacts should be deleted")
}
