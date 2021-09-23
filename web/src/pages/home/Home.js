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

const SHITCAMP_START_TIME = "2021-09-26T19:00:00.00-07:00";

class Home extends React.PureComponent {
  constructor(props) {
    super(props);

    this.state = {
      selectedUserStream: "",
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

      let events = [];
      for (const dateSchedule of dates) {
        for (const e of dateSchedule.events) {
          if (Date.parse(e.start_time) > Date.now()) {
            events.push(e);
          }
        }
      }

      this.setState({
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
      selectedUserStream,
      liveStreams,
      scheduleEvents,
      isLatestSchedule,
    } = this.state;

    return (
      <React.Fragment>
        <Container className="data-alert">
          <Alert variant="success">
            <Countdown
              date={SHITCAMP_START_TIME}
              renderer={({ days, hours, minutes, seconds, completed }) => {
                if (completed) {
                  return <React.Fragment />;
                }

                return (
                  <div className="countdown-container">
                    Shitcamp starts in
                    <div>
                      <span className="countdown-val">{days}</span> days{" "}
                      <span className="countdown-val">{hours}</span> hours{" "}
                      <span className="countdown-val">{minutes}</span> minutes{" "}
                      <span className="countdown-val">{seconds}</span> seconds
                    </div>
                  </div>
                );
              }}
            />
          </Alert>
        </Container>

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
            {/* <Alert variant="warning">
            This site will be updated with more accurate schedule data once
            Shitcamp starts
          </Alert> */}
            {Array.isArray(scheduleEvents) && scheduleEvents.length > 0 ? (
              <React.Fragment>
                <ScheduleAlert isLatestSchedule={isLatestSchedule} />

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
