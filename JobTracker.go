package main

import (
	. "fmt"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("AppliedJobs.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // closes the file when program finishes

	GetExisting()
	Println("What is your first name?")
	var firstName string
	_, err = Scanln(&firstName)
	if err != nil {
		log.Fatal(err)
	}
	Println("What is your last name?")
	var lastName string
	_, err = Scanln(&lastName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(firstName +" "+ lastName + "\n")
	if err != nil {
		log.Fatal(err)
	}
	GetExisting()
}

func GetExisting() {
	existingBytes, err := os.ReadFile("AppliedJobs.txt")
	if err != nil {
		log.Fatal(err)
	}
	existing := string(existingBytes)
	Println(existing)
}
