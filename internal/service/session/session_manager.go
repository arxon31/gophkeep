package session

import (
	"time"

	"github.com/arxon31/gophkeep/pkg/cache"
	"github.com/google/uuid"
)

const expiration = time.Hour

type sessionStorage interface {
	Set(key string, value struct{})
	Get(key string) (value struct{}, exists bool)
}

type sessionManager struct {
	cache sessionStorage
}

func NewManager() *sessionManager {
	return &sessionManager{cache: cache.New[string, struct{}](expiration)}
}

func (s *sessionManager) Create() (sessionID string) {
	sessionID = uuid.New().String()

	s.cache.Set(sessionID, struct{}{})

	return sessionID
}

func (s *sessionManager) IsExists(sessionID string) bool {
	_, ok := s.cache.Get(sessionID)
	return ok
}
