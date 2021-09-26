import React from "react";

import { Alert, Container } from "react-bootstrap";
import Countdown from "react-countdown";

import AccordianWrapper from "components/accordian/AccordianWrapper";
import StreamEmbed from "components/twitch/StreamEmbed";
import Streams from "components/videos/Streams";
import Vods from "components/videos/Vods";
import Events from "components/event/Events";
import ScheduleAlert from "components/event/ScheduleAlert";

import { getLiveStreams, getSchedule } from "apis";

import "pages/home/Home.css";

const SHITCAMP_START_TIME = "2021-09-26T18:00:00.00-07:00";

class Home extends React.PureComponent {
  constructor(props) {
    super(props);

    this.state = {
      selectedUserStream: "",
      shitcampStartTime: SHITCAMP_START_TIME,
      liveStreams: [],
      scheduleEvents: [],
      isLatestSchedule: true,
    };
  }

  async componentDidMount() {
    const { userNames } = this.props;

    let ret = await getLiveStreams(userNames);
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      const { streams } = ret.resp;

      this.setState({
        liveStreams: streams,
        selectedUserStream: streams.length === 0 ? "" : streams[0].user_name,
      });
    }

    ret = await getSchedule();
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      const { dates, is_latest_schedule } = ret.resp;

      let startTime = "";
      let events = [];
      for (const dateSchedule of dates) {
        for (const e of dateSchedule.events) {
          if (startTime === "") {
            startTime = e.start_time;
          }

          if (Date.parse(e.start_time) > Date.now()) {
            events.push(e);
          }
        }
      }

      this.setState({
        shitcampStartTime: startTime,
        scheduleEvents: events,
        isLatestSchedule: Boolean(is_latest_schedule),
      });
    }
  }

  handleStreamClick = (e, stream) => {
    e.preventDefault();
    this.setState({
      selectedUserStream: stream.user_name,
    });
  };

  handleStreamClick = (e, stream) => {
    e.preventDefault();
    this.setState({
      selectedUserStream: stream.user_name,
    });
  };

  render() {
    const { userNames } = this.props;
    const {
      shitcampStartTime,
      selectedUserStream,
      liveStreams,
      scheduleEvents,
      isLatestSchedule,
    } = this.state;

    const timeZone = new Date().toTimeString().split(/ (.+)/)[1];

    return (
      <React.Fragment>
        <Countdown
          date={shitcampStartTime}
          renderer={({ days, hours, minutes, seconds, completed }) => {
            if (completed) {
              return <React.Fragment />;
            }

            return (
              <Container className="data-alert">
                <Alert variant="success">
                  <div className="countdown-container">
                    Shitcamp starts in
                    <div>
                      <span className="countdown-val">{days}</span> days{" "}
                      <span className="countdown-val">{hours}</span> hours{" "}
                      <span className="countdown-val">{minutes}</span> minutes{" "}
                      <span className="countdown-val">{seconds}</span> seconds
                    </div>
                  </div>
                </Alert>
              </Container>
            );
          }}
        />

        {selectedUserStream !== "" && (
          <StreamEmbed channel={selectedUserStream} id={"homepage-stream"} />
        )}
        <Container className="home">
          <AccordianWrapper title="Live Now">
            <Streams
              streams={liveStreams}
              handleStreamClick={this.handleStreamClick}
            />
          </AccordianWrapper>

          <AccordianWrapper title="Upcoming events">
            {Array.isArray(scheduleEvents) && scheduleEvents.length > 0 ? (
              <React.Fragment>
                <ScheduleAlert isLatestSchedule={isLatestSchedule} />
                <p>
                  The times shown are for your timezone: <b>{timeZone}</b>
                </p>

                <Events
                  events={scheduleEvents}
                  onVideoClick={() => {}}
                  displayStartDate={true}
                />
              </React.Fragment>
            ) : (
              <h5>No upcoming events</h5>
            )}
          </AccordianWrapper>

          <AccordianWrapper title="Latest streams">
            <Vods userNames={userNames} />
          </AccordianWrapper>
        </Container>
      </React.Fragment>
    );
  }
}

export default Home;
