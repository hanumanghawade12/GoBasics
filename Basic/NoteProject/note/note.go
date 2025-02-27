package note

import (
	"errors"
	"time"
)

type Note struct {
	Title       string
	Description string
	createdAt   time.Time
}

func New(title, description string) (Note, error) {

	if title == "" || description == "" {
		return Note{}, errors.New("Title or description is empty")
	}
	return Note{
		Title:       title,
		Description: description,
		createdAt:   time.Now(),
	}, nil
}
