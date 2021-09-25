import React from "react";
import { Card, Container, Row, Col } from "react-bootstrap";
import DropdownMultiselect from "react-multiselect-dropdown-bootstrap";
import debounce from "lodash.debounce";

import { getClips } from "apis";
import * as utils from "utils";

import "components/videos/Clips.css";

const THUMBNAIL_SIZE = {
  width: 320,
  height: 180,
};

function ClipCard(props) {
  const { clip, onVideoClick } = props;
  const {
    title,
    broadcaster_name,
    url,
    duration,
    view_count,
    thumbnail_url,
    created_at,
  } = clip;

  var thumbnailUrl = thumbnail_url;
  if (thumbnailUrl) {
    thumbnailUrl = thumbnailUrl.replace("%{width}", THUMBNAIL_SIZE.width);
    thumbnailUrl = thumbnailUrl.replace("{width}", THUMBNAIL_SIZE.width);

    thumbnailUrl = thumbnailUrl.replace("%{height}", THUMBNAIL_SIZE.height);
    thumbnailUrl = thumbnailUrl.replace("{height}", THUMBNAIL_SIZE.height);
  }

  return (
    <Card>
      <div className="thumbnail" onClick={(e) => onVideoClick(e, clip)}>
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
          {utils.getRelativeTimeBeforeNow(created_at)}
        </small>
      </div>
      {/* <Image src={thumbnailUrl} thumbnail /> */}

      <Card.Body>
        <Card.Title className="card-title truncate">
          <a href={url} target="_blank" rel="noreferrer" className="link">
            <h6>{title}</h6>
          </a>
        </Card.Title>
      </Card.Body>

      <Card.Footer>
        <small>
          <a
            href={`https://www.twitch.tv/${broadcaster_name}`}
            target="_blank"
            rel="noreferrer"
            className="link text-muted"
          >
            {broadcaster_name}
          </a>
        </small>
      </Card.Footer>
    </Card>
  );
}

const DEBOUNCE_DELAY = 500;

class Clips extends React.Component {
  constructor(props) {
    super(props);

    const { userNames } = this.props;

    this.state = {
      selectedUserNames: userNames,
      videos: [],
    };
  }

  async componentDidMount() {
    const { selectedUserNames } = this.state;

    let ret = await getClips(selectedUserNames);
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      const { clips } = ret.resp;

      this.setState({
        videos: clips,
      });
    }
  }

  shouldComponentUpdate(nextProps, nextState) {
    const { userNames: currUsers } = this.props;
    const { userNames: nextUsers } = nextProps;

    if (currUsers.length !== nextUsers.length) {
      return true;
    }

    const { selectedUserNames: currSelectedUserNames, videos: currClips } =
      this.state;
    const { selectedUserNames: nextSelectedUserNames, videos: nextClips } =
      nextState;

    if (currSelectedUserNames.length !== nextSelectedUserNames.length) {
      return true;
    }
    if (currClips != null && currClips.length !== nextClips.length) {
      return true;
    }

    let selected = {};
    currSelectedUserNames.forEach((u) => {
      selected[u] = true;
    });
    nextSelectedUserNames.forEach((u) => {
      delete selected[u];
    });
    if (Object.keys(selected).length !== 0) {
      return true;
    }

    let clips = {};
    currClips.forEach((s) => {
      clips[s.id] = true;
    });
    nextClips.forEach((s) => {
      delete clips[s.id];
    });
    if (Object.keys(clips).length !== 0) {
      return true;
    }

    return false;
  }

  async componentDidUpdate(prevProps) {
    const { userNames: oldUserNames } = prevProps;
    const { userNames: newUserNames } = this.props;
    if (newUserNames.length !== oldUserNames.length) {
      this.setState({
        selectedUserNames: newUserNames,
      });
    }

    const { selectedUserNames } = this.state;

    let ret = await getClips(selectedUserNames);
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      this.setState({
        videos: ret.resp.clips,
      });
    }
  }

  // TODO: refactor to HOC?
  handleSelectionChange = (selected) => {
    this.setState({
      selectedUserNames: selected,
    });
  };

  render() {
    const { userNames, onClipClick } = this.props;
    const { videos, selectedUserNames } = this.state;

    return (
      <React.Fragment>
        {userNames.length !== 0 && (
          <DropdownMultiselect
            options={userNames}
            selected={selectedUserNames}
            placeholder="Filter by streamer"
            placeholderMultipleChecked="Filter by streamer"
            handleOnChange={debounce((selected) => {
              this.handleSelectionChange(selected);
            }, DEBOUNCE_DELAY)}
            name="streamers"
            className="dropdown"
          />
        )}
        {videos.length > 0 ? (
          <Container>
            <p>{"Select a clip to watch"}</p>
            <Row xs={1} sm={2} md={3} lg={5} xl={8} className="g-4">
              {videos.map((c) => (
                <Col
                  key={c.id}
                  style={{ display: "flex" } /* make all columns same height */}
                >
                  <ClipCard clip={c} onVideoClick={onClipClick} />
                </Col>
              ))}
            </Row>
          </Container>
        ) : (
          <h5>No clips found</h5>
        )}
      </React.Fragment>
    );
  }
}

export default Clips;
