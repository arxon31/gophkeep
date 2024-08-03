package session

import (
	"time"

	"github.com/arxon31/gophkeep/pkg/cache"
	"github.com/google/uuid"
)

const expiration = time.Hour

type sessionStorage interface {
	Set(key string, value any)
	Get(key string) (value any, exists bool)
}

type sessionManager struct {
	cache sessionStorage
}

type fileInfo struct {
	meta     string
	filename string
	chunks   int64
}

func NewManager() *sessionManager {
	return &sessionManager{cache: cache.New[string, any](expiration)}
}

func (s *sessionManager) CreateSessionForFile(info any) (sessionID string) {
	sessionID = uuid.New().String()

	s.cache.Set(sessionID, info)

	return sessionID
}

func (s *sessionManager) GetInfo(sessionID string) (any, bool) {
	val, ok := s.cache.Get(sessionID)
	if !ok {
		return nil, false
	}

	return val, true
}
