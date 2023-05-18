package service

import "sync"

type MapCache struct {
	cache sync.Map
}

func NewMapCache() *MapCache {
	return &MapCache{
		cache: sync.Map{},
	}
}

func (m *MapCache) Get(cacheKey string) ([]byte, bool) {
	v, ok := m.cache.Load(cacheKey)
	if !ok {
		return nil, false
	}

	return v.([]byte), true
}

func (m *MapCache) Set(cacheKey string, value []byte) {
	m.cache.Store(cacheKey, value)
}
