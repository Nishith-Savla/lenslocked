package models

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	ID           uint16
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	email = strings.ToLower(email)
	user := &User{
		Email:        email,
		PasswordHash: string(bytes),
	}

	result := us.DB.QueryRow(`
		INSERT INTO users (email, password_hash) 
		VALUES ($1, $2) RETURNING id;
	`, user.Email, user.PasswordHash)
	err = result.Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
