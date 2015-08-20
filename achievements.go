package steamapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type GlobalAchievementPercentages struct {
	AchievementPercentages struct {
		Achievements []struct {
			Name    string  `json:"name"`
			Percent float32 `json:"percent"`
		} `json:"achievements"`
	} `json:"achievementpercentages"`
}

func (s SteamAPI) GetGlobalAchievementPercentages(appid int) (*GlobalAchievementPercentages, error) {
	// Convert argument in method from
	// an Integer to a String

	sAppId := strconv.Itoa(appid)

	// Prepare the RequestURI
	// for our request
	requestURI := "http://api.steampowered.com/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002/?gameid=" + sAppId + "&format=json"

	// Create client struct to format
	// our request safely

	client := &http.Client{Timeout: time.Duration(60) * time.Second}

	req, err := http.NewRequest("GET", requestURI, nil)
	if err != nil {
		return nil, err
	}

	// Send our request

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Umarshal our json body into our struct and return it
	var jsonResp GlobalAchievementPercentages

	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	return &jsonResp, nil

}
