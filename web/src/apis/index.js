import { get, getElapsedDuration } from "utils";

export async function getUsers() {
  // let users = [
  //   "ludwig",
  //   "qtcinderella",
  //   "nmplol",
  //   "mizkif",
  //   "hasanabi",
  //   "tonytigre",
  //   "AnnaCramling",
  // ];
  // return {
  //   resp: {
  //     users: users,
  //   },
  //   err: null,
  // };

  const ret = await get("/api/shitcamp/get_streamer_names");
  return ret;
}

export async function getLiveStreams(userNames) {
  // let streams = [
  //   {
  //     id: "43740004317",
  //     user_name: "tonytigre",
  //     title: "drawing doodles for new followers! | !insta !doodle",
  //     view_count: 102,
  //     created_at: "2021-09-18T08:17:04Z",
  //     url: `https://www.twitch.tv/tonytigre`,
  //     thumbnail_url:
  //       "https://static-cdn.jtvnw.net/previews-ttv/live_user_tonytigre-{width}x{height}.jpg",
  //   },
  //   {
  //     id: "43742009789",
  //     user_name: "AnnaCramling",
  //     title:
  //       "24 HOUR STREAM !info ~ Playing Against Subs & Later Guess The ELO And EVAL With Dad <3 ",
  //     view_count: 2756,
  //     created_at: "2021-09-18T13:05:29Z",
  //     thumbnail_url:
  //       "https://static-cdn.jtvnw.net/previews-ttv/live_user_annacramling-{width}x{height}.jpg",
  //   },
  // ];
  // streams.forEach((stream) => {
  //   stream.duration = getElapsedDuration(stream.created_at);
  // });
  // return {
  //   resp: {
  //     streams: streams,
  //   },
  //   err: null,
  // };

  let params = userNames.map((u) => {
    return ["user", u];
  });

  const ret = await get("/api/twitch/get_live_streams", params);
  if (ret.error != null) {
    return ret;
  }

  ret.resp.streams.forEach((stream) => {
    stream.duration = getElapsedDuration(stream.created_at);
  });

  return ret;
}

export async function getVods(userNames) {
  // let vods = [
  //   {
  //     id: "1151405309",
  //     user_name: "ludwig",
  //     title:
  //       "CAN THREE GAMERS BEAT THE BEST SMASH ULTIMATE PLAYER IN THE WORLD?* (usa*)",
  //     created_at: "2021-09-17T17:11:07Z",
  //     url: "https://www.twitch.tv/videos/1148914392",
  //     thumbnail_url:
  //       "https://static-cdn.jtvnw.net/cf_vods/d1m7jfoe9zdc1j/4ec1d0730e7592010c4b_ludwig_43704072301_1631722525//thumb/thumb0-%{width}x%{height}.jpg",
  //     view_count: 161803,
  //     duration: "37m38s",
  //     featured_users: ["Ludwig", "Stanz"],
  //   },
  //   {
  //     id: "1143190672",
  //     user_name: "ludwig",
  //     title: "BEATING JUMP KING TONIGHT LETS GOOOOO",
  //     created_at: "2021-09-09T01:27:05Z",
  //     url: "https://www.twitch.tv/videos/1143190672",
  //     thumbnail_url:
  //       "https://static-cdn.jtvnw.net/cf_vods/d1m7jfoe9zdc1j/ee21b2ae3d8b76f0ee08_ludwig_43616236173_1631150814//thumb/thumb0-%{width}x%{height}.jpg",
  //     view_count: 291012,
  //     duration: "5h6m42s",
  //   },
  //   {
  //     id: "1137117005",
  //     user_name: "ludwig",
  //     title: "POGCHAMPS COVERAGE, TIK TOK TIME, MARIO KART WITH FRIENDS",
  //     created_at: "2021-09-02T19:56:12Z",
  //     url: "https://www.twitch.tv/videos/1137117005",
  //     thumbnail_url:
  //       "https://static-cdn.jtvnw.net/cf_vods/d1m7jfoe9zdc1j/2d5043b59311bf1cade0_ludwig_43530010093_1630612563//thumb/thumb0-%{width}x%{height}.jpg",
  //     view_count: 445795,
  //     duration: "7h3m28s",
  //   },
  // ];

  // return {
  //   resp: {
  //     vods: vods,
  //   },
  //   error: null,
  // };

  let params = userNames.map((u) => {
    return ["user", u];
  });

  const ret = await get("/api/twitch/get_vods", params);
  return ret;
}

