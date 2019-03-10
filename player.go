package truckersmp

type PlayerResponse struct {
	Error  bool   `json:"error"`
	Player Player `json:"response"`
}

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

type Permissions struct {
	IsGameAdmin           bool `json:"isGameAdmin"`
	ShowDetailedOnWebMaps bool `json:"showDetailedOnWebMaps"`
}
