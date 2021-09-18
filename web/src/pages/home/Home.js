import React from "react";

import TwitchEmbed from "components/twitchEmbed/TwitchEmbed";
import Vods from "components/vods/Vods";
import { getVods } from "apis/api";

const Home = () => {
  return (
    <React.Fragment>
      {"Shit's gettin' litty, come thru <3"}
      <span>Home</span>

      <Vods vods={getVods()} />

      <TwitchEmbed
        id="qtcinderella-1"
        channel="ludwig" // name of the chat room and channel to stream
        // video="1144988822" // ID of VOD to play
        // chat="mobile" // TODO: detect mobile devices
      />
    </React.Fragment>
  );
};

export default Home;
