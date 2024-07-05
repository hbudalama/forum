package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"learn.reboot01.com/git/hbudalam/forum/pkg/db"
	"learn.reboot01.com/git/hbudalam/forum/pkg/structs"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if !MethodsGuard(w, r, "POST") {
		return
	}
	// username := r.FormValue("username")
	// email := r.FormValue("email")
	// password := r.FormValue("password")
	var requestData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, `{"reason": "Invalid request"}`, http.StatusBadRequest)
		return
	}
	username := requestData.Username
	email := requestData.Email
	password := requestData.Password
	if strings.TrimSpace(username) == "" {
		http.Error(w, `{"reason": "User ID is required"}`, http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(email) == "" {
		http.Error(w, `{"reason": "Email is required"}`, http.StatusBadRequest)
		return
	}
	if !validEmail(email) {
		http.Error(w, `{"reason": "Invalid email format"}`, http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(password) == "" {
		http.Error(w, `{"reason": "Password is required"}`, http.StatusBadRequest)
		return
	}
	if !validatePassword(password) {
		http.Error(w, `{"reason": "Password must be at least 8 characters long "}`, http.StatusBadRequest)
		return
	}
	exists, err := db.CheckUsernameExists(username)
	if err != nil {
		Error500Handler(w, r)
		return
	}
	if exists {
		http.Error(w, `{"reason": "Username already taken"}`, http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		Error500Handler(w, r)
		return
	}
	_, err = db.AddUser(username, email, string(hashedPassword))
	if err != nil {
		errMsg := fmt.Sprintf(`{"reason": "%s"}`, err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		println(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func AddLikesHandler(w http.ResponseWriter, r *http.Request) {
	if !LoginGuard(w, r) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	postIDStr := r.PathValue("id")
	fmt.Printf("postIDStr: %s\n", postIDStr)
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

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
	username := session.User.Username

	err = db.InsertOrUpdateInteraction(postID, username, 1)
	if err != nil {
		Error500Handler(w, r)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func AddDislikesHandler(w http.ResponseWriter, r *http.Request) {
	if !LoginGuard(w, r) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	postIDStr := r.PathValue("id")
	fmt.Printf("postIDStr: %s\n", postIDStr)
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

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
	username := session.User.Username

	err = db.InsertOrUpdateInteraction(postID, username, 0)
	if err != nil {
		Error500Handler(w, r)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if !MethodsGuard(w, r, "GET", "POST") {
		return
	}
	if r.Method == "POST" {
		var requestData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, `{"reason": "Invalid request"}`, http.StatusBadRequest)
			return
		}
		username := strings.TrimSpace(requestData.Username)
		password := strings.TrimSpace(requestData.Password)
		if username == "" {
			http.Error(w, `{"reason": "Username is required"}`, http.StatusBadRequest)
			return
		}
		if password == "" {
			http.Error(w, `{"reason": "Password is required"}`, http.StatusBadRequest)
			return
		}
		exists, err := db.CheckUsernameExists(username)
		if err != nil {
			log.Printf("LoginHandler: Error checking username: %s\n", err.Error())
			Error500Handler(w, r)
			return
		}
		if !exists {
			Error404Handler(w, r)
			return

		}
		passwordMatches, err := db.CheckPassword(username, password)
		if err != nil {
			log.Printf("LoginHandler: Error checking password: %s\n", err.Error())
			Error500Handler(w, r)
			return
		}
		if !passwordMatches {
			http.Error(w, `{"reason": "Invalid password"}`, http.StatusUnauthorized)
			return
		}
		token, err := db.CreateSession(username)
		if err != nil {
			log.Printf("LoginHandler: Error creating session: %s\n", err.Error())
			Error500Handler(w, r)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})
	}
	http.ServeFile(w, r, filepath.Join("pages", "login.html"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// msg := fmt.Sprintf("Page Not Found, Did you forget to implement: '%s' handler?\n", r.URL.Path)
		// http.Error(w, msg, http.StatusNotFound)
		Error404Handler(w, r)
		return

	}

	if !MethodsGuard(w, r, "GET") {
		http.Error(w, "Method Not Allowed HomeHandler", http.StatusMethodNotAllowed)
		return
	}
	ctx := structs.HomeContext{}
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
		fmt.Printf("Fick this shot: %+v\n", session.User)
		ctx.LoggedInUser = &session.User
	}

	ctx.Posts = db.GetAllPosts()

	tmpl, err := template.ParseFiles(filepath.Join("pages", "index.html"))
	if err != nil {
		log.Printf("can't parse the template: %s\n", err.Error())
		Error500Handler(w, r)
		return
	}

	err = tmpl.Execute(w, ctx)
	if err != nil {
		log.Printf("can't execute the template: %s\n", err.Error())
		Error500Handler(w, r)
		return
	}
}

func validEmail(email string) bool {
	return strings.Contains(email, "@") && strings.HasSuffix(email, ".com")
}

func validatePassword(password string) bool {
	return len(password) >= 8
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if !MethodsGuard(w, r, "DELETE") {
		http.Error(w, "only DELETE requests allowed", http.StatusMethodNotAllowed)
		return
	}
	if !LoginGuard(w, r) {
		http.Error(w, "You have to be logged in", http.StatusUnauthorized)
		return
	}
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No session token found", http.StatusUnauthorized)
			return
		}
		log.Printf("can't get the cookie: %s\n", err.Error())
		Error500Handler(w, r)
		return
	}
	err = db.DeleteSession(cookie.Value)
	if err != nil {
		log.Printf("LogoutHandler: %s", err.Error())
		Error500Handler(w, r)
		return
	}
	// Clear the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		// MaxAge: -1,
		Path: "/",
	})
	w.WriteHeader(http.StatusOK)
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.URL.Path[len("/posts/"):]
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := db.GetPost(postID)
	if err != nil {
		Error404Handler(w, r)
		return
	}

	comments, err := db.GetComments(postID)
	if err != nil {
		Error500Handler(w, r)
		return
	}
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
	user := session.User

	data := struct {
		Post         structs.Post
		Comments     []structs.Comment
		LoggedInUser *structs.User
	}{
		Post:         post,
		Comments:     comments,
		LoggedInUser: &user,
	}

	tmpl, err := template.ParseFiles("templates/posts.html")
	if err != nil {
		Error500Handler(w, r)
		return
	}

	tmpl.Execute(w, data)
}

// func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
// 	postIDStr := r.URL.Path[len("/api/posts/") : len(r.URL.Path)-len("/comments")]
// 	postID, err := strconv.Atoi(postIDStr)
// 	if err != nil {
// 		http.Error(w, "Invalid post ID", http.StatusBadRequest)
// 		return
// 	}

// 	cookie, err := r.Cookie("session_token")
// 	if err != nil {
// 		log.Printf("can't get the cookie: %s\n", err.Error())
// 		return
// 	}

// 	token := cookie.Value

// 	session, err := db.GetSession(token)
// 	if err != nil {
// 		log.Printf("can't get the session: %s\n", err.Error())
// 		return
// 	}
// 	username := session.User.Username

// 	comment := r.FormValue("comment")
// 	if strings.TrimSpace(comment) == "" {
// 		http.Error(w, "Comment cannot be empty", http.StatusBadRequest)
// 		return
// 	}

// 	err = db.AddComment(postID, username, comment)
// 	if err != nil {
// 		http.Error(w, "Error adding comment", http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, "/posts/"+strconv.Itoa(postID), http.StatusSeeOther)
// }

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	if !LoginGuard(w, r) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	if r.Method == http.MethodPost {
		postIDStr := r.URL.Path[len("/api/posts/") : len(r.URL.Path)-len("/comments")]
		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		comment := r.FormValue("comment")
		if strings.TrimSpace(comment) == "" {
			http.Error(w, "Comment cannot be empty", http.StatusBadRequest)
			return
		}
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

		username := session.User.Username

		err = db.AddComment(postID, username, comment)
		if err != nil {
			Error500Handler(w, r)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/posts/%d", postID), http.StatusSeeOther)
	}
}

func AddPostsHandler(w http.ResponseWriter, r *http.Request) {
	if !LoginGuard(w, r) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	if strings.TrimSpace(title) == "" || strings.TrimSpace(content) == "" {
		http.Error(w, "The post cant be created without a title or content", http.StatusBadRequest)
		return
	}
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
		Error500Handler(w, r)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	postIDStr := r.URL.Path[len("/posts/"):]
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := db.GetPost(postID)
	if err != nil {
		Error404Handler(w, r)
		return
	}

	comments, err := db.GetComments(postID)
	if err != nil {
		Error500Handler(w, r)
		return
	}

	var user *structs.User
	if LoginGuard(w, r) {
		cookie, err := r.Cookie("session_token")
		if err == nil {
			token := cookie.Value
			session, err := db.GetSession(token)
			if err == nil {
				user = &session.User
			}
		}
	}

	ctx := struct {
		Post         structs.Post
		Comments     []structs.Comment
		LoggedInUser *structs.User
	}{
		Post:         post,
		Comments:     comments,
		LoggedInUser: user,
	}

	tmpl, err := template.ParseFiles(filepath.Join("pages", "posts.html"))
	if err != nil {
		log.Printf("can't parse the template: %s\n", err.Error())
		Error500Handler(w, r)
		return
	}

	err = tmpl.Execute(w, ctx)
	if err != nil {
		log.Printf("can't execute the template: %s\n", err.Error())
		Error500Handler(w, r)
		return
	}
}

func Error404Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles(filepath.Join("pages", "error404.html"))
	if err != nil {
		log.Printf("can't parse the template: %s\n", err.Error())
		Error500Handler(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("can't execute the template: %s\n", err.Error())
		Error500Handler(w, r)
	}
}


func Error500Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles(filepath.Join("pages", "error500.html"))
	if err != nil {
		log.Printf("can't parse the template: %s\n", err.Error())
		Error500Handler(w, r)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("can't execute the template: %s\n", err.Error())
		http.Error(w, "Internal Server Error Error500Handler", http.StatusInternalServerError)
	}
}
