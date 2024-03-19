package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	url := BASE_URL + "/" + "pokemon" + "/" + pokemon
	if dat, ok := c.cache.Get(url); ok {
		RespPoke := Pokemon{}
		err := json.Unmarshal(dat, &RespPoke)
		if err != nil {
			return Pokemon{}, err
		}
		return RespPoke, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	fmt.Println(url)
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	respPoke := Pokemon{}
	err = json.Unmarshal(dat, &respPoke)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, dat)
	return respPoke, nil
}
