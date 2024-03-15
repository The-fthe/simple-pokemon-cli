package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/The-fthe/pokedex/internal/pokecache"
)

func (c *Client) ListLocation(pageURL *string, cache *pokecache.Cache) (RespLocations, error) {
	url := BASE_URL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, err
	}
	// c.cache := pokecache.NewCache(5 * time.Second)
	cache.Add(url, dat)

	locationResp := RespLocations{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocations{}, err
	}

	return locationResp, nil

}
