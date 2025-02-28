package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NoteProject/note"
	"github.com/NoteProject/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputabble interface {
	saver
	Display()
}

func getNoteData() (string, string, error) {
	title := getUserInput("Content Title : ")
	description := getUserInput("Content Description: ")

	return title, description, nil
}

func main() {
	title, description, err := getNoteData()
	textdata := getUserInput("Enter text: ")
	text, err := todo.New(textdata)
	note, err := note.New(title, description)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(text)
	if err != nil {
		return
	}

	err = outputData(note)
}

func outputData(s outputabble) error {
	s.Display()
	err := s.Save()
	if err != nil {
		fmt.Println("error saving note: ", err)
		return err
	}
	fmt.Println("Note saved successfully")
	return nil
}

func saveData(s saver) error {
	err := s.Save()
	if err != nil {
		fmt.Println("error saving note: ", err)
		return err
	}
	fmt.Println("Note saved successfully")
	return nil
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
