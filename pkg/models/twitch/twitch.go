package twitch

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/shitcamp-unofficial/shitcamp/pkg/cache"
	"github.com/shitcamp-unofficial/shitcamp/pkg/config"
	"github.com/shitcamp-unofficial/shitcamp/pkg/models/shitcamp"
	"github.com/shitcamp-unofficial/shitcamp/pkg/utils"
)

const (
	// never changes
	usersCacheExpiry = 2 * time.Hour

	// available videos only change when there is a new stream.
	// note that the video data (eg. view_count) can change more frequently though
	videosCacheExpiry = 10 * time.Minute

	// can change relatively frequently
	clipsCacheExpiry = 5 * time.Minute
)

type twitchResp struct {
	Data       interface{} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"` // we don't need to use this since our data requests are small enough
}

type twitchUser struct {
	ID              string    `json:"id"`
	Login           string    `json:"login"`
	DisplayName     string    `json:"display_name"`
	Type            string    `json:"type"`
	BroadcasterType string    `json:"broadcaster_type"`
	Description     string    `json:"description"`
	ProfileImageUrl string    `json:"profile_image_url"`
	OfflineImageUrl string    `json:"offline_image_url"`
	ViewCount       int       `json:"view_count"`
	CreatedAt       time.Time `json:"created_at"`
}

type twitchStream struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserLogin    string    `json:"user_login"`
	UserName     string    `json:"user_name"`
	GameId       string    `json:"game_id"`
	GameName     string    `json:"game_name"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"started_at"`
	Language     string    `json:"language"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	TagIDs       []string  `json:"tag_ids"`
	IsMature     bool      `json:"is_mature"`
}

type twitchVideo struct {
	ID            string      `json:"id"`
	StreamID      string      `json:"stream_id"`
	UserID        string      `json:"user_id"`
	UserLogin     string      `json:"user_login"`
	UserName      string      `json:"user_name"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	CreatedAt     time.Time   `json:"created_at"`
	PublishedAt   time.Time   `json:"published_at"`
	URL           string      `json:"url"`
	ThumbnailURL  string      `json:"thumbnail_url"`
	Viewable      string      `json:"viewable"`
	ViewCount     int         `json:"view_count"`
	Language      string      `json:"language"`
	Type          string      `json:"type"`
	Duration      string      `json:"duration"`
	MutedSegments interface{} `json:"muted_segments"`
}

const twitchMaxClips = 100

type twitchClip struct {
	ID              string    `json:"id"`
	URL             string    `json:"url"`
	EmbedURL        string    `json:"embed_url"`
	BroadcasterID   string    `json:"broadcaster_id"`
	BroadcasterName string    `json:"broadcaster_name"`
	CreatorID       string    `json:"creator_id"`
	CreatorName     string    `json:"creator_name"`
	VideoID         string    `json:"video_id"`
	GameID          string    `json:"game_id"`
	Language        string    `json:"language"`
	Title           string    `json:"title"`
	ViewCount       int       `json:"view_count"`
	CreatedAt       time.Time `json:"created_at"`
	ThumbnailURL    string    `json:"thumbnail_url"`
	Duration        float64   `json:"duration"`
}

func GetUserIDs(userNames []string) (userIDMap map[string]string, err error) {
	if len(userNames) == 0 {
		return nil, fmt.Errorf("no user names specified")
	}

	const usersURL = "/users"

	cacheKey := usersURL
	allUserIDsMap := make(map[string]string)
	if v := cache.Get(cacheKey); v != nil {
		allUserIDsMap = v.(map[string]string)
	}

	params := url.Values{}
	for _, name := range userNames {
		if _, ok := allUserIDsMap[name]; !ok {
			params.Add("login", name)
		}
	}
	queryString := params.Encode()
	if queryString != "" {
		queryURL := usersURL + "?" + queryString

		var users []*twitchUser
		resp := &twitchResp{Data: &users}
		err = utils.TwitchRequest(queryURL, http.MethodGet, &resp)
		if err != nil {
			return nil, err
		}

		for _, user := range users {
			allUserIDsMap[strings.ToLower(user.DisplayName)] = user.ID
		}
		cache.Set(cacheKey, allUserIDsMap, usersCacheExpiry)
	}

	userIDMap = make(map[string]string)
	for _, name := range userNames {
		userIDMap[name] = allUserIDsMap[strings.ToLower(name)]
	}

	return userIDMap, nil
}

