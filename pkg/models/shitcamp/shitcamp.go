package shitcamp

import (
	"encoding/json"
	"sort"

	logger "github.com/sirupsen/logrus"
)

func GetStreamerNames() []string {
	streamers := make([]string, 0)

	for n, _ := range allStreamers {
		streamers = append(streamers, n)
	}

	sort.Strings(streamers)

	return streamers
}

func GetStreamers() []*User {
	streamers := make([]*User, 0)

	for _, s := range allStreamers {
		streamers = append(streamers, s)
	}

	sort.SliceStable(streamers, func(i, j int) bool {
		s1 := streamers[i]
		s2 := streamers[j]

		return s1.Name <= s2.Name
	})

	return streamers
}

func GetSchedule() []*DateSchedule {
	return schedule
}

func SetSchedule(s []*DateSchedule) error {
	for _, dateS := range s {
		for _, e := range dateS.Events {
			featuredStreamers := GetFeaturedStreamersForVod(e.VideoID)
			if len(featuredStreamers) == 0 {
				featuredStreamers = make([]string, 0)
			}

			e.FeaturedUsers = featuredStreamers
		}
	}

	schedule = s

	scheduleStr, err := json.Marshal(schedule)
	l := logger.WithField("schedule", scheduleStr)
	if err != nil {
		l.WithError(err).Error("SetSchedule_error")
		return err
	}

	l.Error("SetSchedule_success")

	return nil
}

func SetFeaturedUsersForVod(vodID string, userNames []string) error {
	vodIDFeaturedMap[vodID] = userNames

	mapStr, err := json.Marshal(vodIDFeaturedMap)
	l := logger.WithField("vodIDFeaturedMap", mapStr)
	if err != nil {
		l.WithError(err).Error("SetFeaturedStreamersForVod_error")
		return err
	}

	l.WithField("vodID", vodID).Error("SetFeaturedStreamersForVod_success")

	return nil
}

func GetFeaturedStreamersForVod(vodID string) []string {
	featuredStreamers, ok := vodIDFeaturedMap[vodID]
	if !ok {
		return []string{}
	}

	return featuredStreamers
}
