import React from "react";

import { Accordion, Container } from "react-bootstrap";

import TwitchEmbed from "components/twitchEmbed/TwitchEmbed";
import Streams from "components/videos/Streams";
import Vods from "components/videos/Vods";

import { getLiveStreams, getUsers } from "apis";

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

class Home extends React.PureComponent {
  constructor(props) {
    super(props);

    this.state = {
      selectedUserStream: "",
      liveStreams: [],
      userNames: [],
    };
  }

  async componentDidMount() {
    let ret = await getLiveStreams();
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      this.setState({
        liveStreams: ret.resp,
      });
    }
    ret = await getUsers();
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      this.setState({
        userNames: ret.resp,
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
        {selectedUserStream !== "" && (
          <TwitchEmbed
            channel={selectedUserStream}
            id={"homepage-stream"}
            // chat="mobile" // TODO: detect mobile devices
          />
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
