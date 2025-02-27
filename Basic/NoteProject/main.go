package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NoteProject/note"
)

func getNoteData() (string, string, error) {
	title := getUserInput("Content Title : ")
	description := getUserInput("Content Description: ")

	return title, description, nil
}

func main() {
	title, description, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
	note, err := note.New(title, description)
	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()
	err = note.Save()
	if err != nil {
		fmt.Println("error saving note: ", err)
		return
	}
	fmt.Println("Note saved successfully")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	// var name string
	// fmt.Scanln(&name)
	// if name == "" {
	// 	return ""
	// }

	bufioReader := bufio.NewReader(os.Stdin)
	name, err := bufioReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	name = strings.TrimSuffix(name, "\n")
	name = strings.TrimSuffix(name, "\r")

	return name
}
