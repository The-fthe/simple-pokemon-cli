package main

import (
	"testing"
	"time"

	"github.com/The-fthe/pokedex/internal/pokeapi"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HellO World ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}

func TestURLBodyParsing(t *testing.T) {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	c := &Config{
		pokeapiClient: pokeClient,
	}
	data, err := c.pokeapiClient.ListLocation(c.nextLocationURL)
	if err != nil {
		t.Errorf("readbody error: %s", err.Error())
	}

	expected := "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"
	if data.Next == nil {
		t.Error("data is error")
	}

	if *data.Next != expected {
		t.Errorf("url don't match Actual :\n '%s' \nvs Expected: \n '%s'", *data.Next, expected)
	}
}

func TestMapNameIsMatch(t *testing.T) {
	expectedNames := []string{
		"canalave-city-area",
		"eterna-city-area",
		"pastoria-city-area",
		"sunyshore-city-area",
		"sinnoh-pokemon-league-area",
		"oreburgh-mine-1f",
		"oreburgh-mine-b1f",
		"valley-windworks-area",
		"eterna-forest-area",
		"fuego-ironworks-area",
		"mt-coronet-1f-route-207",
		"mt-coronet-2f",
		"mt-coronet-3f",
		"mt-coronet-exterior-snowfall",
		"mt-coronet-exterior-blizzard",
		"mt-coronet-4f",
		"mt-coronet-4f-small-room",
		"mt-coronet-5f",
		"mt-coronet-6f",
		"mt-coronet-1f-from-exterior",
	}
	pokeClient := pokeapi.NewClient(5 * time.Second)
	c := &Config{
		pokeapiClient: pokeClient,
	}

	data, err := c.pokeapiClient.ListLocation(c.nextLocationURL)
	if err != nil {
		t.Errorf("readbody error: %s", err.Error())
	}

	if data.Next == nil {
		t.Error("data is error")
	}

	for i, name := range expectedNames {
		if data.Results[i].Name != name {
			t.Errorf("Expected: %s \n Actual: %s", name, data.Results[i].Name)
			return
		}
	}
}

func TestPreviousURL(t *testing.T) {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	c := &Config{
		pokeapiClient: pokeClient,
	}

	data1, err := c.pokeapiClient.ListLocation(c.nextLocationURL)
	if err != nil {
		t.Errorf("readbody error: %s", err.Error())
		return
	}
	c.nextLocationURL = data1.Next
	c.prevLocationURL = data1.Previous

	expected := "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"
	if *data1.Next != expected {
		t.Errorf("url don't match Actual :\n '%s' \nvs Expected: \n '%s'", *data1.Next, expected)
		return
	}

	data2, err := c.pokeapiClient.ListLocation(c.nextLocationURL)
	if err != nil {
		t.Errorf("readbody error: %s", err.Error())
		return
	}
	c.nextLocationURL = data2.Next
	c.prevLocationURL = data2.Previous

	if data2.Next == nil {
		t.Error("data is error")
		return
	}

	expected_previousValue := ""
	if data2.Previous == nil {
		t.Error("previous is nil")
		return
	}

	if *data2.Previous == expected_previousValue {
		t.Errorf("previos don't match actual : \n'%s' \nvs Expeted: \n '%s'", *data2.Previous, expected_previousValue)
	}

}
