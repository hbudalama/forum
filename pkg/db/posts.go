package db

import (
	"errors"

	"learn.reboot01.com/git/hbudalam/forum/pkg/structs"
)

// this function will be reused in the functions below
func postExists(id int) bool {
	var status bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM post WHERE ID = $1)", id).Scan(&status)

	if err != nil {
		return false
	}

	return status
}

func userExists(username string) bool {
	var status bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM user WHERE username = $1)", username).Scan(&status)

	if err != nil {
		return false
	}

	return status
}

func isOwner(post int, username string) bool {
	if username == "" || post < 1 {
		return false
	}

	if !postExists(post) || !userExists(username) {
		return false
	}

	var status bool

	//TO DO: db.QueryRow

	return status
}

func CreatePost(title string, content string, username string) error {
	_, err := db.Exec("INSERT INTO post (Title, Content, username) VALUES ($1, $2, $3)", title, content, username)

	if err != nil {
		return err
	}

	return nil
}

func DeletePost(id int, user string) error {
	if !postExists(id) || !isOwner(id, user) {
		return errors.New("post does not exist")
	}
	_, err := db.Exec("DELETE FROM post WHERE ID = $1", id)

	if err != nil {
		return err
	}

	return nil
}

func Interact(post int, username string, interaction int) error {
	if !postExists(post) || !userExists(username) {
		return errors.New("post does not exist")
	}

	//TO DO: Check if the user didn't already interact with the post

	_, err := db.Exec("INSERT INTO interaction (PostID, Username, Interaction) VALUES ($1, $2, $3)", post, username, interaction)

	if err != nil {
		return err
	}

	return nil
}

func GetPost(postID int) (structs.Post, error) {
	var post structs.Post

	row := db.QueryRow("SELECT PostID, Title, Content, CreatedDate, username FROM Post WHERE PostID = $1", postID)
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedDate, &post.Username)
	if err != nil {
		return post, err
	}

	rows, err := db.Query("SELECT CategoryName FROM Category c INNER JOIN PostCategory pc ON c.CategoryID = pc.CategoryID WHERE pc.PostID = $1", postID)
	if err != nil {
		return post, err
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return post, err
		}
		post.Categories = append(post.Categories, category)
	}

	return post, nil
}
