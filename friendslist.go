package steamapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type FriendsList struct {
	Friends struct {
		Friend []struct {
			SteamId      string `json:"steamid"`
			Relationship string `json:"relationship"`
			FriendSince  int    `json:"friend_since"`
		} `json:"friends"`
	} `json:"friendslist"`
}

func (s SteamAPI) GetFriendsList(steamid int, relationship string) (*FriendsList, error) {
	// Fetches friends list form
	// steams API for Steam Friends

	// Prepare our RequestURI

	sSteamId := strconv.Itoa(steamid)

	requestURI := "http://api.steampowered.com/ISteamUser/GetFriendList/v0001/?key=" + s.APIKey + "&steamid=" + sSteamId + "&relationship=" + relationship

	client := http.Client{Timeout: time.Duration(60) * time.Second}

	req, err := http.NewRequest("GET", requestURI, nil)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jsonResp FriendsList
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	return &jsonResp, nil

}
