package twitch

import (
	"net/http"
	"net/url"
	"time"

	"github.com/shitcamp-unofficial/shitcamp/pkg/cache"

	"github.com/shitcamp-unofficial/shitcamp/pkg/utils"
)

const (
	// never changes
	usersCacheExpiry    = 30 * time.Minute
	channelsCacheExpiry = 30 * time.Minute

	// only changes at end of stream
	videosCacheExpiry = 10 * time.Minute

	// can change relatively frequently
	clipsCacheExpiry = 5 * time.Minute
)

func GetUsers() (*Setup, error) {
	params := url.Values{}
	params.Add("login", "hasanabi")
	params.Add("login", "ludwig")

	queryURL := "/users" + params.Encode()
	if v := cache.Get(queryURL); v != nil {
		return v, nil
	}

	resp := make(map[string]interface{})
	err := utils.TwitchRequest(queryURL, http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}

	cache.Set(queryURL, resp, usersCacheExpiry)
	return resp, nil
}

func GetStreams() ([]*LiveStream, error) {
	params := url.Values{}
	params.Add("user_id", "247808909")
	params.Add("user_id", "207813352")

	params.Add("user_login", "qtcinderella")

	queryURL := "/streams" + params.Encode()
	if v := cache.Get(queryURL); v != nil {
		return v.([]*LiveStream), nil
	}

	resp := make(map[string]interface{})
	err := utils.TwitchRequest(queryURL, http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}

	cache.SetDefault(queryURL, resp)
	return resp, nil
}

func GetVods() ([]*Vod, error) {
	params := url.Values{}
	params.Add("user_id", "40934651")
	params.Add("sort", "time")
	params.Add("period", "week")

	queryURL := "/videos" + params.Encode()
	if v := cache.Get(queryURL); v != nil {
		return v.([]*Vod), nil
	}

	resp := make(map[string]interface{})
	err := utils.TwitchRequest(queryURL, http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}

	cache.Set(queryURL, resp, videosCacheExpiry)
	return resp, nil
}

func GetChannels() (*Setup, error) {
	params := url.Values{}
	params.Add("broadcaster_id", "40934651")
	params.Add("broadcaster_id", "247808909")

	queryURL := "/channels" + params.Encode()
	if v := cache.Get(queryURL); v != nil {
		return v, nil
	}

	resp := make(map[string]interface{})
	err := utils.TwitchRequest(queryURL, http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}

	cache.Set(queryURL, resp, channelsCacheExpiry)
	return resp, nil
}

func GetClips() ([]*Clip, error) {
	params := url.Values{}
	params.Add("broadcaster_id", "40934651")
	params.Add("broadcaster_id", "247808909")

	params.Add("started_at", time.Now().Format(time.RFC3339))

	queryURL := "/clips" + params.Encode()
	if v := cache.Get(queryURL); v != nil {
		return v.([]*Clip), nil
	}

	resp := make(map[string]interface{})
	err := utils.TwitchRequest(queryURL, http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}

	cache.Set(queryURL, resp, clipsCacheExpiry)
	return resp, nil
}
