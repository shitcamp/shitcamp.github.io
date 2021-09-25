package schedule

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shitcamp-unofficial/shitcamp/pkg/models/twitch"

	"github.com/shitcamp-unofficial/shitcamp/pkg/cache"
	logger "github.com/sirupsen/logrus"
)

const (
	scheduleCacheExpiry = 5 * time.Minute
)

func updateScheduleData(s *Schedule) error {
	var videoIDsToGet []string
	for _, dateS := range s.Dates {
		for _, e := range dateS.Events {
			if e.VideoID != "" {
				videoIDsToGet = append(videoIDsToGet, e.VideoID)
			}
		}
	}

	videosMap, err := twitch.GetVodsByIDs(videoIDsToGet)
	if err != nil {
		return fmt.Errorf("updateScheduleData error: %v", err)
	}

	for _, dateS := range s.Dates {
		for _, e := range dateS.Events {
			if video, ok := videosMap[e.VideoID]; ok {
				e.Vod = video
			}
		}
	}

	return nil
}

const (
	scheduleCacheKey = "shitcamp_schedule"
)

func GetSchedule() (*Schedule, error) {
	if v := cache.Get(scheduleCacheKey); v != nil {
		return v.(*Schedule), nil
	}

	err := updateScheduleData(schedule)
	if err != nil {
		return nil, err
	}

	cache.Set(scheduleCacheKey, schedule, scheduleCacheExpiry)
	return schedule, nil
}

func SetSchedule(s *Schedule) error {
	err := updateScheduleData(s)
	if err != nil {
		return err
	}

	schedule = s
	cache.Delete(scheduleCacheKey)

	scheduleStr, err := json.Marshal(schedule)
	l := logger.WithField("schedule", string(scheduleStr))
	if err != nil {
		l.WithError(err).Error("SetSchedule_error")
		return err
	}
	l.Info("SetSchedule_success")

	return nil
}
