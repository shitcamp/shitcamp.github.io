import React from "react";
import DropdownMultiselect from "react-multiselect-dropdown-bootstrap";

import Videos from "components/videos/Videos";

import { getVods } from "apis";

import "components/videos/Vods.css";

class Vods extends React.Component {
  constructor(props) {
    super(props);

    const { userNames } = this.props;

    this.state = {
      selectedUserNames: userNames,
      pastStreams: [],
    };
  }

  async componentDidMount() {
    const { selectedUserNames } = this.state;

    let ret = await getVods(selectedUserNames);
    if (ret.error != null) {
      console.error(ret.error);
    }
    this.setState({
      pastStreams: ret.resp,
    });
    console.log("Mounted", ret);
  }

  shouldComponentUpdate(nextProps, nextState) {
    const {
      selectedUserNames: currSelectedUserNames,
      pastStreams: currStreams,
    } = this.state;
    const {
      selectedUserNames: nextSelectedUserNames,
      pastStreams: nextStreams,
    } = nextState;

    if (currSelectedUserNames.length !== nextSelectedUserNames.length) {
      return true;
    }
    if (currStreams.length !== nextStreams.length) {
      return true;
    }

    let users = {};
    currSelectedUserNames.forEach((u) => {
      users[u] = true;
    });
    nextSelectedUserNames.forEach((u) => {
      delete users[u];
    });
    if (Object.keys(users).length !== 0) {
      return true;
    }

    let streams = {};
    currStreams.forEach((s) => {
      streams[s.id] = true;
    });
    nextStreams.forEach((s) => {
      delete streams[s.id];
    });
    if (Object.keys(streams).length !== 0) {
      return true;
    }

    return false;
  }

  async componentDidUpdate() {
    console.log("upda");
    const { selectedUserNames } = this.state;

    let ret = await getVods(selectedUserNames);
    if (ret.error != null) {
      console.error(ret.error);
    }
    this.setState({
      pastStreams: ret.resp,
    });
  }

  handleSelectionChange = (selected) => {
    this.setState({
      selectedUserNames: selected,
    });
  };

  render() {
    const { userNames } = this.props;
    const { pastStreams, selectedUserNames } = this.state;

    return (
      <React.Fragment>
        <DropdownMultiselect
          options={userNames}
          selected={selectedUserNames}
          placeholder="Filter by streamer"
          placeholderMultipleChecked="Filter by streamer"
          handleOnChange={(selected) => {
            this.handleSelectionChange(selected);
          }}
          name="streamers"
          className="dropdown  btn-dark"
        />
        <Videos videos={pastStreams} emptyErrMsg="No past live streams found" />
      </React.Fragment>
    );
  }
}

export default Vods;
