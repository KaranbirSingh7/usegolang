package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

//TMP test func for  file upload scenario
func handleUpload(w http.ResponseWriter, r *http.Request) {
	file, fileHandler, err := r.FormFile("avatar")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("./uploaded/"+fileHandler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	io.Copy(f, file)
	log.Println("file saved locally as:", fileHandler.Filename)
	fmt.Fprintf(w, "file uploaded onto server")
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

	tpl = views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "upload-file.gohtml"))
	r.Get("/upload", controllers.StaticHandler(tpl))
	r.Post("/upload/new", handleUpload)
	// controller
	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "signup.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	r.Get("/gallery/{userID}", galleryHandler)

	// hanlde unknown routes
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// start server
	fmt.Println("Server starting on :3000")
	http.ListenAndServe(":3000", r)
}
