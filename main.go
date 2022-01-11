package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	City string
	// State string
	// Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {

	dir := "./"

	fmt.Println(os.Geteuid())

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"DevOps", "12", "42342343", "devops", Address{"not appropriate", "1234"}},
		{"DevOps", "12", "42342343", "devops", Address{"not appropriate", "1234"}},
		{"DevOps", "12", "42342343", "devops", Address{"not appropriate", "1234"}},
		{"DevOps", "12", "42342343", "devops", Address{"not appropriate", "1234"}},
		{"DevOps", "12", "42342343", "devops", Address{"not appropriate", "1234"}},
		{"DevOps", "12", "42342343", "devops", Address{"not appropriate", "1234"}},
		{"DevOps", "12", "42342343", "devops", Address{"not appropriate", "1234"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(records)

	allUsers := []User{}

	for _, value := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(value), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allUsers = append(allUsers, employeeFound)
	}
	fmt.Println(allUsers)

	if err := db.Delete("user", "DevOps"); err != nil {
		fmt.Println("Error", err)
	}

	if err := db.DeleteAll("user", ""); err != nil {
		fmt.Println("Error", err)
	}
}
