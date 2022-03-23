package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// users object will have a template already parsed that can be rendered
	u.Templates.New.Execute(w, nil)
}

// Create creates new user sent by signUp form
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// email := r.PostForm.Get("email")
	// password := r.PostForm.Get("password")

	// // validate
	// if len(email) == 0 {
	// 	http.Error(w, "email cannot be empty", http.StatusBadRequest)
	// 	return
	// }
	// if len(password) == 0 {
	// 	http.Error(w, "password cannot be empty", http.StatusBadRequest)
	// 	return
	// }
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Fprintf(w, email, password)
}
