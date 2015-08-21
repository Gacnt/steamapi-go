package steamapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type UserStatsForGame struct {
	PlayerStats struct {
		SteamId  string `json:"steamID"`
		GameName string `json:"gameName"`
		Stats    []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"stats"`
		Achievements []struct {
			Name     string `json:"name"`
			Achieved int    `json:"achieved"`
		} `json:"achievements"`
	} `json:"playerstats"`
}

func (s SteamAPI) GetUserStatsForGame(steamid int, appid int, lang ...string) (*UserStatsForGame, error) {

	sSteamId := strconv.Itoa(steamid)
	sAppId := strconv.Itoa(appid)

	var requestURI string

	if lang[0] == "" {
		requestURI = "http://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002/?appid=" + sAppId + "&key=" + s.APIKey + "&steamid=" + sSteamId
	} else {
		requestURI = "http://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002/?appid=" + sAppId + "&key=" + s.APIKey + "&steamid=" + sSteamId + "&l=" + lang[0]

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

	var jsonResp UserStatsForGame
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	return &jsonResp, nil

}
