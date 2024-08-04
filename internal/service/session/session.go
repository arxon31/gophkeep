package session

import "github.com/google/uuid"

func (s *sessionService) Create(info any) (sessionID string) {
	sessionID = uuid.New().String()

	s.cache.Set(sessionID, info)

	return sessionID
}

func (s *sessionService) Info(sessionID string) (any, bool) {
	val, ok := s.cache.Get(sessionID)
	if !ok {
		return nil, false
	}

	return val, true
}
