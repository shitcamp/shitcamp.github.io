import ReactTwitchEmbedVideo from "react-twitch-embed-video";

function getRandomString() {
  return (Math.random() + 1).toString(36).substring(7);
}

// TODO:
const DEFAULT_HEIGHT = "480"; // 600
const DEFAULT_WIDTH = "940"; // '100%'

// https://dev.twitch.tv/docs/embed/everything
// https://github.com/talk2MeGooseman/react-twitch-embed-video
function TwitchEmbed(props) {
  var params = {
    id: props.id, // || getRandomString(),
    autoplay: !!props.channel,
    height: props.height || DEFAULT_HEIGHT,
    width: props.width || DEFAULT_WIDTH,
    chat: props.chat,
  };

  return (
    <div className="twitch-embed">
      <ReactTwitchEmbedVideo
        targetId={params.id} // custom id to target, if we have multiple embedded players on the page
        autoplay={params.autoplay}
        channel={props.channel} // name of the chat room and channel to stream
        video={props.video} // ID of VOD to play
        // collection={{ video: "124085610", collection: "GMEgKwTQpRQwyA" }} // VOD collection to play

        height={params.height}
        width={params.width}
        chat={params.chat}
      />
    </div>
  );
}

export default TwitchEmbed;
