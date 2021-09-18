import { get, getElapsedDuration } from "utils";

export async function getLiveStreams() {
  let streams = [
    {
      id: "43740004317",
      user_name: "tonytigre",
      title: "drawing doodles for new followers! | !insta !doodle",
      view_count: 102,
      created_at: "2021-09-18T08:17:04Z",
      url: `https://www.twitch.tv/tonytigre`,
      thumbnail_url:
        "https://static-cdn.jtvnw.net/previews-ttv/live_user_tonytigre-{width}x{height}.jpg",
    },
  ];

  streams.forEach((stream) => {
    stream.duration = getElapsedDuration(stream.created_at);
  });

  return {
    resp: streams,
    err: null,
  };
}

export async function getVods(userNames) {
  let vods = [
    {
      id: "1151405309",
      user_name: "ludwig",
      title:
        "CAN THREE GAMERS BEAT THE BEST SMASH ULTIMATE PLAYER IN THE WORLD?* (usa*)",
      created_at: "2021-09-17T17:11:07Z",
      url: "https://www.twitch.tv/videos/1148914392",
      thumbnail_url:
        "https://static-cdn.jtvnw.net/cf_vods/d1m7jfoe9zdc1j/4ec1d0730e7592010c4b_ludwig_43704072301_1631722525//thumb/thumb0-%{width}x%{height}.jpg",
      view_count: 161803,
      duration: "37m38s",
      featured_users: ["Ludwig", "Stanz"],
    },
    {
      id: "1143190672",
      user_name: "ludwig",
      title: "BEATING JUMP KING TONIGHT LETS GOOOOO",
      created_at: "2021-09-09T01:27:05Z",
      url: "https://www.twitch.tv/videos/1143190672",
      thumbnail_url:
        "https://static-cdn.jtvnw.net/cf_vods/d1m7jfoe9zdc1j/ee21b2ae3d8b76f0ee08_ludwig_43616236173_1631150814//thumb/thumb0-%{width}x%{height}.jpg",
      view_count: 291012,
      duration: "5h6m42s",
    },
    {
      id: "1137117005",
      user_name: "ludwig",
      title: "POGCHAMPS COVERAGE, TIK TOK TIME, MARIO KART WITH FRIENDS",
      created_at: "2021-09-02T19:56:12Z",
      url: "https://www.twitch.tv/videos/1137117005",
      thumbnail_url:
        "https://static-cdn.jtvnw.net/cf_vods/d1m7jfoe9zdc1j/2d5043b59311bf1cade0_ludwig_43530010093_1630612563//thumb/thumb0-%{width}x%{height}.jpg",
      view_count: 445795,
      duration: "7h3m28s",
    },
  ];

  return {
    resp: vods,
    error: null,
  };
}
