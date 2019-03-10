package truckersmp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TruckersClient defines a new client for go-TruckersMP
type TruckersClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new TruckersClient
func NewClient(httpClient *http.Client) *TruckersClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &TruckersClient{
		baseURL:    "https://api.truckersmp.com/v2",
		httpClient: httpClient,
	}
}

func (c *TruckersClient) doRequest(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.baseURL+path, nil)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("error while performing request, did not recive 200 status code: %s", body)
	}

	return body, nil
}

// PlayerLookup looks up for a given player and returs the player
// or an error if it occurs
func (c TruckersClient) PlayerLookup(p string) (*Player, error) {
	b, err := c.doRequest("/player/" + p)
	if err != nil {
		return nil, err
	}

	var res PlayerResponse
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	if res.Error {
		return nil, fmt.Errorf("failed to lookup player %s", p)
	}

	return &res.Player, nil
}
