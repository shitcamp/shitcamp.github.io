import React from "react";
import { Container } from "react-bootstrap";

import StreamEmbed from "components/twitch/StreamEmbed";
import DateSchedule from "components/event/Events";
import AccordianWrapper from "components/accordian/AccordianWrapper";

import { getSchedule } from "apis";

import "pages/schedule/Schedule.css";

class Schedule extends React.PureComponent {
  constructor(props) {
    super(props);

    this.state = {
      scheduleDates: [],
      selectedVideoID: "",
    };
  }

  async componentDidMount() {
    let ret = await getSchedule();
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      const { dates } = ret.resp;

      this.setState({
        scheduleDates: dates,
      });
    }
  }

  handleVideoClick = (e, video_id) => {
    e.preventDefault();

    this.setState({
      selectedVideoID: video_id,
    });
  };

  render() {
    const { selectedVideoID, scheduleDates } = this.state;

    return (
      <React.Fragment>
        {selectedVideoID !== "" && (
          <StreamEmbed videoID={selectedVideoID} id={"selected-video"} />
        )}

        <Container className="schedule-page">
          <h3>Schedule</h3>

          {Array.isArray(scheduleDates) && scheduleDates.length > 0 ? (
            scheduleDates.map((d) => (
              <AccordianWrapper key={d.date} title={d.date}>
                <DateSchedule
                  events={d.events}
                  onVideoClick={this.handleVideoClick}
                />
              </AccordianWrapper>
            ))
          ) : (
            <h5>Schedule is not available</h5>
          )}
        </Container>
      </React.Fragment>
    );
  }
}

export default Schedule;
