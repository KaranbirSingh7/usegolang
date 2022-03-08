package controllers

import (
	"html/template"
	"net/http"

	"github.com/karanbirsingh7/usegolang/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQs(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes we do offer 30 days trail",
		},
		{
			Question: "What's your support hours?",
			Answer:   "24/7 but little slow on weekends",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
