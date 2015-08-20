package steamapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type SteamAPI struct {
	APIKey string
}

type GetAppNews struct {
	AppNews struct {
		AppId     int `json:"appid"`
		NewsItems []struct {
			Gid           string `json:"gid"`
			Title         string `json:"title"`
			Url           string `json:"url"`
			IsExternalUrl bool   `json:"is_external_url"`
			Author        string `json:"author"`
			Contents      string `json:"contents"`
			Feedlabel     string `json:"feedlabel"`
			Date          int    `json:"date"`
		} `json:"newsitems"`
	} `json:"appnews"`
}

func (s SteamAPI) GetNewsForApp(appid, count, maxlength int) (*GetAppNews, error) {

	// Convert the `int` variables from
	// the arguments supplied to the
	// method

	sAppid := strconv.Itoa(appid)
	sCount := strconv.Itoa(count)
	sMaxlength := strconv.Itoa(maxlength)

	// Set the appropriate URI for the
	// API request

	requestURI := "http://api.steampowered.com/ISteamNews/GetNewsForApp/v0002/?appid=" + sAppid + "&count=" + sCount + "&maxlength=" + sMaxlength + "&format=json"

	// Do this the proper way so that
	// there is no error for any api
	// requests that may take too long
	// to return

	client := &http.Client{Timeout: time.Duration(60) * time.Second}
	req, err := http.NewRequest("GET", requestURI, nil)
	if err != nil {
		return nil, err
	}

	// Send off our request to the
	// API service

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Get and send our data to our struct

	var jsonResp GetAppNews
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	// Everybody is happy! Send back structured data
	return &jsonResp, nil

}
