package main

import "fmt"

type Contact struct {
	Id    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact with name %s already exists.\n", name)
		return
	}

	contact := Contact{
		Id:    len(contactList) + 1,
		Name:  name,
		Email: email,
		Phone: phone,
	}

	contactList = append(contactList, contact)
	contactIndexByName[name] = contact.Id - 1
	fmt.Printf("Contact %s added successfully.\n", name)
}

func findContactByName(name string) *Contact {
	if index, exists := contactIndexByName[name]; exists {
		return &contactList[index]
	}
	return nil
}

func listContacts() {
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("Contact List:")
	for _, contact := range contactList {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)
	}
}

func main() {
	addContact("Alice", "alice@example.com", "123-456-7890")
	addContact("Bob", "bob@example.com", "098-765-4321")
	addContact("Charlie", "charlie@example.com", "555-555-5555")
	addContact("Alice", "alice2@example.com", "111-111-1111")

	fmt.Println("Contact Management System")
	fmt.Println("-------------------------")
	listContacts()
}
