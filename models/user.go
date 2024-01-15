package models

import "database/sql"

type User struct {
	ID           uint16
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}
