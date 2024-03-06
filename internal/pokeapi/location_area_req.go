package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(pageURL *string) (LocationAreaResponse, error) {
	var endpoint = "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close() // close the response object when done

	if resp.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body) // read data from response body
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// Unmarshal
	locationAreaResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResp, nil
}
