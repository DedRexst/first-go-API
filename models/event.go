package models

import (
	"example.com/rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query :=
		`INSERT INTO events(name, location, date_time, user_id, description, date_time) 
		VALUES (?, ?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	exec, err := stmt.Exec(e.Name, e.Location, e.DateTime, e.UserID, e.Description, e.DateTime)
	if err != nil {
		return err
	}

	id, err := exec.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id
	return nil
}

func GetEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	var event Event

	query := `SELECT * FROM events WHERE id=?`
	rows, err := db.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserID,
	)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `UPDATE events 
		SET name=?, location=?, description=?, date_time=? 
		WHERE id=?`
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(event.Name, event.Location, event.Description, event.DateTime, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEventById(id int64) error {
	query := `
	DELETE FROM events where id=?
`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (event Event) Register(userID int64) error {
	query := `INSERT INTO registrations (user_id, event_id) VALUES (?,?)`
	_, err := db.DB.Exec(query, userID, event.ID)

	return err
}

func (event Event) CancelRegistration(userID int64) error {
	query := `DELETE FROM registrations where user_id = ? AND event_id = ?`
	_, err := db.DB.Exec(query, userID, event.ID)

	return err
}
