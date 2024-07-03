package db

import (
	"log"

	"learn.reboot01.com/git/hbudalam/forum/pkg/structs"
)

func GetAllPosts() []structs.Post {
    var posts []structs.Post

    rows, err := db.Query("SELECT Title, Content, Username FROM post")
    if err != nil {
        log.Printf("Query error: %s", err)
        return []structs.Post{}
    }
    defer rows.Close()

    for rows.Next() {
        var post structs.Post
        err := rows.Scan( &post.Title, &post.Content, &post.Username)
        if err != nil {
            log.Printf("Scan error: %s", err)
            continue
        }
        posts = append(posts, post)
    }

    if err := rows.Err(); err != nil {
        log.Printf("Rows error: %s", err)
        return []structs.Post{}
    }

    return posts
}


func GetFilteredPosts(category string) []structs.Post {
	var filteredPosts []structs.Post
	err := db.QueryRow("SELECT * FROM post WHERE category = $1", category).Scan(&filteredPosts)

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

	err := db.QueryRow("SELECT * FROM post WHERE id = $1", postId).Scan(&thisPost.ID, &thisPost.Title, &thisPost.Content, &thisPost.Categories)

	if err != nil {
		return structs.Post{}, structs.User{}, []structs.Comment{}, []structs.Interaction{}
	}

	return thisPost, thisUser, theseComments, theseInteractions
}
