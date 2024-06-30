package db

import (
	"errors"
	"log"
	"strings"

	"github.com/mattn/go-sqlite3"
	"learn.reboot01.com/git/hbudalam/forum/pkg/structs"
)

func AddUser(username string, email string, hashedPassword string) (*structs.User, error) {
	// Insert user into the database
	dbMutex.Lock()
	defer dbMutex.Unlock()
	_, err := db.Exec("INSERT INTO user (username, email, password) VALUES ($1, $2, $3)", username, email, hashedPassword)
	if err != nil {
		log.Printf("AddUser: %s\n", err.Error())
		sqliteErr, ok := err.(sqlite3.Error)
		if ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			if strings.Contains(sqliteErr.Error(), "username") {
				return nil, errors.New("username already exists")
			} else if strings.Contains(sqliteErr.Error(), "email") {
				return nil, errors.New("email already exists")
			}
		}

		return nil, errors.New("internal server error")
	}

	return nil, nil
}
