package db

import (
	"database/sql"
	"log"

	"learn.reboot01.com/git/hbudalam/forum/pkg/structs"
)

func GetAllPosts() []structs.Post {
	var posts []structs.Post
	dbMutex.Lock()
	defer dbMutex.Unlock()
	rows, err := db.Query(`
	SELECT 
	p.PostID, p.Title, p.Content, p.Username, 
	IFNULL(likes.likes, 0) as likes, 
	IFNULL(dislikes.dislikes, 0) as dislikes 
	FROM post p
	LEFT JOIN (SELECT PostID, COUNT(*) as likes FROM interaction WHERE Kind = 1 GROUP BY PostID) likes 
	ON p.PostID = likes.PostID
	LEFT JOIN (SELECT PostID, COUNT(*) as dislikes FROM interaction WHERE Kind = 0 GROUP BY PostID) dislikes 
	ON p.PostID = dislikes.PostID ORDER BY "CreatedDate" DESC 
    `)
	if err != nil {
		log.Printf("Query error: %s", err)
		return []structs.Post{}
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Username, &post.Likes, &post.Dislikes)
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
	// fmt.Printf("%+v\n", posts)
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

func GetCommentDetails(commentId int) (structs.Comment, structs.User, []structs.CommentInteraction) {
	var (
		thisComment              structs.Comment
		thisUser                 structs.User
		theseCommentInteractions []structs.CommentInteraction
	)

	err := db.QueryRow("SELECT * FROM comment WHERE id = $1", commentId).Scan(&thisComment.ID, &thisComment.Content, &thisComment.PostID)

	if err != nil {
		return structs.Comment{}, structs.User{}, []structs.CommentInteraction{}
	}

	return thisComment, thisUser, theseCommentInteractions
}

func InsertOrUpdateInteraction(postID int, username string, kind int) error {
	var existingKind int
	err := db.QueryRow("SELECT Kind FROM Interaction WHERE PostID = ? AND Username = ?", postID, username).Scan(&existingKind)
	if err != nil {
		if err == sql.ErrNoRows {
			// No existing interaction, insert a new one
			_, err = db.Exec(
				"INSERT INTO Interaction (PostID, Username, Kind) VALUES (?, ?, ?)",
				postID, username, kind,
			)
			if err != nil {
				log.Printf("InsertInteraction error: %s", err)
				return err
			}
		} else {
			// Some other error occurred
			log.Printf("Query error: %s", err)
			return err
		}
	} else {
		// Existing interaction found, update it if necessary
		if existingKind != kind {
			_, err = db.Exec(
				"UPDATE Interaction SET Kind = ? WHERE PostID = ? AND Username = ?",
				kind, postID, username,
			)
			if err != nil {
				log.Printf("UpdateInteraction error: %s", err)
				return err
			}
		}
	}
	return nil
}
