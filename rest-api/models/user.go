package models

import (
	"github.com/BaGorK/lets-go/rest-api/db"
	"github.com/BaGorK/lets-go/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=6" json:"password"`
}

func (user *User) CreateUser() error {
	query := `
	insert into users (email, password) values (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(user.Email, hashedPass)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = userID

	return nil
}
