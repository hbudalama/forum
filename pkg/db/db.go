package db

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

var dbMutex = sync.Mutex{}

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

func CheckUsernameExists(username string) (bool, error) {
	var foundUsername string
	dbMutex.Lock()
	defer dbMutex.Unlock()
	err := db.QueryRow("SELECT username FROM user WHERE username = ?", username).Scan(&foundUsername)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func GetHashedPasswordByUsername(username string) (string, error) {
	var hashedPassword string
	dbMutex.Lock()
	defer dbMutex.Unlock()
	err := db.QueryRow("SELECT password FROM user WHERE username = ?", username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return hashedPassword, nil
}

func CheckPassword(username, password string) (bool, error) {
	hashedPassword, err := GetHashedPasswordByUsername(username)
	if err != nil {
		return false, err
	}
	if hashedPassword == "" {
		return false, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// func GetUser(userId int) (structs.User, error) {
// 	panic("not implemented")
// }

// func GetPost(postId int) (structs.Post, error) {
// 	panic("not implemented")
// }

// func GetComment(commentId int) (structs.Comment, error) {
// 	panic("not implemented")
// }

// func GetSession(token string) (structs.Session, error) {
// 	panic("not implemented")
// }
