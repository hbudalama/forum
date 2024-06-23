package main

import (
	"database/sql"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	database, dbErr := sql.Open("sqlite3", "./database.db")

	if dbErr != nil {
		log.Fatal("Error opening database:", dbErr)
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/login", loginHandler)
	// mux.HandleFunc("/api/posts/{id}/comments", )
	// mux.HandleFunc("/api/posts", )
	// mux.HandleFunc("/api/posts/{id}", )
	// mux.HandleFunc("/api/posts/{id}/dislike", )
	// mux.HandleFunc("/api/posts/{id}/like", )
	// mux.HandleFunc("/api/categories", )
	// mux.HandleFunc("signup", )
	mux.HandleFunc("/", homeHandler)
	// http://localhost:8080/login
	log.Println("Serving on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("Error starting server:", err)
	}
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, filepath.Join("pages", "login.html"))
		return
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, filepath.Join("pages", "index.html"))
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
