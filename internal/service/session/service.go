package session

import (
	"time"

	"github.com/arxon31/gophkeep/pkg/cache"
)

const expiration = time.Hour

//go:generate moq -out storage_moq_test.go . sessionStorage
type sessionStorage interface {
	Set(key string, value any)
	Get(key string) (value any, exists bool)
	Delete(key string)
}

type sessionService struct {
	cache sessionStorage
}

func NewService() *sessionService {
	return &sessionService{cache: cache.New[string, any](expiration)}
}
