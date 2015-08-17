package steamapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SteamAPI struct {
	APIKey string
}

type GetAppNews struct {
	AppNews struct {
		AppId     int `json:"appid"`
		NewsItems []struct {
			Gid           int    `json:"gid"`
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

type JsonResponse map[string]GetAppNews

func (s SteamAPI) GetNewsForApp(appid, count, maxlength int) error {
	sAppid := strconv.Itoa(appid)
	sCount := strconv.Itoa(count)
	sMaxlength := strconv.Itoa(maxlength)

	resp, err := http.Get("http://api.steampowered.com/ISteamNews/GetNewsForApp/v0002/?appid=" + sAppid + "&count=" + sCount + "&maxlength=" + sMaxlength + "&format=json")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var jsonReturn JsonResponse

	json.Unmarshal(body, &jsonReturn)

	fmt.Println(jsonReturn)

	return nil

}
