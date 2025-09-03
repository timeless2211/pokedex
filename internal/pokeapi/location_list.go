package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/timeless2211/pokedexcli/internal/pokecache"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	interval := 5 * time.Millisecond
	cache := pokecache.NewCache(interval)

	if pageURL != nil {
		if val, ok := cache.Get(*pageURL); ok {
			locationsResp := RespShallowLocations{}
			err := json.Unmarshal(val, &locationsResp)
			if err != nil {
				return RespShallowLocations{}, err
			}
		}
	}

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	if pageURL != nil {
		cache.Add(*pageURL, dat)
		return locationsResp, nil
	} else {
		cache.Add(url, dat)
		return locationsResp, nil
	}

	return locationsResp, nil
}
