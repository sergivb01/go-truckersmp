package truckersmp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

// GetPlayer looks up for a given player and returs the player
// or an error if it occurs
func (c TruckersClient) GetPlayer(id int) (*Player, error) {
	b, err := c.doRequest("/player/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	var res PlayerResponse
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	if res.Error {
		return nil, fmt.Errorf("failed to lookup player %d", id)
	}

	return &res.Player, nil
}

// GetPlayerBans looks up for the latest 5 bans of a player and
// returns the list of bans or an error
func (c TruckersClient) GetPlayerBans(id int) (*[]Ban, error) {
	b, err := c.doRequest("/bans/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	var res BansResponse
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	if res.Error {
		return nil, fmt.Errorf("failed to lookup bans for player %d", id)
	}

	return &res.Bans, nil
}

// GetServers looks up for the list of servers and returns
// the servers or an error
func (c TruckersClient) GetServers() (*[]Server, error) {
	b, err := c.doRequest("/servers")
	if err != nil {
		return nil, err
	}

	var res ServersResponse
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	if res.Error != "false" {
		return nil, errors.New("failed to lookup servers")
	}

	return &res.Servers, nil
}

// GetGameTime looks up for the game time and returns the time
// or an error
func (c TruckersClient) GetGameTime() (*int64, error) {
	b, err := c.doRequest("/game_time")
	if err != nil {
		return nil, err
	}

	var res GameTimeResponse
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	if res.Error {
		return nil, errors.New("failed to lookup for game time")
	}

	return &res.GameTime, nil
}
