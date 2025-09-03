package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	val     []byte
	created time.Time
}

type Cache struct {
	cache    map[string]CacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	return &Cache{
		cache:    make(map[string]CacheEntry),
		interval: interval,
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = CacheEntry{
		val:     val,
		created: time.Now(),
	}
	go c.ReapLoop(c.interval)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	cacheEntry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return cacheEntry.val, true
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.cache, key)
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mutex.Lock()
		for key, cacheEntry := range c.cache {
			if time.Since(cacheEntry.created) > interval {
				delete(c.cache, key)
			}
		}
		c.mutex.Unlock()
	}
}
