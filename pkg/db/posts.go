package db

import "errors"

// this function will be reused in the functions below
func postExists(id int) bool {
	var status bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM posts WHERE ID = $1)", id).Scan(&status)

	if err != nil {
		return false
	}

	return status
}

func userExists(username string) bool {
	var status bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&status)

	if err != nil {
		return false
	}

	return status
}

func isOwner(post int, username string) bool {
	if !postExists(post) || !userExists(username) {
		return false
	}
	if username == "" || post < 1 {
		return false
	}
	var owner bool

	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM posts WHERE ID = $1)", post).Scan(&owner)

	return false //default
}

func CreatePost(title string, content string, username string) error {
	_, err := db.Exec("INSERT INTO posts (Title, Content, username) VALUES ($1, $2, $3)", title, content, username)

	if err != nil {
		return err
	}

	return nil
}

func DeletePost(id int) error {
	if !postExists(id) {
		return errors.New("post does not exist")
	}
	_, err := db.Exec("DELETE FROM posts WHERE ID = $1", id)

	if err != nil {
		return err
	}

	return nil
}

func Interact(post int, username string, interaction int) error {
	//TO DO: Check if the post exists
	//TO DO: Check if the user didn't already interact with the post
	_, err := db.Exec("INSERT INTO interactions (PostID, Username, Interaction) VALUES ($1, $2, $3)", post, username, interaction)

	if err != nil {
		return err
	}

	return nil

}

/*
type User struct{
	Username string
	Email string
}
type Post struct{
	ID int
	Title string
	Content string
	CreatedDate time.Time
	UserID *User
	Categories []string
	Interactions []Interaction
}
*/
