package steamapi

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

type SteamAPI struct {
	APIKey string
	Client *http.Client
}

func New(apikey string, timeout int) (*SteamAPI, error) {
	cookiesjar, err := cookiejar.New(&cookiejar.Options{nil})
	if err != nil {
		return nil, err
	}

	newSteamAPI := SteamAPI{
		APIKey: apikey,
		Client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
			Jar:     cookiesjar,
		},
	}

	return &newSteamAPI, nil
}
