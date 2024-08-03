package cache

import (
	"sync"
	"time"
)

type cache[K comparable, V any] struct {
	mu         sync.RWMutex
	storage    map[K]cacheItem[V]
	expiration time.Duration
}

type cacheItem[V any] struct {
	value     V
	expiredAt int64
}

func New[K comparable, V any](keyExpiration time.Duration) *cache[K, V] {
	return &cache[K, V]{storage: make(map[K]cacheItem[V])}
}

func (c *cache[K, V]) Set(key K, value V) {
	expiredAt := time.Now().Add(c.expiration).Unix()
	c.mu.Lock()
	c.storage[key] = cacheItem[V]{
		value:     value,
		expiredAt: expiredAt,
	}
	c.mu.Unlock()
}

func (c *cache[K, V]) Get(key K) (value V, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.storage[key]
	if !ok || val.isExpired() {
		return *new(V), false
	}

	return val.value, true
}

func (i cacheItem[V]) isExpired() bool {
	if time.Now().Unix() > i.expiredAt {
		return true
	}
	return false
}
