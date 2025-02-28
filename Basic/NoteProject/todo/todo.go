package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

// opmment adeed

func (todo Todo) Display() {
	fmt.Println(todo)
}

func (todo Todo) Save() error {
	filename := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}

func New(title string) (Todo, error) {

	if title == "" {
		return Todo{}, errors.New("title or description is empty")
	}
	return Todo{
		Text: title,
	}, nil
}
