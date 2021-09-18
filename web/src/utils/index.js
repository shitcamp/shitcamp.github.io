function pad(n) {
  return n > 9 ? "" + n : "0" + n;
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

// 2021-09-02T19:56:12Z ->
export function getRelativeTime(t) {
  var date1 = Date.parse(t);
  var date2 = Date.now();
  var diffMs = Math.abs(date2 - date1);

  const diffMonths = (diffMs / (1000 * 60 * 60 * 24 * 30)).toFixed(0);
  if (diffMonths >= 1) {
    if (diffMonths === 1) {
      return `${diffMonths} month ago`;
    }
    return `${diffMonths} months ago`;
  }

  const diffDays = (diffMs / (1000 * 60 * 60 * 24)).toFixed(0);
  if (diffDays >= 1) {
    if (diffDays === 1) {
      return `Yesterday`;
    }
    return `${diffDays} days ago`;
  }

  const diffHours = (diffMs / (1000 * 60 * 60)).toFixed(0);
  if (diffHours >= 1) {
    if (diffHours === 1) {
      return `${diffHours} hour ago`;
    }
    return `${diffHours} hours ago`;
  }

  const diffMins = (diffMs / (1000 * 60)).toFixed(0);
  if (diffMins === 1) {
    return `${diffMins} minute ago`;
  }
  return `${diffMins} minutes ago`;
}

// `params` should be a string or an array of pairs. eg. [['k1', 'v1'], ['k2', 1], ['k2', 2]]
export async function get(url, params) {
  url = new URL(url);

  if (params && params.length > 0) {
    if (typeof params === "string") {
      url.search = params;
    } else {
      url.search = new URLSearchParams(params).toString();
    }
  }

  try {
    const response = await fetch(url);
    if (response.status < 200 || response.status > 299) {
      var err = new Error(
        `API status code ${response.status}: ${response.statusText}`
      );
      return null, err;
    }

    const jsonResponse = await response.json();
    // console.log(jsonResponse);
    return jsonResponse, null;
  } catch (err) {
    var e = new Error(`API error: ${err}`);
    return null, e;
  }
}
