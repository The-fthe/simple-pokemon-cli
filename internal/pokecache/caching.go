package pokecache

import (
	"fmt"
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheTable: make(map[string]CacheEntry),
		mu:         &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(url string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheTable[url] = CacheEntry{
		createAt: time.Now().UTC(),
		val:      data,
	}
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheTable[url]
	return entry.val, ok
}

// remote cacheEntry that is older than interval
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for url, entry := range c.cacheTable {
		if entry.createAt.Before(now.Add(-interval)) {
			delete(c.cacheTable, url)
			fmt.Printf("\ncache is being deleted\n")
		}
	}

}
