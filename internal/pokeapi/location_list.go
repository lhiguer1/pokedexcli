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

	return locationAreasRes, nil
}
