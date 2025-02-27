package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (n Note) Display() {
	fmt.Printf("Title: %s\nDescription: %s\nCreated At: %s\n", n.Title, n.Description, n.CreatedAt)
}

func (n Note) Save() error {
	filename := strings.ReplaceAll(n.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}

func New(title, description string) (Note, error) {

	if title == "" || description == "" {
		return Note{}, errors.New("Title or description is empty")
	}
	return Note{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}
