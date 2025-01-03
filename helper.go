package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber

}

func CreateFile() {
	fmt.Println("Writing file")
	file, err := os.Create("test.txt")

	if err != nil {
		panic(err)
	}

	length, err := file.WriteString("welcome to golang" +
		"demonstrates reading and writing operations to a file in golang.")

	if err != nil {
		panic(err)
	}
	fmt.Printf("File name: %s", file.Name())
	fmt.Printf("\nfile length: %d\n", length)
}

func ReadFile() {
	fmt.Println("Reading file")
	fileName := "test.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("file name " + fileName)
	fmt.Printf("file size %d\n", len(data))
	fmt.Printf("file content : %s\n", data)

}

func RemoveFile() {
	e := os.Remove("test.txt")
	if e != nil {
		log.Fatal(e)
	}
}
