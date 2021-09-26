import { get, getElapsedDuration } from "utils";

export async function getUsers() {
  const ret = await get("/api/shitcamp/get_streamer_names");
  return ret;
}

export async function getLiveStreams(userNames) {
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
  let params = userNames.map((u) => {
    return ["user", u];
  });

  const ret = await get("/api/twitch/get_vods", params);
  return ret;
}

export async function getClips(userNames) {
  let params = userNames.map((u) => {
    return ["user", u];
  });

  const ret = await get("/api/twitch/get_clips", params);
  return ret;
}

export async function getSchedule() {
  const ret = await get("/api/shitcamp/get_schedule");
  return ret;
}

export async function getTeamsInfo() {
  // const resp = {
  //   teams: [
  //     {
  //       name: "Team A",
  //       users: [
  //         "AdeptTheBest",
  //         "AustinShow",
  //         "Cyr",
  //         "EsfandTV",
  //         "HasanAbi",
  //         "JustaMinx",
  //         "Jschlatt",
  //         "Kaceytron",
  //       ],
  //     },
  //     {
  //       name: "Team B",
  //       users: [
  //         "Ludwig",
  //         "Malena",
  //         "Myth",
  //         "Nmplol",
  //         "QTCinderella",
  //         "Sodapoppin",
  //         "WillNeff",
  //         "xQcOW",
  //       ],
  //     },
  //   ],
  // };

  // return {
  //   resp: resp,
  //   err: null,
  // };

  const ret = await get("/api/shitcamp/get_teams_info");
  return ret;
}
