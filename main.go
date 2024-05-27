package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gouae/golang.ae/internal/template"
)

func main() {
	port := "8080"

	if value, ok := os.LookupEnv("PORT"); ok {
		port = value
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		template.Home("Under Construction").Render(context.TODO(), w)
	})

	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/favicon.ico")
	})

	mux.HandleFunc("GET /public/", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path[len("/public/"):]
		fullPath := filepath.Join(".", "public", filePath)
		http.ServeFile(w, r, fullPath)
	})

	mux.HandleFunc("GET /discord", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://discord.gg/FXhjYCkvGg", 301)
	})

	mux.HandleFunc("GET /whatsapp", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://chat.whatsapp.com/DEmS5AmfJSfBmH40aYLih1", 301)
	})

	log.Printf("Server is running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
