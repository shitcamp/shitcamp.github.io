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

func GetFeaturedStreamersForStreams() map[string][]string {
	return streamIDFeaturedMap
}

func GetFeaturedStreamersForStream(streamID string) []string {
	featuredStreamers, ok := streamIDFeaturedMap[streamID]
	if !ok {
		return []string{}
	}

	return featuredStreamers
}

func SetFeaturedUsersForStream(streamID string, userNames []string) error {
	streamIDFeaturedMap[streamID] = userNames

	mapStr, err := json.Marshal(streamIDFeaturedMap)
	l := logger.WithField("streamIDFeaturedMap", string(mapStr))
	if err != nil {
		l.WithError(err).Error("SetFeaturedUsersForStream_error")
		return err
	}
	l.WithField("streamID", streamID).Info("SetFeaturedUsersForStream_success")

	return nil
}

func GetFeaturedStreamersForVods() map[string][]string {
	return vodIDFeaturedMap
}

func GetFeaturedStreamersForVod(vodID string) []string {
	featuredStreamers, ok := vodIDFeaturedMap[vodID]
	if !ok {
		return []string{}
	}

	return featuredStreamers
}

func SetFeaturedUsersForVod(vodID string, userNames []string) error {
	vodIDFeaturedMap[vodID] = userNames

	mapStr, err := json.Marshal(vodIDFeaturedMap)
	l := logger.WithField("vodIDFeaturedMap", string(mapStr))
	if err != nil {
		l.WithError(err).Error("SetFeaturedStreamersForVod_error")
		return err
	}
	l.WithField("vodID", vodID).Info("SetFeaturedStreamersForVod_success")

	return nil
}
