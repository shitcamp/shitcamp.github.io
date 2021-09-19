import "components/twitch/ClipEmbed.css";

// https://dev.twitch.tv/docs/embed/video-and-clips/#non-interactive-iframes-for-clips
// https://github.com/talk2MeGooseman/react-twitch-embed-video
// https://philna.sh/blog/2020/03/23/responsive-twitch-embed/
function ClipEmbed(props) {
  const { clipName } = props;

  const domainName = window.location.hostname;

  let url = new URL("https://clips.twitch.tv/embed");
  let qParams = new URLSearchParams();
  qParams.append("allowfullscreen", true);
  qParams.append("clip", clipName);
  qParams.append("parent", domainName);
  qParams.append("preload", "auto");
  url.search = qParams.toString();
  const videoSrc = url.toString();

  return (
    <div className="twitch-clip">
      <div className="twitch-clip-embed">
        <iframe
          src={videoSrc}
          frameBorder="0"
          scrolling="no"
          allowFullScreen={true}
        ></iframe>
      </div>
    </div>
  );
}

export default ClipEmbed;