func GetStreams(userNames []string) ([]*LiveStream, error) {
	liveStreams := make([]*LiveStream, 0)

	if len(userNames) == 0 {
		return liveStreams, nil
	}

	params := url.Values{}
	params.Add("sort", "time")
	params.Add("period", "week")
	for _, user := range userNames {
		params.Add("user_login", user)
	}

	queryURL := "/streams" + "?" + params.Encode()

	cacheKey := queryURL
	if v := cache.Get(cacheKey); v != nil {
		return v.([]*LiveStream), nil
	}

	var streams []*twitchStream
	resp := &twitchResp{Data: &streams}
	err := utils.TwitchRequest(queryURL, http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}

	for _, s := range streams {
		liveStreams = append(liveStreams, &LiveStream{Video{
			ID:                s.ID,
			UserName:          s.UserName,
			Title:             s.Title,
			CreatedAt:         s.StartedAt,
			URL:               fmt.Sprintf(`https://www.twitch.tv/%s`, s.UserName),
			ThumbnailURL:      s.ThumbnailUrl,
			ViewCount:         s.ViewerCount,
			FeaturedStreamers: shitcamp.GetFeaturedStreamersForVod(s.ID),
		}})
	}

	cache.SetDefault(cacheKey, liveStreams)
	return liveStreams, nil
}

func GetVodsByIDs(videoIDs []string) (map[string]*Vod, error) {
	params := url.Values{}
	params.Add("sort", "time")

	for _, id := range videoIDs {
		params.Add("id", id)
	}

	queryURL := "/videos" + "?" + params.Encode()

	var videos []*twitchVideo
	resp := &twitchResp{Data: &videos}
	err := utils.TwitchRequest(queryURL, http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}

	vodMap := make(map[string]*Vod)
	for _, v := range videos {
		vodMap[v.ID] = &Vod{
			Video: Video{
				ID:                v.ID,
				UserName:          v.UserName,
				Title:             v.Title,
				CreatedAt:         v.CreatedAt,
				URL:               v.URL,
				ThumbnailURL:      v.ThumbnailURL,
				ViewCount:         v.ViewCount,
				FeaturedStreamers: shitcamp.GetFeaturedStreamersForVod(v.ID),
			},
			Duration: v.Duration,
		}
	}

	return vodMap, nil
}

func getVodsForUserID(userID string) ([]*Vod, error) {
	params := url.Values{}
	params.Add("user_id", userID)
	params.Add("sort", "time")
	params.Add("period", "week")

	queryURL := "/videos" + "?" + params.Encode()

	cacheKey := queryURL
	if v := cache.Get(cacheKey); v != nil {
		return v.([]*Vod), nil
	}

	var videos []*twitchVideo
	resp := &twitchResp{Data: &videos}
	err := utils.TwitchRequest(queryURL, http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}

	var vods []*Vod
	for _, v := range videos {
		if v.CreatedAt.Before(config.GetConfig().Shitcamp.GetOldestUploadTime()) {
			break
		}

		vods = append(
			vods, &Vod{
				Video: Video{
					ID:                v.ID,
					UserName:          v.UserName,
					Title:             v.Title,
					CreatedAt:         v.CreatedAt,
					URL:               v.URL,
					ThumbnailURL:      v.ThumbnailURL,
					ViewCount:         v.ViewCount,
					FeaturedStreamers: shitcamp.GetFeaturedStreamersForVod(v.ID),
				},
				Duration: v.Duration,
			},
		)
	}

	cache.Set(cacheKey, vods, videosCacheExpiry)
	return vods, nil
}

