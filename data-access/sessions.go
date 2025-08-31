package dataAccess

import (
	"github.com/lemarsupial/go-mailing-list/db"
)

type Session struct {
	Id        int64
	SessionId string
}

func GetSession(sessionId string) (Session, error) {
	stmt, err := db.DbHandle.Prepare("SELECT id, sessionId FROM sessions WHERE sessionId = ?")
	if err != nil {
		return Session{}, err
	}
	defer stmt.Close()

	var foundSession Session
	err = stmt.QueryRow(sessionId).Scan(&foundSession.Id, &foundSession.SessionId)
	if err != nil {
		return Session{}, err
	}

	return foundSession, nil
}

func CreateSession(email string) (Session, error) {
	stmt, err := db.DbHandle.Prepare("INSERT INTO sessions (sessionId) VALUES (?) RETURNING id, sessionId")
	if err != nil {
		return Session{}, err
	}
	defer stmt.Close()

	var createdSession Session
	err = stmt.QueryRow(email).Scan(&createdSession.Id, &createdSession.SessionId)
	if err != nil {
		return Session{}, err
	}

	return createdSession, nil
}

func DeleteAllSessions() error {
	stmt, err := db.DbHandle.Prepare("DELETE FROM sessions")
	if err != nil {
		return err

	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
