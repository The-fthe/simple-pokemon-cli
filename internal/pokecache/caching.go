package pokecache

import (
	"fmt"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheTable: make(map[string]CacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(url string, data []byte) {
	c.Lock()
	defer c.Unlock()

	entry := CacheEntry{
		createAt: time.Now(),
		val:      data,
	}

	c.cacheTable[url] = entry
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()

	entry, ok := c.cacheTable[url]
	if !ok {
		fmt.Printf("URL not found in cache")
		return nil, false
	}

	return entry.val, true
}

// remote cacheEntry that is older than interval
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.Lock()
		for url, entry := range c.cacheTable {
			if time.Since(entry.createAt) > interval {
				delete(c.cacheTable, url)
				fmt.Printf("cache is being deleted\n")
			}
		}
		c.Unlock()
	}
}
