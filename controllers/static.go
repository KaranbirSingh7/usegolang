package controllers

import (
	"net/http"

	"github.com/karanbirsingh7/usegolang/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
