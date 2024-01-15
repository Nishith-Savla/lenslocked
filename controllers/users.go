package controllers

import (
	"fmt"
	"github.com/Nishith-Savla/lenslocked/models"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
	UserService *models.UserService
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	user, err := u.UserService.Create(r.PostFormValue("email"), r.PostFormValue("password"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %+v\n", user)
}
