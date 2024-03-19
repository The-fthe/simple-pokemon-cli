package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocation(pageURL *string) (RespLocations, error) {
	url := BASE_URL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	if dat, ok := c.cache.Get(url); ok {
		locationResp := RespLocations{}
		err := json.Unmarshal(dat, &locationResp)
		if err != nil {
			return RespLocations{}, err
		}
		return locationResp, nil
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

	locationResp := RespLocations{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil

}

func (c *Client) ExploreLocation(location string) (RespExp, error) {
	url := BASE_URL + "/location-area" + "/" + location
	if dat, ok := c.cache.Get(url); ok {
		RespExp := RespExp{}
		err := json.Unmarshal(dat, &RespExp)
		if err != nil {
			return RespExp, err
		}
		return RespExp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExp{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExp{}, err
	}

	// fmt.Println(url)
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExp{}, err
	}

	respExp := RespExp{}
	err = json.Unmarshal(dat, &respExp)
	if err != nil {
		return RespExp{}, err
	}

	c.cache.Add(url, dat)
	return respExp, nil

}
