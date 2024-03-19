package pokeapi

import (
	"github.com/The-fthe/pokedex/internal/pokecache"
	"net/http"
	"time"
)

// Client
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) *Client {
	return &Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
