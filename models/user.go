package models

import (
	"fmt"

	"github.com/qsmsoft/go-rest-api/db"
	"github.com/qsmsoft/go-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	err = stmt.QueryRow(u.Email, hashedPassword).Scan(&u.ID)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
