import React from "react";

import { Accordion, Alert, Container } from "react-bootstrap";
import Countdown from "react-countdown";

import StreamEmbed from "components/twitch/StreamEmbed";
import Streams from "components/videos/Streams";
import Vods from "components/videos/Vods";

import { getLiveStreams } from "apis";

import "pages/home/Home.css";

function AccordianWrapper(props) {
  const { title, children } = props;

  return (
    <Accordion defaultActiveKey="0">
      <Accordion.Item eventKey="0">
        <Accordion.Header>
          <h5>{title}</h5>
        </Accordion.Header>
        <Accordion.Body>{children}</Accordion.Body>
      </Accordion.Item>
    </Accordion>
  );
}

const SHITCAMP_START_TIME = "2021-09-26T19:00:00.00-07:00";

class Home extends React.PureComponent {
  constructor(props) {
    super(props);

    const { userNames } = this.props;

    this.state = {
      userNames: userNames,
      selectedUserStream: "",
      liveStreams: [],
    };
  }

  async componentDidMount() {
    let ret = await getLiveStreams();
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      let streams = ret.resp;

      this.setState({
        liveStreams: streams,
        selectedUserStream: streams.length === 0 ? "" : streams[0].user_name,
      });
    }
  }

  handleStreamClick = (e, stream) => {
    e.preventDefault();
    this.setState({
      selectedUserStream: stream.user_name,
    });
  };

  render() {
    const { selectedUserStream, liveStreams, userNames } = this.state;

    return (
      <React.Fragment>
        <Container className="data-alert">
          {/* <Alert variant="warning">
            This site will be updated with more accurate stream data once
            Shitcamp starts
          </Alert> */}
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

          <AccordianWrapper title="Upcoming streams">
            {/* next 3 streams? */}
          </AccordianWrapper>

          <AccordianWrapper title="Past streams">
            <Vods userNames={userNames} />
          </AccordianWrapper>
        </Container>
      </React.Fragment>
    );
  }
}

export default Home;
