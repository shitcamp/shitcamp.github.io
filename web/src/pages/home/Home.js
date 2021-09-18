import React from "react";

import { Accordion, Container } from "react-bootstrap";

import TwitchEmbed from "components/twitchEmbed/TwitchEmbed";
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

class Home extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      selectedUserStream: "ludwig",
      liveStreams: [],
    };
  }

  async componentDidMount() {
    let ret = await getLiveStreams();
    if (ret.error != null) {
      console.error(ret.error);
    }
    this.setState({
      liveStreams: ret.resp,
    });
  }

  render() {
    const { selectedUserStream, liveStreams } = this.state;

    return (
      <React.Fragment>
        <Container className="home">
          {selectedUserStream !== "" && (
            <TwitchEmbed
              channel={selectedUserStream}
              // id="qtcinderella-1"
              // video="1144988822" // ID of VOD to play
              // chat="mobile" // TODO: detect mobile devices
            />
          )}

          <AccordianWrapper title="Live Now">
            <Streams streams={liveStreams} />
          </AccordianWrapper>

          <AccordianWrapper title="Upcoming streams">
            {/* next 3 streams? */}
          </AccordianWrapper>

          <AccordianWrapper title="Past streams">
            <Vods
              userNames={[
                "Australia",
                "Canada",
                "USA",
                "Poland",
                "Spain",
                "France",
              ]}
            />
          </AccordianWrapper>
        </Container>
      </React.Fragment>
    );
  }
}

export default Home;
