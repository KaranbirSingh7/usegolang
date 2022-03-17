package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/karanbirsingh7/usegolang/controllers"
	"github.com/karanbirsingh7/usegolang/templates"
	"github.com/karanbirsingh7/usegolang/views"
)

func galleryHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userID")
	fmt.Fprintf(w, "Passed value: %v", id)
}

func main() {

	r := chi.NewRouter()
	// log incoming requests
	r.Use(middleware.Logger)

	// parse templates

	tpl := views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "faq.gohtml"))
	r.Get("/faq", controllers.FAQs(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "signup.gohtml"))
	r.Get("/signup", controllers.StaticHandler(tpl))

	r.Get("/gallery/{userID}", galleryHandler)

	// hanlde unknown routes
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// start server
	fmt.Println("Server starting on :3000")
	http.ListenAndServe(":3000", r)
}