func GetVods(userNames []string) ([]*Vod, error) {
	vods := make([]*Vod, 0)

	if len(userNames) == 0 {
		return vods, nil
	}

	cacheKey := "vods/" + strings.Join(userNames, "&")
	if v := cache.Get(cacheKey); v != nil {
		return v.([]*Vod), nil
	}

	userIDMap, err := GetUserIDs(userNames)
	if err != nil {
		return nil, err
	}

	for _, userID := range userIDMap {
		userVods, err := getVodsForUserID(userID)
		if err != nil {
			return nil, err
		}

		vods = append(vods, userVods...)
	}

	sort.SliceStable(vods, func(i, j int) bool {
		v1 := vods[i]
		v2 := vods[j]

		return v1.CreatedAt.After(v2.CreatedAt)
	})

	cache.Set(cacheKey, vods, videosCacheExpiry)
	return vods, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getClipsForBroadcaster(broadcasterID string, maxNumClips int) ([]*Clip, error) {
	params := url.Values{}
	params.Add("broadcaster_id", broadcasterID)
	params.Add("started_at", config.GetConfig().Shitcamp.GetOldestUploadTime().Format(time.RFC3339))
	params.Add("first", fmt.Sprintf("%d", min(maxNumClips, twitchMaxClips)))

	queryURL := "/clips" + "?" + params.Encode()

	cacheKey := queryURL
	if v := cache.Get(cacheKey); v != nil {
		return v.([]*Clip), nil
	}

	var tClips []*twitchClip
	resp := &twitchResp{Data: &tClips}
	err := utils.TwitchRequest(queryURL, http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}

	var clips []*Clip
	for _, c := range tClips {
		clips = append(clips, &Clip{
			ID:              c.ID,
			BroadcasterName: c.BroadcasterName,
			Title:           c.Title,
			CreatedAt:       c.CreatedAt,
			URL:             c.URL,
			ThumbnailURL:    c.ThumbnailURL,
			ViewCount:       c.ViewCount,
			Duration:        fmt.Sprintf("%.0fs", c.Duration),
		})
	}

	cache.Set(cacheKey, clips, clipsCacheExpiry)
	return clips, nil
}

func GetClips(broadcasterNames []string) ([]*Clip, error) {
	clips := make([]*Clip, 0)

	if len(broadcasterNames) == 0 {
		return clips, nil
	}

	cacheKey := "clips/" + strings.Join(broadcasterNames, "&")
	if v := cache.Get(cacheKey); v != nil {
		return v.([]*Clip), nil
	}

	userIDMap, err := GetUserIDs(broadcasterNames)
	if err != nil {
		return nil, err
	}

	var maxNumClips = (twitchMaxClips / len(broadcasterNames)) + 10

	for _, userID := range userIDMap {
		userClips, err := getClipsForBroadcaster(userID, maxNumClips)
		if err != nil {
			return nil, err
		}

		clips = append(clips, userClips...)
	}

	sort.SliceStable(clips, func(i, j int) bool {
		c1 := clips[i]
		c2 := clips[j]

		return c1.ViewCount >= c2.ViewCount
	})

	if len(clips) > twitchMaxClips {
		clips = clips[:twitchMaxClips]
	}

	cache.Set(cacheKey, clips, videosCacheExpiry)
	return clips, nil
}

//func GetChannels() (*Setup, error) {
//	params := url.Values{}
//	params.Add("broadcaster_id", "40934651")
//	params.Add("broadcaster_id", "247808909")
//
//	queryURL := "/channels" +"?" + params.Encode()
//	if v := cache.Get(queryURL); v != nil {
//		return v, nil
//	}
//
//	resp := make(map[string]interface{})
//	err := utils.TwitchRequest(queryURL, http.MethodGet, &resp)
//	if err != nil {
//		return nil, err
//	}
//
//	cache.Set(queryURL, resp, channelsCacheExpiry)
//	return resp, nil
//}
