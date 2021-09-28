import React from "react";
import { Card } from "react-bootstrap";

import * as utils from "utils";

import "components/event/EventCard.css";

const THUMBNAIL_SIZE = {
  width: 320,
  height: 180,
};

function EventCard(props) {
  // onVideoClick
  const { event, displayStartDate } = props;
  const {
    title: event_title,
    start_time,
    user_name,
    featured_users: event_featured_users,
    video_id,
    thumbnail_url,
    vod,
    description,
  } = event;

  let videoUrl = `https://twitch.tv/${user_name}`;
  if (video_id) {
    videoUrl = `https://twitch.tv/videos/${video_id}`;
  }

  let featuredUsers = event_featured_users;
  // let handleVideoClick = () => {};

  let duration = "";
  let viewCount = "";
  let createdAt = "";
  let thumbnailUrl = thumbnail_url;

  if (vod) {
    if (vod.url) {
      videoUrl = vod.url;
    }

    if (Array.isArray(vod.featured_users) && vod.featured_users.length > 0) {
      featuredUsers = vod.featured_users;
    }

    // if (vod.id) {
    // handleVideoClick = (e) => onVideoClick(e, vod.id);
    // }

    duration = vod.duration;
    viewCount = vod.view_count;
    if (Date.parse(vod.created_at) > 0) {
      createdAt = vod.created_at;
    }

    if (vod.thumbnail_url) {
      thumbnailUrl = vod.thumbnail_url;

      thumbnailUrl = thumbnailUrl.replace("%{width}", THUMBNAIL_SIZE.width);
      thumbnailUrl = thumbnailUrl.replace("{width}", THUMBNAIL_SIZE.width);

      thumbnailUrl = thumbnailUrl.replace("%{height}", THUMBNAIL_SIZE.height);
      thumbnailUrl = thumbnailUrl.replace("{height}", THUMBNAIL_SIZE.height);
    }
  }

  if (!thumbnailUrl) {
    thumbnailUrl = `https://static-cdn.jtvnw.net/previews-ttv/live_user_${user_name}-${THUMBNAIL_SIZE.width}x${THUMBNAIL_SIZE.height}.jpg`;
  }

  let thumbnailComponent = (
    <div className="thumbnail">
      {/* onClick={handleVideoClick} */}
      <a href={videoUrl} target="_blank" rel="noreferrer" className="link">
        <Card.Img variant="top" alt="thumbnail" src={thumbnailUrl} />
      </a>

      {!!duration && (
        <small className="indicator duration">
          {utils.getDisplayedDuration(duration)}
        </small>
      )}

      {!!viewCount && (
        <small className="indicator views">
          {utils.getDisplayedViewCount(viewCount)}
        </small>
      )}

      <small className="indicator created-at">
        {!!createdAt ? (
          utils.getRelativeTimeBeforeNow(createdAt)
        ) : (
          <React.Fragment>
            Starts in {utils.getRelativeTimeFromNow(start_time)}
          </React.Fragment>
        )}
      </small>
    </div>
  );

  let startTimeStr = "";
  const d = new Date(start_time);
  if (displayStartDate) {
    startTimeStr += d.toDateString().split(/ (.+)/)[1] + ", ";
  }
  startTimeStr += utils.formatAMPM(d);

  return (
    <Card>
      {thumbnailComponent}

      <Card.Body>
        <Card.Title className="card-title">
          <h6>
            <a
              href={videoUrl}
              target="_blank"
              rel="noreferrer"
              className="link"
            >
              {event_title}
            </a>
            {event_title === "PJ Party" && (
              <img src="./jammies.gif" alt="Jammies" className="jammies-gif" />
            )}
          </h6>
        </Card.Title>

        <Card.Text>
          {Date.parse(start_time) > Date.now() && (
            <React.Fragment>
              <small>
                <span className="description">Starts at</span>: {startTimeStr}
              </small>
              <br />
            </React.Fragment>
          )}
          {Array.isArray(featuredUsers) && featuredUsers.length > 0 && (
            <small>
              {description && description !== "" && (
                <React.Fragment>
                  <span className="description">Description</span>:{" "}
                  {description}
                  {event_title === "PJ Party" && (
                    <React.Fragment>
                      <br />
                      {"Get your own matching PJs at the "}
                      <b>
                        <a
                          href="https://shitcamp.gg/product/peepo-pjs"
                          target="_blank"
                          rel="noreferrer"
                        >
                          Shitcamp Merch Store
                        </a>
                      </b>
                    </React.Fragment>
                  )}
                  {event_title === "French toast breakfast" && (
                    <React.Fragment>
                      {" "}
                      <img src="./YEP.png" alt="YEP" className="yep-image" />
                    </React.Fragment>
                  )}
                  <br />
                </React.Fragment>
              )}
              <span className="description">Featuring</span>:{" "}
              {featuredUsers.map((user, i, users) => (
                <React.Fragment key={user}>
                  <a
                    href={`https://www.twitch.tv/${user}`}
                    target="_blank"
                    rel="noreferrer"
                    className="link"
                  >
                    {user}
                  </a>
                  {i !== users.length - 1 ? ", " : ""}
                </React.Fragment>
              ))}
            </small>
          )}
        </Card.Text>
      </Card.Body>

      <Card.Footer>
        <small>
          <a
            href={`https://www.twitch.tv/${user_name}`}
            target="_blank"
            rel="noreferrer"
            className="link text-muted"
          >
            {user_name ? user_name : "TBA"}
          </a>
        </small>
      </Card.Footer>
    </Card>
  );
}

export default EventCard;
