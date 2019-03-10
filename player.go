package truckersmp

// PlayerResponse is the response of a PlayerLookup
type PlayerResponse struct {
	Error  bool   `json:"error"`
	Player Player `json:"response"`
}

// Player contains the basic information of a player
type Player struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Avatar      string      `json:"avatar"`
	SmallAvatar string      `json:"smallAvatar"`
	JoinDate    string      `json:"joinDate"`
	SteamID64   int64       `json:"steamID64"`
	GroupName   string      `json:"groupName"`
	GroupID     int64       `json:"groupID"`
	Banned      bool        `json:"banned"`
	BannedUntil interface{} `json:"bannedUntil"`
	Permissions Permissions `json:"permissions"`
}

// Permissions defines player permissions
type Permissions struct {
	IsGameAdmin           bool `json:"isGameAdmin"`
	ShowDetailedOnWebMaps bool `json:"showDetailedOnWebMaps"`
}

// BansResponse is the response of a player ban lookup
type BansResponse struct {
	Error bool  `json:"error"`
	Bans  []Ban `json:"response"`
}

// Ban contains a ban information
type Ban struct {
	Expiration *string `json:"expiration"`
	TimeAdded  string  `json:"timeAdded"`
	Active     bool    `json:"active"`
	Reason     string  `json:"reason"`
	AdminName  string  `json:"adminName"`
	AdminID    int64   `json:"adminID"`
}
