package models

import (
	"errors"
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users(email, password)
	VALUES (?, ?)`

	var err error
	user.Password, err = utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := db.DB.Exec(query, user.Email, user.Password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	user.ID = id

	return err
}

func (user User) FindByEmail() (*User, error) {
	query := `SELECT password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)
	err := row.Scan(
		&user.Password,
	)

	return &user, err
}

func (user *User) ValidateCredentials() error {
	insertedPassword := user.Password
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)
	err := row.Scan(
		&user.ID,
		&user.Password,
	)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(insertedPassword, user.Password)

	if !passwordIsValid {
		return errors.New("Credential password")
	}
	return nil
}
