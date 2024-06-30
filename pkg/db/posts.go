package db

func CreatePost(title string, content string, username string) error {
	_, err := db.Exec("INSERT INTO posts (Title, Content, username) VALUES ($1, $2, $3)", title, content, username)

	if err != nil {
		return err
	}

	return nil
}

func DeletePost(id int) error {
	//TO DO: Check if the post exists
	//TO DO: Check if the user is the owner of the post
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
