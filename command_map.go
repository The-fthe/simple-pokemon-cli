package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	URL      string
	Next     string
	Previous string
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type ResultData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Result
}

func (c *Config) createNewConfig() error {
	c.URL = "https://pokeapi.co/api/v2/location-area"

	data, err := readBody(c.URL)
	if err != nil {
		return fmt.Errorf("read body error: {%s}", err.Error())
	}

	c.Next = data.Next
	c.Previous = data.Previous
	return nil
}

func (c *Config) map_next() error {
	//if next is empty print current url result location-area
	fmt.Printf("start c url: %s \n", c.URL)
	if c.URL == "" {
		err := c.createNewConfig()
		if err != nil {
			return err
		}

		data, err := readBody(c.URL)
		if err != nil {
			return err
		}

		c.Next = data.Next
		c.Previous = data.Previous

		mapNames := c.GetResultMapNames(data)
		for _, mapName := range mapNames {
			fmt.Println(mapName)
		}
		return nil
	}
	fmt.Printf("next available c next: %s \n", c.Next)

	//if next exist print next url location-area
	data, err := readBody(c.Next)
	if err != nil {
		return err
	}

	c.Next = data.Next
	c.Previous = data.Previous

	mapNames := c.GetResultMapNames(data)
	for _, mapName := range mapNames {
		fmt.Println(mapName)
	}

	return nil
}

func (c *Config) map_previous() error {
	if c.URL == "" || c.Previous == "" {
		err := c.createNewConfig()
		if err != nil {
			return err
		}

		data, err := readBody(c.URL)
		if err != nil {
			return err
		}

		c.Next = data.Next
		c.Previous = data.Previous

		mapNames := c.GetResultMapNames(data)
		for _, mapName := range mapNames {
			fmt.Println(mapName)
		}
		return nil
	}

	//if previous exist print next url location-area
	data, err := readBody(c.Previous)
	if err != nil {
		return err
	}

	c.Next = data.Next
	c.Previous = data.Previous

	mapNames := c.GetResultMapNames(data)
	for _, mapName := range mapNames {
		fmt.Println(mapName)
	}

	return nil
}

func (c *Config) GetResultMapNames(data ResultData) []string {
	var mapNames []string
	for _, m := range data.Results {
		mapNames = append(mapNames, m.Name)
	}
	return mapNames
}

func readBody(url string) (ResultData, error) {
	res, err := http.Get(url)
	if err != nil {
		return ResultData{}, errors.New("get url error")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResultData{}, errors.New("read errors")
	}

	if res.StatusCode > 299 {
		return ResultData{}, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}

	res.Body.Close()

	var data ResultData

	err = json.Unmarshal(body, &data)
	if err != nil {
		return ResultData{}, fmt.Errorf("Unmarshal error")
	}
	return data, nil
}
