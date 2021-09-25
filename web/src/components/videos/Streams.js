import React from "react";

import Videos from "components/videos/Videos";

function Streams(props) {
  const { streams, handleStreamClick } = props;

  return (
    <Videos
      videos={streams}
      title={
        <React.Fragment>
          <h5>{"Shit's gettin' litty, come thru <3"}</h5>
          <p>{"Select a stream to watch"}</p>
        </React.Fragment>
      }
      emptyErrMsg="No live streams found"
      onVideoClick={handleStreamClick}
    />
  );
}

export default Streams;
