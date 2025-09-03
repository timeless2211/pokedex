package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (RespLocation, error) {
	url := baseURL + "/location-area/" + locationName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}
	locationResp := RespLocation{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespLocation{}, err
	}

	return locationResp, nil
}
