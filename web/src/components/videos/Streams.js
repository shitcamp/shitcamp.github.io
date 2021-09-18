import React from "react";

import Videos from "components/videos/Videos";

class Streams extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {}

  handleStreamClick = (stream) => {
    console.log("stream is clicked", stream);
  };

  render() {
    const { streams } = this.props;

    return (
      <Videos
        videos={streams}
        titleMsg="Shit's gettin' litty, come thru <3"
        emptyErrMsg="No live streams found"
        onVideoClick={this.handleStreamClick}
      />
    );
  }
}

export default Streams;
