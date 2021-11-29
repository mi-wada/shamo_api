package session_store

import (
	"api/infrastructure/database"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

const (
	SESSION_TABLE_NAME = "shamo_session"
	SESSION_TIMEOUT    = 600 // 10minutes for test
)

type SessionStore struct {
	database *dynamo.DB
}

func NewSessionStore() SessionStore {
	return SessionStore{
		database: database.Connect(),
	}
}

func (s SessionStore) GetUserId(sessionId string) (string, error) {
	table := s.database.Table(SESSION_TABLE_NAME)

	session := new(Session)
	err := table.Get("id", sessionId).One(&session)

	if err != nil {
		return "", err
	}

	return session.userId, nil
}

func (s SessionStore) Set(userId string) string {
	table := s.database.Table(SESSION_TABLE_NAME)

	sessionId := uuid.NewString()
	session := Session{
		id:     sessionId,
		userId: userId,
		ttl:    int(time.Now().Unix()) + SESSION_TIMEOUT,
	}

	table.Put(session).Run()

	return sessionId
}

type Session struct {
	id     string `dynamo:"id"`
	userId string `dynamo:"user_id"`
	ttl    int    `dynamo:"ttl"`
}
