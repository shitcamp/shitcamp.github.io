package schedule

import (
	"encoding/json"

	data2 "github.com/shitcamp-unofficial/shitcamp/pkg/data"

	logger "github.com/sirupsen/logrus"
)

var schedule = make([]*DateSchedule, 0)

//"id": "43740004317",
//"user_name": "QTCinderella",
//"title": "drawing doodles for new followers! | !insta !doodle",
//"created_at": "2021-09-18T08:17:04Z",
//"url": "https://www.twitch.tv/QTCinderella",
//"thumbnail_url":"https://static-cdn.jtvnw.net/previews-ttv/live_user_tonytigre-{width}x{height}.jpg",
//"view_count": 102

func InitScheduleData() {
	var newSchedule = make([]*DateSchedule, 0)
	err := json.Unmarshal([]byte(data2.InitialScheduleStr), &newSchedule)
	if err != nil {
		logger.WithError(err).Error("init_unmarshal_schedule_error")
		return
	}

	err = SetSchedule(newSchedule)
	if err != nil {
		logger.WithError(err).Error("init_SetSchedule_error")
	}
}
