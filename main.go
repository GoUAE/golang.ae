package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/gouae/golang.ae/internal/template"
)

func main() {
	port := "8080"

	// parse command line arguments for --port flag if no value, use default
	portFlag := flag.String("port", port, "The port to listen on")

	flag.Parse()

	if portFlag != nil {
		port = *portFlag
	}

	safeURL := templ.SafeURL("/")

	navbar := []template.NavbarItem{
		{Name: "Organizers", Link: safeURL},
		{Name: "Meetups", Link: safeURL},
		{Name: "Projects", Link: safeURL},
		{Name: "Gallery", Link: safeURL},
		{Name: "Contact", Link: safeURL},
	}

	socials := []template.Social{
		{Icon: "fa-twitter", Link: templ.SafeURL("/twitter")},
		{Icon: "fa-whatsapp", Link: templ.SafeURL("/whatsapp")},
		{Icon: "fa-linkedin", Link: templ.SafeURL("/linkedin")},
		{Icon: "fa-discord", Link: templ.SafeURL("/discord")},
		{Icon: "fa-github", Link: templ.SafeURL("/github")},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		template.Home("golang.ae", navbar, socials).Render(context.TODO(), w)
	})

	mux.HandleFunc("GET /public/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path[len("/public/"):]
		fullPath := filepath.Join(".", "public", filePath)
		http.ServeFile(w, r, fullPath)
	})

	mux.HandleFunc("GET /discord", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://discord.gg/FXhjYCkvGg", http.StatusMovedPermanently)
	})

	mux.HandleFunc("GET /whatsapp", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://chat.whatsapp.com/DEmS5AmfJSfBmH40aYLih1", http.StatusMovedPermanently)
	})

	mux.HandleFunc("GET /github", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/gouae", http.StatusMovedPermanently)
	})

	mux.HandleFunc("GET /linkedin", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://linkedin.com/company/gouae", http.StatusMovedPermanently)
	})

	mux.HandleFunc("GET /twitter", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://twitter.com/gouae", http.StatusMovedPermanently)
	})

	log.Printf("Server is running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
