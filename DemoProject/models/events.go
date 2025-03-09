package models

import (
	"demoProject/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)`
	statment, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statment.Close()
	result, err := statment.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	e.ID = id
	return err
}

func GetEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	var event Event
	query := "SELECT * FROM EVENTS WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil

}

func (e Event) Upadate() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ? WHERE id = ?`
	statment, err := db.DB.Prepare(query)
	if err != nil {
		return err

	}
	defer statment.Close()
	_, err = statment.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId, e.ID)
	return err

}

func (e Event) DeleteEvent() error {
	query := `DELETE FROM events WHERE id = ?`
	statment, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statment.Close()
	_, err = statment.Exec(e.ID)
	return err
}

func (e Event) RegisterForEvent(userId int64) error {
	query := `INSERT INTO registrations (event_id, user_id) VALUES (?, ?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.ID, userId)
	return err
}

func (e Event) UnregisterFromEvent(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.ID, userId)
	return err
}
