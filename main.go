package main

import (
	"database/sql"
	"fmt"
	"github.com/Nishith-Savla/lenslocked/models"
	"log"
	"net/http"

	"github.com/Nishith-Savla/lenslocked/controllers"
	"github.com/Nishith-Savla/lenslocked/templates"
	"github.com/Nishith-Savla/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"contact.gohtml", "tailwind.gohtml",
	))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS,
		"faq.gohtml", "tailwind.gohtml",
	))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	usersController := controllers.Users{
		UserService: &models.UserService{DB: db},
	}
	usersController.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	r.Get("/signup", usersController.New)
	r.Post("/users", usersController.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("starting server on :3000...")
	log.Fatalln(http.ListenAndServe(":3000", r))
}
