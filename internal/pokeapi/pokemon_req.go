package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + pokemonName

	// cache hit
	if data, ok := c.cache.Get(fullURL); ok {
		pokemon := Pokemon{}

		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	// create http request
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return Pokemon{}, err
	}

	// perform request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	// process request
	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// update cache
	c.cache.Add(fullURL, data)

	return pokemon, nil
}
