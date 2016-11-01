package main

import (
	"errors"
	"net"

	"github.com/satori/go.uuid"
)

// Session represents a terminal session
type Session struct {
	SharePipe *Pipe
	JoinPipe  *Pipe
	ShareConn net.Conn
	JoinConn  net.Conn
}

// NewSession creates a new session object and stores the share connection
func NewSession(sc net.Conn) *Session {
	return &Session{
		SharePipe: NewPipe(),
		JoinPipe:  NewPipe(),
		ShareConn: sc,
	}
}

var store = make(map[string]*Session)

func saveSession(token string, session *Session) {
	store[token] = session
}

func lookup(token string) (*Session, error) {
	s, ok := store[token]

	if !ok {
		return nil, errors.New("session not found")
	}

	return s, nil
}

func generateToken() string {
	return uuid.NewV4().String()
}
