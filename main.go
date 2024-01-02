package main

import (
	"fmt"
	"net/http"

	"github.com/Nishith-Savla/lenslocked/controllers"
	"github.com/Nishith-Savla/lenslocked/templates"
	"github.com/Nishith-Savla/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {

		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("starting server on :3000...")
	http.ListenAndServe(":3000", r)

}
