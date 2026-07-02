package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID = 0

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	// Check whether the contact is already there;
	_, exist := contactIndexByName[name]
	if exist {
		fmt.Println("Contact already exist")
		return
	}
	contact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}

	contactList = append(contactList, contact)
	contactIndexByName[name] = contact.ID
	nextID++

	fmt.Printf("contact added: %+v\n", contact)
}

func findContact(name string) *Contact {
	// Check even if name exists in the first place.
	contactIndex, exist := contactIndexByName[name]
	if !exist {
		return nil
	}

	return &contactList[contactIndex]
}

func listContacts() {
	if len(contactList) == 0 {
		fmt.Println("No contacts")
		return
	}

	for _, contact := range contactList {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	}
}

func main() {
	addContact("John", "abc@gmail.com", "123456789")
	addContact("Kunle", "def@gmail.com", "987654321")
	addContact("Wale", "fgh@gmail.com", "443456789")

	listContacts()
	contact := findContact("Kunle")

	if contact == nil {
		fmt.Println("Contact not found")
	} else {
		fmt.Println("Contact pointer:", contact.Name)
	}
}
