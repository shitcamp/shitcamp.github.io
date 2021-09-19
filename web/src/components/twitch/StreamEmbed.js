import "components/twitch/StreamEmbed.css";

// https://dev.twitch.tv/docs/embed/everything
// https://github.com/talk2MeGooseman/react-twitch-embed-video
//https://philna.sh/blog/2020/03/23/responsive-twitch-embed/
function StreamEmbed(props) {
  const { channel } = props;

  const domainName = window.location.hostname;

  let url = new URL("https://player.twitch.tv");
  let qParams = new URLSearchParams();
  qParams.append("autoplay", true);
  qParams.append("channel", channel);
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
    <div className="twitch">
      <div className="twitch-video">
        <iframe
          src={videoSrc}
          frameBorder="0"
          scrolling="no"
          allowFullScreen={true}
        ></iframe>
      </div>

      <div className="twitch-chat">
        <iframe frameBorder="0" scrolling="no" src={chatSrc}></iframe>
      </div>
    </div>
  );
}

export default StreamEmbed;
