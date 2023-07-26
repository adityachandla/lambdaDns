package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const baseUrl = "https://j0k3usty63.execute-api.us-east-1.amazonaws.com/stage/"

type ResponseStatus uint8

const (
	REDIRECT ResponseStatus = iota
	FETCHED
	FAILED
)

func (rs *ResponseStatus) MarshalJSON() ([]byte, error) {
	switch *rs {
	case REDIRECT:
		return []byte("\"REDIRECT\""), nil
	case FETCHED:
		return []byte("\"FETCHED\""), nil
	case FAILED:
		return []byte("\"FAILED\""), nil
	}
	return nil, fmt.Errorf("Invalid value")
}

func (rs *ResponseStatus) UnmarshalJSON(inp []byte) error {
	strInp := strings.Trim(string(inp), "\"")
	switch strInp {
	case "REDIRECT":
		*rs = REDIRECT
	case "FETCHED":
		*rs = FETCHED
	case "FAILED":
		*rs = FAILED
	default:
		return fmt.Errorf("Got value %s %v", strInp, inp)
	}
	return nil
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
	log.Printf("Sending request to %s with %s", endPoint, remaining)
	req, err := json.Marshal(UrlRequest{remaining})
	Check(err)
	res, err := http.Post(baseUrl+endPoint, "application/json", bytes.NewBuffer(req))
	Check(err)
	defer res.Body.Close()
	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	Check(err)
	return &response
}
