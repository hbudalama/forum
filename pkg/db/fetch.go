package db

import (
	"learn.reboot01.com/git/hbudalam/forum/pkg/structs"
)

func GetAllPosts() []structs.Post {
	var posts []structs.Post

	err := db.QueryRow("SELECT * FROM posts").Scan(&posts)

	if err != nil {
		return []structs.Post{}
	}

	return posts
}

func GetFilteredPosts(category string) []structs.Post {
	var filteredPosts []structs.Post
	err := db.QueryRow("SELECT * FROM posts WHERE category = $1", category).Scan(&filteredPosts)

	if err != nil {
		return []structs.Post{}
	}
	return filteredPosts
}

func GetPostDetails(postId int) (structs.Post, structs.User, []structs.Comment, []structs.Interaction) {
	var (
		thisPost          structs.Post
		thisUser          structs.User
		theseComments     []structs.Comment
		theseInteractions []structs.Interaction
	)

	err := db.QueryRow("SELECT * FROM posts WHERE id = $1", postId).Scan(&thisPost.ID, &thisPost.Title, &thisPost.Content, &thisPost.Categories)

	if err != nil {
		return structs.Post{}, structs.User{}, []structs.Comment{}, []structs.Interaction{}
	}

	return thisPost, thisUser, theseComments, theseInteractions
}
