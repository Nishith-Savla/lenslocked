package models

type User struct {
	ID           uint16
	Email        string
	PasswordHash string
}
