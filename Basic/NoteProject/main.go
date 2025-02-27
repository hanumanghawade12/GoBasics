package main

import (
	"errors"
	"fmt"
)

func getNoteData() (string, string, error) {
	title, err := getUserInput("Content Title : ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	description, err := getUserInput("Content Description: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return title, description, nil
}

func main() {
	title, description, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Title: ", title, "\nDescription: ", description)
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var name string
	fmt.Scanln(&name)
	if name == "" {
		return "", errors.New("Name is empty")
	}

	return name, nil
}
