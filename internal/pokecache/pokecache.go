package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	entry		map[string]cacheEntry
	mutex		sync.Mutex
	interval	time.Duration
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache {
		entry:		make(map[string]cacheEntry),
		interval:	interval,
	}
	go newCache.reapLoop()
	return newCache
}

func newCacheEntry(val []byte) cacheEntry {
	return cacheEntry {
		createdAt:	time.Now(),
		val:		val,
	}
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.entry[key] = newCacheEntry(val)
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	entry, ok := cache.entry[key]
	if ok {
		return entry.val, ok
	}
	return []byte{}, ok
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	for range ticker.C {
		cache.mutex.Lock()
		for key, value := range cache.entry {
			if value.createdAt.Before(time.Now().Add(-1 * cache.interval)) {
				delete(cache.entry, key)
			}
		}
		cache.mutex.Unlock()
	}
}
