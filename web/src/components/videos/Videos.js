import React from "react";
import { Row, Col, Container, Card } from "react-bootstrap";

import * as utils from "utils";

import "components/videos/Videos.css";

const THUMBNAIL_SIZE = {
  width: 320,
  height: 180,
};

function VideoCard(props) {
  const { video, onVideoClick } = props;
  const {
    title,
    user_name,
    url,
    duration,
    view_count,
    thumbnail_url,
    created_at,
    featured_users,
  } = video;

  var thumbnailUrl = thumbnail_url;
  if (thumbnailUrl) {
    thumbnailUrl = thumbnailUrl.replace("%{width}", THUMBNAIL_SIZE.width);
    thumbnailUrl = thumbnailUrl.replace("{width}", THUMBNAIL_SIZE.width);

    thumbnailUrl = thumbnailUrl.replace("%{height}", THUMBNAIL_SIZE.height);
    thumbnailUrl = thumbnailUrl.replace("{height}", THUMBNAIL_SIZE.height);
  }

  return (
    <Card>
      <div className="thumbnail" onClick={onVideoClick}>
        <a href={url} target="_blank" rel="noreferrer" className="link">
          <Card.Img variant="top" alt="thumbnail" src={thumbnailUrl} />
        </a>
        <small className="indicator duration">
          {utils.getDisplayedDuration(duration)}
        </small>
        <small className="indicator views">
          {utils.getDisplayedViewCount(view_count)}
        </small>
        <small className="indicator createdAt">
          {utils.getRelativeTime(created_at)}
        </small>
      </div>
      {/* <Image src={thumbnailUrl} thumbnail /> */}

      <Card.Body>
        <Card.Title className="card-title truncate">
          <a href={url} target="_blank" rel="noreferrer" className="link">
            {title}
          </a>
        </Card.Title>

        {Array.isArray(featured_users) && featured_users.length > 0 ? (
          <Card.Text className="truncate">
            Featuring:{" "}
            {featured_users.map((user, i, users) => (
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
          </Card.Text>
        ) : null}
      </Card.Body>

      <Card.Footer>
        <small>
          <a
            href={`https://www.twitch.tv/${user_name}`}
            target="_blank"
            rel="noreferrer"
            className="link text-muted"
          >
            {user_name}
          </a>
        </small>
      </Card.Footer>
    </Card>
  );
}

class Videos extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {}

  render() {
    const { videos, titleMsg, emptyErrMsg, onVideoClick } = this.props;

    return (
      <React.Fragment>
        {Array.isArray(videos) && videos.length > 0 ? (
          <React.Fragment>
            {titleMsg != null && <h5>{titleMsg}</h5>}
            <Container>
              <Row xs={1} sm={2} md={3} lg={5} xl={8} className="g-4">
                {videos.map((v) => (
                  <Col
                    key={v.id}
                    style={
                      { display: "flex" } /* make all columns same height */
                    }
                  >
                    <VideoCard video={v} onVideoClick={onVideoClick} />
                  </Col>
                ))}
              </Row>
            </Container>
          </React.Fragment>
        ) : (
          <h5>{emptyErrMsg}</h5>
        )}
      </React.Fragment>
    );
  }
}

export default Videos;
