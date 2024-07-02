package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"learn.reboot01.com/git/hbudalam/forum/pkg/db"
	"learn.reboot01.com/git/hbudalam/forum/pkg/structs"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if !MethodsGuard(w, r, "POST") {
		http.Error(w, "Method Not Allowed SignupHandler", http.StatusMethodNotAllowed)
		return
	}
	// if LoginGuard(r) {
	//  http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	//  return
	// }ß

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if strings.TrimSpace(username) == "" {
		http.Error(w, `{"error": "User ID is required"}`, http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(email) == "" {
		http.Error(w, `{"error": "Email is required"}`, http.StatusBadRequest)
		return
	}
	if !validEmail(email) {
		http.Error(w, `{"error": "Invalid email format"}`, http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(password) == "" {
		http.Error(w, `{"error": "Password is required"}`, http.StatusBadRequest)
		return
	}

	if !validatePassword(password) {
		http.Error(w, `{"error": "Password must be at least 8 characters long "}`, http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}
	_, err = db.AddUser(username, email, string(hashedPassword))
	if err != nil {
		errMsg := fmt.Sprintf(`{"error": "%s"}`, err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func CommentsHandler(w http.ResponseWriter, r *http.Request) {

}

func AddPostsHandler(w http.ResponseWriter, r *http.Request) {
	if !LoginGuard(w, r) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	var user structs.User
	cookie, err := r.Cookie("session_token")
	if err != nil {
		log.Printf("can't get the cookie: %s\n", err.Error())
		return
	}

	token := cookie.Value

	session, err := db.GetSession(token)
	if err != nil {
		log.Printf("can't get the session: %s\n", err.Error())
		return
	}
	user = session.User

	err = db.CreatePost(title, content, user.Username)
	if err != nil {
		log.Printf("failed to create post: %s\n", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {

}

func AddDislikesHandler(w http.ResponseWriter, r *http.Request) {
	if !LoginGuard(w, r) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
}

func AddLikesHandler(w http.ResponseWriter, r *http.Request) {
	if !LoginGuard(w, r) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// if LoginGuard(r) {
	//  http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	//  fmt.Fprint(w, "Already logged in!")
	//  return

	// }
	if !MethodsGuard(w, r, "GET", "POST") {
		http.Error(w, "Method Not Allowed LoginHandler", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		exists, err := db.CheckUsernameExists(username)
		if err != nil {
			log.Printf("LoginHandler: Error checking username: %s\n", err.Error())
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		if !exists {
			http.Error(w, "Username not found", http.StatusUnauthorized)
			return
		}

		passwordMatches, err := db.CheckPassword(username, password)
		if err != nil {
			log.Printf("LoginHandler: Error checking password: %s\n", err.Error())
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		if !passwordMatches {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		token, err := db.CreateSession(username)
		if err != nil {
			log.Printf("LoginHandler: Error creating session: %s\n", err.Error())
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, filepath.Join("pages", "login.html"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		msg := fmt.Sprintf("Page Not Found, Did you forget to implement: '%s' handler?\n", r.URL.Path)
		http.Error(w, msg, http.StatusNotFound)
	}

	if !MethodsGuard(w, r, "GET") {
		http.Error(w, "Method Not Allowed HomeHandler", http.StatusMethodNotAllowed)
		return
	}

	var user structs.User
	if LoginGuard(w, r) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			log.Printf("can't get the cookie: %s\n", err.Error())
			return
		}

		token := cookie.Value

		session, err := db.GetSession(token)
		if err != nil {
			log.Printf("can't get the session: %s\n", err.Error())
			return
		}
		user = session.User
	}

	posts := db.GetAllPosts()
	ctx := structs.HomeContext{LoggedInUser: &user, Posts: posts}

	tmpl, err := template.ParseFiles(filepath.Join("pages", "index.html"))
	if err != nil {
		log.Printf("can't parse the template: %s\n", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, ctx)
	if err != nil {
		log.Printf("can't execute the template: %s\n", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func validEmail(email string) bool {
	return strings.Contains(email, "@") && strings.HasSuffix(email, ".com")
}

func validatePassword(password string) bool {
	return len(password) >= 8
}
