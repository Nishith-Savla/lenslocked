package main

import (
	"database/sql"
	"fmt"
	"github.com/Nishith-Savla/lenslocked/models"
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
	fmt.Printf("%s", cfg)
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

	// id := 4
	// row := db.QueryRow(`
	// 	SELECT name, email FROM users WHERE id = $1;
	// `, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err != nil && !errors.Is(err, sql.ErrNoRows) {
	// 	panic(err)
	// }
	// fmt.Printf("User information: name=%q email=%q\n", name, email)

	// userID := 1
	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err = db.Exec(`
	// 		INSERT INTO orders (user_id, amount, description)
	// 		VALUES ($1, $2, $3);
	// 	`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Created fake orders!")
	//
	// type Order struct {
	// 	ID          int16
	// 	UserID      int16
	// 	Amount      int
	// 	Description string
	// }
	//
	// var (
	// 	orders []Order
	// 	userID int16 = 1
	// )
	// rows, err := db.Query(`
	// 	SELECT id, amount, description FROM orders WHERE user_id = $1;
	// `, userID)
	// if err != nil {
	// 	panic(err)
	// }
	// defer func(rows *sql.Rows) {
	// 	err := rows.Close()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }(rows)
	// for rows.Next() {
	// 	order := Order{UserID: userID}
	// 	err := rows.Scan(&order.ID, &order.Amount, &order.Description)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	orders = append(orders, order)
	// }
	// if err := rows.Err(); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Orders: %#v\n", orders)

	us := models.UserService{DB: db}
	user, err := us.Create("nishith@email.com", "password123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created user: %#v\n", user)
}
