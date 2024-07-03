package structs

import "time"

type User struct {
	// ID int
	Username string
	// FirstName string
	// LastName string
	Email string
}

type Post struct {
	ID           int
	Title        string
	Content      string
	CreatedDate  time.Time
	Username string
	Categories   []string
	Interactions []Interaction
}

type Comment struct {
	ID           int
	Content      string
	CreatedDate  time.Time
	PostID       int
	UserID       int
	Interactions []Interaction
}

type Interaction struct {
	UserId int
	Kind   int
}

type Session struct {
	Token  string
	Expiry time.Time
	User   User
}

type HomeContext struct {
	LoggedInUser *User
	Posts        []Post
}

type PostContext struct {
	LoggedInUser *User
	Categories   []string
	Post         Post
	Comments     []Comment
}
