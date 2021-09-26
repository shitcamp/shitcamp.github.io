export function getRelUrl(path) {
  // use prefix for GitHub Pages homepage, and query param to redirect to 404.html to fix routing.
  return process.env.PUBLIC_URL + "/?" + path;
}

export function pad(n) {
  return n > 9 ? "" + n : "0" + n;
}

export function formatAMPM(date) {
  var hours = date.getHours();
  var minutes = date.getMinutes();

  var ampm = hours >= 12 ? "pm" : "am";

  hours = hours % 12;
  hours = hours ? hours : 12; // the hour '0' should be '12'

  minutes = pad(minutes);

  var strTime = hours + ":" + minutes + "" + ampm;
  return strTime;
}

// 2021-09-02T19:56:12Z -> 7h3m28s
export function getElapsedDuration(t) {
  var date1 = Date.parse(t);
  var date2 = Date.now();
  var diffMs = Math.abs(date2 - date1);

  let h = 0,
    m = 0,
    s = 0;

  s = (diffMs / 1000).toFixed(0);

  m = (s / 60).toFixed(0);
  s = (s % 60).toFixed(0);

  h = (m / 60).toFixed(0);
  m = (m % 60).toFixed(0);

  if (h > 0) {
    return `${h}h${m}m${s}s`;
  }
  if (m > 0) {
    return `${m}m${s}s`;
  }
  return `${s}s`;
}

// 7h3m28s -> 7:03:28
// 3m28s -> 3:28
// 8s -> 0:08
export function getDisplayedDuration(d) {
  var h = 0,
    m = 0,
    s = 0;

  var parts = d.split("h");
  if (parts.length > 1) {
    h = parts[0];
    d = parts[1];
  }

  parts = d.split("m");
  if (parts.length > 1) {
    m = parts[0];
    d = parts[1];
  }

  parts = d.split("s");
  s = parts[0];

  if (h > 0) {
    return `${h}:${pad(m)}:${pad(s)}`;
  }
  if (m > 0) {
    return `${m}:${pad(s)}`;
  }
  return `0:${pad(s)}`;
}

// 445795 -> 445.8K
export function getDisplayedViewCount(v) {
  const million = 1000000;
  if (v >= million) {
    var views = v / million;
    if ((v % million) / 100000 < 1) {
      return `${views.toFixed(0)}M views`;
    }

    return `${views.toFixed(1)}M views`;
  }

  const thousand = 1000;
  if (v >= thousand) {
    views = v / thousand;

    if ((v % thousand) / 100 < 1) {
      return `${views.toFixed(0)}K views`;
    }
    return `${views.toFixed(1)}K views`;
  }

  if (v === 1) {
    return `${v} view`;
  }

  return `${v} views`;
}

// 2021-09-02T19:56:12Z -> 2 days ago
export function getRelativeTimeBeforeNow(t) {
  let date1 = Date.parse(t);
  let date2 = Date.now();
  let diffMs = date2 - date1;

  let diffMonths = diffMs / (1000 * 60 * 60 * 24 * 30);
  if (diffMonths >= 1) {
    diffMonths = Math.round(diffMonths);
    if (diffMonths === 1) {
      return `${diffMonths} month ago`;
    }
    return `${diffMonths} months ago`;
  }

  let diffDays = diffMs / (1000 * 60 * 60 * 24);
  if (diffDays >= 1) {
    diffDays = Math.round(diffDays);
    if (diffDays === 1) {
      return `Yesterday`;
    }
    return `${diffDays} days ago`;
  }

  let diffHours = diffMs / (1000 * 60 * 60);
  if (diffHours >= 1) {
    diffHours = Math.round(diffHours);
    if (diffHours === 1) {
      return `${diffHours} hour ago`;
    }
    return `${diffHours} hours ago`;
  }

  let diffMins = Math.round(diffMs / (1000 * 60));
  if (diffMins === 1) {
    return `${diffMins} minute ago`;
  }
  return `${diffMins} minutes ago`;
}

// 2021-09-02T19:56:12Z -> 2 days
export function getRelativeTimeFromNow(t) {
  let date1 = Date.now();
  let date2 = Date.parse(t);
  let diffMs = date2 - date1;

  if (diffMs < 0) {
    return "0 mins";
  }

  let diffMonths = diffMs / (1000 * 60 * 60 * 24 * 30);
  if (diffMonths >= 1) {
    diffMonths = Math.round(diffMonths);
    if (diffMonths === 1) {
      return `${diffMonths} month`;
    }
    return `${diffMonths} months`;
  }

  let diffDays = diffMs / (1000 * 60 * 60 * 24);
  if (diffDays >= 1) {
    diffDays = Math.round(diffDays);
    if (diffDays === 1) {
      return `${diffDays} day`;
    }
    return `${diffDays} days`;
  }

  let diffHours = diffMs / (1000 * 60 * 60);
  if (diffHours >= 1) {
    diffHours = Math.round(diffHours);
    if (diffHours === 1) {
      return `${diffHours} hour`;
    }
    return `${diffHours} hours`;
  }

  let diffMins = Math.round(diffMs / (1000 * 60));
  if (diffMins === 1) {
    return `${diffMins} minute`;
  }
  return `${diffMins} minutes`;
}

// `params` should be a string or an array of pairs. eg. [['k1', 'v1'], ['k2', 1], ['k2', 2]]
export async function get(url, params) {
  url = new URL(process.env.REACT_APP_SERVER_DOMAIN + url);

  if (params && params.length > 0) {
    if (typeof params === "string") {
      url.search = params;
    } else {
      url.search = new URLSearchParams(params).toString();
    }
  }

  const headers = {
    Authorization: "Basic c2hpdGNhbXAtdXNlcjpzb21lLXBhc3N3b3Jk",
  };

  try {
    const response = await fetch(url, { headers });
    if (!response.ok || response.status < 200 || response.status > 299) {
      var err = new Error(
        `API status code ${response.status}: ${response.statusText}`
      );
      return { resp: null, error: err };
    }

    const jsonResponse = await response.json();
    return { resp: jsonResponse.data, error: null };
  } catch (err) {
    var e = new Error(`Shitcamp API error: ${err}`);
    return { resp: null, error: e };
  }
}
