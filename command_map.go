package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

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

func map_next() error {
	return errors.New("")
}

func map_previous() error {
	return errors.New("")
}

func GetResultMapNames(data ResultData) []string {
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
