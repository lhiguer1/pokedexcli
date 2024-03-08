package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreasResponse, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	// cache hit
	if data, ok := c.cache.Get(fullURL); ok {
		locationAreasRes := LocationAreasResponse{}
		err := json.Unmarshal(data, &locationAreasRes)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasRes, nil
	}

	// create http request
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	// perform request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer res.Body.Close()

	// process request
	if res.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasRes := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasRes)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	// update cache
	c.cache.Add(fullURL, data)

	return locationAreasRes, nil
}
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	fullURL := baseURL + "/location-area/" + locationAreaName

	// cache hit
	if data, ok := c.cache.Get(fullURL); ok {
		locationArea := LocationArea{}

		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	// create http request
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationArea{}, err
	}

	// perform request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()

	// process request
	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	// update cache
	c.cache.Add(fullURL, data)

	return locationArea, nil
}