export async function getClips(userNames) {
  // let clips = [
  //   {
  //     id: "UgliestFrailGarageNinjaGrumpy-2Vbp2Vo9tOhlPCUT",
  //     url: "https://clips.twitch.tv/UgliestFrailGarageNinjaGrumpy-2Vbp2Vo9tOhlPCUT",
  //     broadcaster_name: "ludwig",
  //     title: "Ludwig on Mizkif and Maya",
  //     view_count: 75088,
  //     created_at: "2021-09-17T00:22:16Z",
  //     thumbnail_url:
  //       "https://clips-media-assets2.twitch.tv/AT-cm%7CJ3q8vdlW6dlaPbToZTwlEw-preview-480x272.jpg",
  //     duration: "30s",
  //   },
  //   {
  //     id: "DarlingBlushingOkapiPrimeMe-gMOq8s7l-4CjXLmr",
  //     url: "https://clips.twitch.tv/DarlingBlushingOkapiPrimeMe-gMOq8s7l-4CjXLmr",
  //     broadcaster_name: "ludwig",
  //     title: "Same Laugh ",
  //     view_count: 46917,
  //     created_at: "2021-09-15T18:06:21Z",
  //     thumbnail_url:
  //       "https://clips-media-assets2.twitch.tv/AT-cm%7C1324612404-preview-480x272.jpg",
  //     duration: "11s",
  //   },
  // ];

  // return {
  //   resp: {
  //     clips: clips,
  //   },
  //   error: null,
  // };

  let params = userNames.map((u) => {
    return ["user", u];
  });

  const ret = await get("/api/twitch/get_clips", params);
  return ret;
}

export async function getSchedule() {
  // let schedule = [
  //   {
  //     date: "2021-09-26",
  //     events: [
  //       {
  //         title: "Opening Ceremony",
  //         start_time: "19:00:00.00-07:00",
  //         user_name: "QTCinderella",
  //         featured_users: ["QTCinderella", "Adeptthebest", "XQCOW"],
  //       },
  //     ],
  //   },
  //   {
  //     date: "2021-09-27",
  //     events: [
  //       {
  //         title: "Pancake breakfast",
  //         start_time: "08:00:00.00-07:00",
  //         user_name: "Nmplol",
  //         featured_users: ["Nmplol"],
  //         vod_link: "",
  //       },
  //       {
  //         title: "Chopped",
  //         start_time: "12:00:00.00-07:00",
  //         user_name: "QTCinderella",
  //         featured_users: ["QTCinderella"],
  //         video_id: "43740004317",
  //         video: {
  //           id: "43740004317",
  //           user_name: "tonytigre",
  //           title: "drawing doodles for new followers! | !insta !doodle",
  //           created_at: "2021-09-18T08:17:04Z",
  //           url: `https://www.twitch.tv/tonytigre`,
  //           thumbnail_url:
  //             "https://static-cdn.jtvnw.net/previews-ttv/live_user_tonytigre-{width}x{height}.jpg",
  //           view_count: 102,
  //         },
  //       },
  //       {
  //         title: "Ghost stories and Smores",
  //         start_time: "19:00:00.00-07:00",
  //         user_name: "WillNeff",
  //         featured_users: ["WillNeff"],
  //       },
  //     ],
  //   },
  //   {
  //     date: "2021-09-28",
  //     events: [
  //       {
  //         title: "French toast breakfast",
  //         start_time: "08:00:00.00-07:00",
  //         user_name: "Nmplol",
  //         featured_users: ["Nmplol"],
  //       },
  //       {
  //         title: "Without a Recipe contest",
  //         start_time: "12:00:00.00-07:00",
  //         user_name: "QTCinderella",
  //         featured_users: ["QTCinderella"],
  //       },
  //     ],
  //   },
  //   {
  //     date: "2021-09-29",
  //     events: [
  //       {
  //         title: "Grand Slam breakfast",
  //         start_time: "08:00:00.00-07:00",
  //         user_name: "Nmplol",
  //         featured_users: ["Nmplol"],
  //       },
  //       {
  //         title: "Closing Ceremony",
  //         start_time: "18:00:00.00-07:00",
  //         user_name: "QTCinderella",
  //         featured_users: ["QTCinderella"],
  //       },
  //     ],
  //   },
  // ];

  // return {
  //   resp: {
  //     dates: schedule,
  //   },
  //   err: null,
  // };

  const ret = await get("/api/shitcamp/get_schedule");
  return ret;
}
