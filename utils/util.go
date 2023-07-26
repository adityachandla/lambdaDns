package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseUrl = "https://j0k3usty63.execute-api.us-east-1.amazonaws.com/stage/"

type ResponseStatus uint8

const (
	REDIRECT ResponseStatus = iota
	FETCHED
	FAILED
)

func (rs ResponseStatus) MarshalJSON() ([]byte, error) {
	switch rs {
	case REDIRECT:
		return []byte("\"REDIRECT\""), nil
	case FETCHED:
		return []byte("\"FETCHED\""), nil
	case FAILED:
		return []byte("\"FAILED\""), nil
	}
	return nil, fmt.Errorf("Invalid value")
}

type Response struct {
	Status   ResponseStatus `json:"status"`
	Redirect *Response      `json:"redirect"`
	NodeId   int            `json:"nodeId"`
}

type UrlRequest struct {
	Url string `json:"url"`
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetRedirectResponse(endPoint, remaining string) *Response {
	req, err := json.Marshal(UrlRequest{remaining})
	Check(err)
	res, err := http.Post(baseUrl+endPoint, "application/json", bytes.NewBuffer(req))
	Check(err)
	defer res.Body.Close()
	var response Response
	json.NewDecoder(res.Body).Decode(&response)
	return &Response{Status: FETCHED, Redirect: &response}
}
