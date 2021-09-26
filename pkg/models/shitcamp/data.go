package shitcamp

import (
	"encoding/json"

	"github.com/shitcamp-unofficial/shitcamp/pkg/data"
	logger "github.com/sirupsen/logrus"
)

var vodIDFeaturedMap = make(map[string][]string)
var streamIDFeaturedMap = make(map[string][]string)
var teamsInfo = &TeamsInfo{}

func init() {
	err := json.Unmarshal([]byte(data.InitialVodIDFeaturedMapStr), &vodIDFeaturedMap)
	if err != nil {
		logger.WithError(err).Error("init_vod_featured_map_error")
	}

	err = json.Unmarshal([]byte(data.InitialStreamIDFeaturedMapStr), &streamIDFeaturedMap)
	if err != nil {
		logger.WithError(err).Error("init_stream_featured_map_error")
	}

	err = json.Unmarshal([]byte(data.InitialTeamsInfo), &teamsInfo)
	if err != nil {
		logger.WithError(err).Error("init_unmarshal_teamsInfo_error")
	}
}
