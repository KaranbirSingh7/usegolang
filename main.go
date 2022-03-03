package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/karanbirsingh7/usegolang/controllers"
	"github.com/karanbirsingh7/usegolang/views"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		http.Error(w, "There was an error parsing template.", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userID")
	fmt.Fprintf(w, "Passed value: %v", id)
}

func main() {

	r := chi.NewRouter()
	// log incoming requests
	r.Use(middleware.Logger)

	// parse templates

	tpl, err := views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	r.Get("/gallery/{userID}", galleryHandler)

	// hanlde unknown routes
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// start server
	fmt.Println("Server starting on :3000")
	http.ListenAndServe(":3000", r)
}
