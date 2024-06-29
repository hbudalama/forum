package db

import (
	"database/sql"

	"learn.reboot01.com/git/hbudalam/forum/pkg/structs"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Connect() error {
	database, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}
	db = database
	return nil
}

func Close() error {
	return db.Close()
}

func GetUser(userId int) (structs.User, error) {
	panic("not implemented")
}

func GetUserPassword(username string) (string, error) {
	panic("not implemented")
}

func GetPost(postId int) (structs.Post, error) {
	panic("not implemented")
}

func GetComment(commentId int) (structs.Comment, error) {
	panic("not implemented")
}

func GetSession(token string) (structs.Session, error) {
	panic("not implemented")
}
