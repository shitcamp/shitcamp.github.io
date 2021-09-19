import React from "react";
import { Container } from "react-bootstrap";

import ClipEmbed from "components/twitch/ClipEmbed";
import Clips from "components/videos/Clips";

class ClipsPage extends React.PureComponent {
  constructor(props) {
    super(props);

    this.state = {
      selectedClipName: "",
    };
  }

  handleClipClick = (e, clip) => {
    e.preventDefault();
    this.setState({
      selectedClipName: clip.id,
    });
  };

  render() {
    const { userNames } = this.props;
    const { selectedClipName } = this.state;

    return (
      <Container className="home">
        <h3>Popular Clips</h3>

        {selectedClipName !== "" && (
          <ClipEmbed clipName={selectedClipName} id={"selected-clip"} />
        )}

        <Clips userNames={userNames} onClipClick={this.handleClipClick} />
      </Container>
    );
  }
}

export default ClipsPage;
