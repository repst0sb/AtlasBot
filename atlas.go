package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const URL = "https://atlasbus.by/api/search"

type AtlasResponse struct {
	Rides []Ride `json:"rides"`
}

type Ride struct {
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	FreeSeats int    `json:"freeSeats"`
	Price     int    `json:"price"`
}

func AtlasRequest(idFrom string, idTo string, date string) (*AtlasResponse, error) {
	params := url.Values{}
	params.Add("from_id", idFrom)
	params.Add("to_id", idTo)
	params.Add("date", date)
	params.Add("passengers", "1")
	url := URL + "?" + params.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ride AtlasResponse
	_ = json.Unmarshal(data, &ride)
	return &ride, nil

}
