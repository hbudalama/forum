package main

import (
	// "database/sql"
	// "log"
	"log"
	"net/http"
	"os"

	"learn.reboot01.com/git/hbudalam/forum/pkg/db"
	"learn.reboot01.com/git/hbudalam/forum/pkg/server"
)

func main() {
	dbErr := db.Connect()
	if dbErr != nil {
		log.Println("Error connecting to the database:", dbErr)
		os.Exit(1)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/login", server.LoginHandler) // ✅
	mux.HandleFunc("/api/posts/{id}/comments", server.CommentsHandler)
	mux.HandleFunc("/api/posts/{id}", server.GetPostsHandler)
	mux.HandleFunc("/api/posts/{id}/dislike", server.AddDislikesHandler)
	mux.HandleFunc("/api/posts/{id}/like", server.AddLikesHandler)
	mux.HandleFunc("/api/categories", server.GetCategoriesHandler)
	mux.HandleFunc("/signup", server.SignupHandler) // ✅
	mux.HandleFunc("/", server.HomeHandler)
	mux.HandleFunc("/add-post",server.AddPostsHandler)

	log.Println("Serving on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("Error starting server:", err)
		os.Exit(1)
	}
}
