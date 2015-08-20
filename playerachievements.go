package steamapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type PlayerStatus struct {
	PlayerStats struct {
		SteamId      string `json:"steamID"`
		GameName     string `json:"gameName"`
		Achievements []struct {
			APIName     string `json:"apiname"`
			Achieved    int    `json:"achieved"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"achievements"`
		Success bool `json:"success"`
	} `json:"playerstats"`
}

func (s SteamAPI) GetPlayerAchievements(steamid int, appid int, lang ...string) (*PlayerStatus, error) {
	// Comments to be added later its quite obvious what it does

	sSteamId := strconv.Itoa(steamid)
	sAppId := strconv.Itoa(appid)

	var requestURI string
	if len(lang) == 0 {
		requestURI = "http://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v0001/?appid=" + sAppId + "&key=" + s.APIKey + "&steamid=" + sSteamId

	} else {
		requestURI = "http://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v0001/?appid=" + sAppId + "&key=" + s.APIKey + "&steamid=" + sSteamId + "&l=" + lang[0]

	}

	client := http.Client{Timeout: time.Duration(60) * time.Second}

	req, err := http.NewRequest("GET", requestURI, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jsonResp PlayerStatus
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	return &jsonResp, nil

}
