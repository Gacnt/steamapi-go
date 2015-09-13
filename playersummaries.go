package steamapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type PlayerSummaries struct {
	Response struct {
		Players []struct {
			SteamId        string `json:"steamid"`
			ComVisState    int    `json:"communityvisibilitystate"`
			ProfState      int    `json:"profilestate"`
			PersonaName    string `json:"personaname"`
			LastLogOff     int    `json:"lastlogoff"`
			ProfileUrl     string `json:"profileurl"`
			Avatar         string `json:"avatar"`
			AvatarMed      string `json:"avatarmedium"`
			AvatarFull     string `json:"avatarfull"`
			PersonaState   int    `json:"personastate"`
			RealName       string `json:"realname"`
			PrimaryClId    string `json:"primaryclanid"`
			TimeCreated    int    `json:"timecreated"`
			PersonaStateFl int    `json:"personastateflags"`
			LocCountryCode string `json:"loccountrycode"`
			LocStateCode   string `json:"locstatecode"`
			LocCityId      int    `json:"loccityid"`
		} `json:"players"`
	} `json:"response"`
}

func (s SteamAPI) GetPlayerSummaries(steamIds ...uint64) (*PlayerSummaries, error) {

	// Get data from the arguments of
	// the method and convert them to
	// a string, as well concatenate
	// them into one string

	var requestURI bytes.Buffer
	requestURI.WriteString("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=" + s.APIKey + "&steamids=")

	for _, id := range steamIds {
		sSteamId := strconv.Itoa(id)
		requestURI.WriteString("," + sSteamId)

	}

	// Prepare our new data into a request

	client := http.Client{Timeout: time.Duration(60) * time.Second}

	req, err := http.NewRequest("GET", requestURI.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Send our new Json Response to our struct and return it

	var jsonResp PlayerSummaries

	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	return &jsonResp, nil
}
