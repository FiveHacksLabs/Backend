package services

import (
	"database/sql"
	"errors"
	"users/models"
	"users/utils"
)

func RegisterUser(user models.User) error {
	query := "INSERT INTO users (email, full_name, password) VALUES (?, ?, ?)"
	_, err := utils.GetDB().Exec(query, user.Email, user.FullName, user.Password)
	return err
}

func LoginUser(email, password string) (models.User, error) {
	var user models.User
	query := "SELECT id, email, full_name FROM users WHERE email = ? AND password = ?"
	row := utils.GetDB().QueryRow(query, email, password)
	err := row.Scan(&user.ID, &user.Email, &user.FullName)
	if err == sql.ErrNoRows {
		return user, errors.New("invalid credentials")
	}
	return user, err
}
