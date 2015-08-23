package steamapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TradeOffers struct {
	Response struct {
		TradeOffersReceived []struct {
			TradeOfferId    string `json:"tradeofferid"`
			AccountIdOther  int    `json:"accountid_other"`
			Message         string `json:"message"`
			ExpirationTime  int    `json:"expiration_time"`
			TradeOfferState int    `json:"trade_offer_state"`
			ItemsToReceive  []struct {
				AppId      string `json:"appid"`
				ContextId  string `json:"contextid"`
				AssetId    string `json:"assetid"`
				ClassId    string `json:"classid"`
				InstanceId string `json:"instanceid"`
				Amount     string `json:"amount"`
				Missing    bool   `json:"missing"`
			} `json:"items_to_receive"`
			IsOurOffer        bool `json:"is_our_offer"`
			TimeCreated       int  `json:"time_created"`
			TimeUpdated       int  `json:"time_updated"`
			FromRealTimeTrade bool `json:"from_real_time_trade"`
		} `json:"trade_offers_received"`
	} `json:"response"`
}

type TradeOfferOptions struct {
	GetSentOffers        string
	GetReceivedOffers    string
	GetDescriptions      string
	Language             string
	ActiveOnly           string
	HistoricalOnly       string
	TimeHistoricalCutoff string
}

func (s SteamAPI) GetTradeOffers(options *TradeOfferOptions) (*TradeOffers, error) {
	requestURI := bytes.NewBufferString("https://api.steampowered.com/IEconService/GetTradeOffers/v1/")

	requestURI.WriteString("?key=" + s.APIKey)

	if options.GetSentOffers != "" {
		requestURI.WriteString("&get_sent_offers=" + options.GetSentOffers)
	}

	if options.GetReceivedOffers != "" {
		requestURI.WriteString("&get_received_offers=" + options.GetReceivedOffers)
	}

	if options.GetDescriptions != "" {
		requestURI.WriteString("&get_descriptions=" + options.GetDescriptions)
	}

	if options.Language != "" {
		requestURI.WriteString("&language=" + options.Language)
	}

	if options.ActiveOnly != "" {
		requestURI.WriteString("&active_only=" + options.ActiveOnly)
	}

	if options.HistoricalOnly != "" {
		requestURI.WriteString("&historical_only=" + options.HistoricalOnly)
	}

	if options.TimeHistoricalCutoff != "" {
		requestURI.WriteString("&time_historical_cutoff=" + options.TimeHistoricalCutoff)
	}

	fmt.Println(requestURI.String())

	req, err := http.NewRequest("GET", requestURI.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	var jsonResp TradeOffers
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return nil, err
	}

	return &jsonResp, nil

}
