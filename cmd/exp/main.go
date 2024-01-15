package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg *PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	// Create a table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SMALLSERIAL PRIMARY KEY,
			name VARCHAR(100),
			email VARCHAR(100) UNIQUE NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS orders (
		    id SMALLSERIAL PRIMARY KEY,
		    user_id SMALLINT NOT NULL,
		    amount INT,
		    description VARCHAR(1000)
		);			
	`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created!")

	// Insert some data
	// name := "Nishith"
	// email := "nish@email.com"
	// row := db.QueryRow(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2) RETURNING id;
	// `, name, email)
	// var id int16
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created!, id:", id)

	id := 1
	row := db.QueryRow(`
		SELECT name, email FROM users WHERE id = $1;
	`, id)
	var name, email string
	err = row.Scan(&name, &email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	fmt.Printf("User information: name=%q email=%q\n", name, email)
}
