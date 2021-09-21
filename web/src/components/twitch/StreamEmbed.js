import "components/twitch/StreamEmbed.css";

// https://dev.twitch.tv/docs/embed/everything
// https://github.com/talk2MeGooseman/react-twitch-embed-video
//https://philna.sh/blog/2020/03/23/responsive-twitch-embed/
function StreamEmbed(props) {
  const { channel, videoID } = props;

  const domainName = window.location.hostname;

  let url = new URL("https://player.twitch.tv");
  let qParams = new URLSearchParams();
  if (channel != null) {
    qParams.append("channel", channel);
  } else if (videoID != null) {
    qParams.append("video", `v${videoID}`);
  }
  qParams.append("autoplay", true);
  qParams.append("parent", domainName);
  qParams.append("muted", false);
  url.search = qParams.toString();
  const videoSrc = url.toString();

  url = new URL(`https://twitch.tv/embed/${channel}/chat`);
  qParams = new URLSearchParams();
  qParams.append("parent", domainName);
  qParams.append("darkpopout", true);
  url.search = qParams.toString();
  const chatSrc = url.toString();

  return (
    <div className="twitch-video">
      <div className="twitch-video-embed">
        <iframe
          src={videoSrc}
          title="Twitch stream"
          frameBorder="0"
          scrolling="no"
          allowFullScreen={true}
        ></iframe>
      </div>

      <div className="twitch-video-chat">
        <iframe
          frameBorder="0"
          scrolling="no"
          title="Twitch chat"
          src={chatSrc}
        ></iframe>
      </div>
    </div>
  );
}

export default StreamEmbed;
