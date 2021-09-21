package schedule

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shitcamp-unofficial/shitcamp/pkg/models/shitcamp"

	"github.com/shitcamp-unofficial/shitcamp/pkg/models/twitch"

	"github.com/shitcamp-unofficial/shitcamp/pkg/cache"
	logger "github.com/sirupsen/logrus"
)

const (
	scheduleCacheExpiry = 5 * time.Minute
)

func updateScheduleData(s []*DateSchedule) error {
	var videoIDsToGet []string
	for _, dateS := range s {
		for _, e := range dateS.Events {
			featuredUsers := shitcamp.GetFeaturedStreamersForVod(e.VideoID)
			if len(featuredUsers) > 0 {
				e.FeaturedUsers = featuredUsers
			}

			if e.VideoID != "" {
				videoIDsToGet = append(videoIDsToGet, e.VideoID)
			}
		}
	}

	videosMap, err := twitch.GetVodsByIDs(videoIDsToGet)
	if err != nil {
		return fmt.Errorf("updateScheduleData error: %v", err)
	}

	for _, dateS := range s {
		for _, e := range dateS.Events {
			if video, ok := videosMap[e.VideoID]; ok {
				e.Vod = video
			}
		}
	}

	return nil
}

func GetSchedule() ([]*DateSchedule, error) {
	cacheKey := "shitcamp_schedule"
	if v := cache.Get(cacheKey); v != nil {
		return v.([]*DateSchedule), nil
	}

	err := updateScheduleData(schedule)
	if err != nil {
		return nil, err
	}

	cache.Set(cacheKey, schedule, scheduleCacheExpiry)
	return schedule, nil
}

func SetSchedule(s []*DateSchedule) error {
	err := updateScheduleData(s)
	if err != nil {
		return err
	}

	schedule = s

	scheduleStr, err := json.Marshal(schedule)
	l := logger.WithField("schedule", string(scheduleStr))
	if err != nil {
		l.WithError(err).Error("SetSchedule_error")
		return err
	}
	l.Info("SetSchedule_success")

	return nil
}
