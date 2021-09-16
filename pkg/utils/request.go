package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/shitcamp-unofficial/shitcamp/pkg/config"
)

type requestOpt func(r *http.Request)

const (
	AuthorizationHeader  = "Authorization"
	TwitchClientIDHeader = "Client-Id"
)

func WithTwitchAccessToken(accessToken string) requestOpt {
	return func(r *http.Request) {
		r.Header.Set(AuthorizationHeader, "Bearer "+accessToken)
	}
}

func WithTwitchClientID(clientID string) requestOpt {
	return func(r *http.Request) {
		r.Header.Set(TwitchClientIDHeader, clientID)
	}
}

func printJSON(data interface{}) {
	switch v := data.(type) {
	case string:
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(v), &m)
		if err == nil {
			data = m
		}

	default:
		break
	}

	formattedBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return
	}

	fmt.Println(string(formattedBytes))
}

func Request(reqURL, method string, queryParams url.Values, body interface{}, response interface{}, opts ...requestOpt) (err error) {
	// create request
	var tempBody []byte
	tempBody, err = json.Marshal(body)
	if err != nil {
		return fmt.Errorf("request_body_error: %v", err)
	}

	req, err := http.NewRequest(method, reqURL, bytes.NewReader(tempBody))
	if err != nil {
		return fmt.Errorf("request_creation_error: %v", err)
	}

	//q := req.URL.Query()
	//for k, values := range queryParams {
	//	for _, v := range values {
	//		q.Add(k, v)
	//	}
	//}
	req.URL.RawQuery = queryParams.Encode()

	req.Header.Set("Content-Type", "application/json")

	for _, opt := range opts {
		opt(req)
	}

	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request_send_error: %v", err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("response_read_error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP code %d: %v", resp.StatusCode, string(respBody))
	} else {
		_ = json.Unmarshal(respBody, response)
	}

	// print debug
	if config.GetConfig().Debug {
		fmt.Println(method, req.URL.String(), resp.StatusCode)

		fmt.Println("Request:")
		printJSON(body)

		fmt.Println("Response:")
		printJSON(string(respBody))

		fmt.Println()
	}

	return err
}

const twitchAPIDomain = "https://api.twitch.tv/helix"

func TwitchRequest(path, method string, response interface{}) (err error) {
	return Request(twitchAPIDomain+path, method, nil, nil, response,
		WithTwitchClientID(config.GetConfig().Twitch.ClientID),
		WithTwitchAccessToken(config.GetConfig().Twitch.AccessToken),
	)
}
