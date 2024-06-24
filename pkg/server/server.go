package server

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if !MethodsGuard(w, r, "POST") {
		return
	}
	if LoginGuard( r) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w,"Already logged in!")
		return
	}

}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {

}

func AddPostsHandler(w http.ResponseWriter, r *http.Request) {

}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {

}

func AddDislikesHandler(w http.ResponseWriter, r *http.Request) {

}

func AddLikesHandler(w http.ResponseWriter, r *http.Request) {

}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {

}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if LoginGuard(r) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w,"Already logged in!")
		return
	}
	if !MethodsGuard(w, r, "GET", "POST") {
		return
	}

	http.ServeFile(w, r, filepath.Join("pages", "login.html"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if !MethodsGuard(w, r, "GET") {
		return
	}

	http.ServeFile(w, r, filepath.Join("pages", "index.html"))
}
