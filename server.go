package truckersmp

// ServersResponse is the response of servers lookup
type ServersResponse struct {
	Error   string   `json:"error"` // for some reason the API returns the error as an string in servers
	Servers []Server `json:"response"`
}

// Server contains a server information
type Server struct {
	ID                   int64  `json:"id"`
	Game                 string `json:"game"`
	IP                   string `json:"ip"`
	Port                 int64  `json:"port"`
	Name                 string `json:"name"`
	Shortname            string `json:"shortname"`
	Online               bool   `json:"online"`
	Players              int64  `json:"players"`
	Queue                int64  `json:"queue"`
	Maxplayers           int64  `json:"maxplayers"`
	Speedlimiter         int64  `json:"speedlimiter"`
	Collisions           bool   `json:"collisions"`
	Carsforplayers       bool   `json:"carsforplayers"`
	Policecarsforplayers bool   `json:"policecarsforplayers"`
	Afkenabled           bool   `json:"afkenabled"`
	Syncdelay            int64  `json:"syncdelay"`
}

// GameTimeResponse is the response of the game servers' time
type GameTimeResponse struct {
	Error    bool  `json:"error"`
	GameTime int64 `json:"game_time"`
}
