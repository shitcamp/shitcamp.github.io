import React from "react";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Container from "react-bootstrap/Container";
import Card from "react-bootstrap/Card";

import * as utils from "utils";

import "components/vods/Vods.css";

const THUMBNAIL_SIZE = {
  width: 320,
  height: 180,
};

function Vod(props) {
  const {
    title,
    user_name,
    url,
    duration,
    view_count,
    thumbnail_url,
    created_at,
    featured_users,
  } = props.vod;

  var thumbnailUrl = thumbnail_url;
  if (thumbnailUrl) {
    thumbnailUrl = thumbnailUrl.replace("%{width}", THUMBNAIL_SIZE.width);
    thumbnailUrl = thumbnailUrl.replace("%{height}", THUMBNAIL_SIZE.height);
  }

  return (
    <Card>
      <div className="thumbnail">
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
              <React.Fragment>
                <a
                  href={`https://www.twitch.tv/${user}`}
                  text-decoration="none"
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
            text-decoration="none"
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

class Vods extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {}

  render() {
    const { vods } = this.props;

    return (
      <Container>
        <Row xs={1} sm={2} md={3} lg={5} xl={8} className="g-4">
          {vods.map((vod) => (
            <Col
              key={vod.id}
              style={{ display: "flex" } /* make all columns same height */}
            >
              <Vod vod={vod} />
            </Col>
          ))}
        </Row>
      </Container>
    );
  }
}

export default Vods;
