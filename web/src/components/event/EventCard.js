import React from "react";
import { Card } from "react-bootstrap";

import * as utils from "utils";

import "components/event/EventCard.css";

const THUMBNAIL_SIZE = {
  width: 320,
  height: 180,
};

function EventCard(props) {
  const { event, displayStartDate } = props;
  const {
    title: event_title,
    start_time,
    user_name,
    featured_users: event_featured_users,
    video_id,
    vod,
    description,
  } = event;

  let featuredUsers = event_featured_users;

  let video_url = `https://twitch.tv/${user_name}`;
  if (video_id) {
    video_url = `https://twitch.tv/videos/${video_id}`;
  }

  let thumbnailUrl = `https://static-cdn.jtvnw.net/previews-ttv/live_user_${user_name}-${THUMBNAIL_SIZE.width}x${THUMBNAIL_SIZE.height}.jpg`;

  let thumbnailComponent = (
    <div className="thumbnail">
      <a href={video_url} target="_blank" rel="noreferrer" className="link">
        <Card.Img variant="top" alt="thumbnail" src={thumbnailUrl} />
      </a>
      <small className="indicator created-at">
        Starts in {utils.getRelativeTimeFromNow(start_time)}
      </small>
    </div>
  );

  if (vod) {
    if (Array.isArray(vod.featured_users) && vod.featured_users.length > 0) {
      featuredUsers = vod.featured_users;
    }

    video_url = vod.url;

    if (vod.thumbnail_url) {
      thumbnailUrl = vod.thumbnail_url;

      thumbnailUrl = thumbnailUrl.replace("%{width}", THUMBNAIL_SIZE.width);
      thumbnailUrl = thumbnailUrl.replace("{width}", THUMBNAIL_SIZE.width);

      thumbnailUrl = thumbnailUrl.replace("%{height}", THUMBNAIL_SIZE.height);
      thumbnailUrl = thumbnailUrl.replace("{height}", THUMBNAIL_SIZE.height);
    }

    thumbnailComponent = (
      <div className="thumbnail">
        {/* onClick={(e) => onVideoClick(e, vod.id)}> */}
        <a href={video_url} target="_blank" rel="noreferrer" className="link">
          <Card.Img variant="top" alt="thumbnail" src={thumbnailUrl} />
        </a>
        <small className="indicator duration">
          {utils.getDisplayedDuration(vod.duration)}
        </small>
        <small className="indicator views">
          {utils.getDisplayedViewCount(vod.view_count)}
        </small>
        <small className="indicator created-at">
          {utils.getRelativeTimeBeforeNow(vod.created_at)}
        </small>
      </div>
    );
  }

  let startTimeStr = "";
  const d = new Date(start_time);
  if (displayStartDate) {
    startTimeStr += d.toDateString().split(/ (.+)/)[1] + ", ";
  }
  startTimeStr += d.toTimeString();

  return (
    <Card>
      {thumbnailComponent}

      <Card.Body>
        <Card.Title className="card-title">
          <a href={video_url} target="_blank" rel="noreferrer" className="link">
            <h6>{event_title}</h6>
          </a>
        </Card.Title>

        <Card.Text>
          {Date.parse(start_time) > Date.now() && (
            <React.Fragment>
              <small>
                <span className="sub-title">Starts at</span>: {startTimeStr}
              </small>
              <br />
            </React.Fragment>
          )}
          {Array.isArray(featuredUsers) && featuredUsers.length > 0 && (
            <small>
              {description && description !== "" && (
                <React.Fragment>
                  <span className="sub-title">Description</span>: {description}
                  <br />
                </React.Fragment>
              )}
              <span className="sub-title">Featuring</span>:{" "}
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
