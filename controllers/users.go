package controllers

import (
	"net/http"

	"github.com/karanbirsingh7/usegolang/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// users object will have a template already parsed that can be rendered
	u.Templates.New.Execute(w, nil)
}
