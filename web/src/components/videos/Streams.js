import React from "react";

import Videos from "components/videos/Videos";

function Streams(props) {
  const { streams, handleStreamClick } = props;

  return (
    <Videos
      videos={streams}
      titleMsg="Shit's gettin' litty, come thru <3"
      emptyErrMsg="No live streams found"
      onVideoClick={handleStreamClick}
    />
  );
}

export default Streams;
