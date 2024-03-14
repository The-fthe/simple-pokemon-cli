package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	cacheTable map[string]CacheEntry
	sync.Mutex
}
