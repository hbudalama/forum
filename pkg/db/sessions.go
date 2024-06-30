package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"learn.reboot01.com/git/hbudalam/forum/pkg/structs"
)

func CreateSession(username string) (string, error) {
	token := uuid.New().String()
	expiry := time.Now().Add(24 * time.Hour)

	_, err := db.Exec("UPDATE User SET sessionToken = ?, sessionExpiration = ? WHERE username = ?", token, expiry, username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetSession(token string) (*structs.Session, error) {
	var session structs.Session
	err := db.QueryRow("SELECT sessionToken, sessionExpiration, username FROM User WHERE sessionToken = ?", token).Scan(&session.Token, &session.Expiry, &session.User.Username)
	if err != nil {
		log.Printf("GetSession: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &session, nil
}

func DeleteSession(token string) error {
	_, err := db.Exec("UPDATE User SET sessionToken = NULL, sessionExpiration = NULL WHERE sessionToken = ?", token)
	return err
}
