import React from "react";
import DropdownMultiselect from "react-multiselect-dropdown-bootstrap";
import debounce from "lodash.debounce";

import Videos from "components/videos/Videos";
import LoadingSpinner from "components/spinner/Spinner";

import { getVods } from "apis";

import "components/videos/Vods.css";

const DEBOUNCE_DELAY = 500; // milliseconds to wait for before executing function

class Vods extends React.Component {
  constructor(props) {
    super(props);

    const { userNames } = this.props;

    this.state = {
      isLoading: true,
      userNames: userNames,
      selectedUserNames: userNames,
      pastStreams: [],
    };
  }

  async componentDidMount() {
    const { selectedUserNames } = this.state;

    let ret = await getVods(selectedUserNames);
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      this.setState({
        pastStreams: ret.resp.vods,
      });
    }

    this.setState({
      isLoading: false,
    });
  }

  shouldComponentUpdate(nextProps, nextState) {
    const { userNames: currUsers } = this.props;
    const { userNames: nextUsers } = nextProps;

    if (currUsers.length !== nextUsers.length) {
      return true;
    }

    const {
      isLoading: currIsLoading,
      selectedUserNames: currSelectedUserNames,
      pastStreams: currStreams,
      userNames: currUserNames,
    } = this.state;
    const {
      isLoading: nextIsLoading,
      selectedUserNames: nextSelectedUserNames,
      pastStreams: nextStreams,
      userNames: nextUserNames,
    } = nextState;

    if (currIsLoading !== nextIsLoading) {
      return true;
    }

    if (currUserNames.length !== nextUserNames.length) {
      return true;
    }

    if (currSelectedUserNames.length !== nextSelectedUserNames.length) {
      return true;
    }
    if (currStreams.length !== nextStreams.length) {
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

  async componentDidUpdate(prevProps) {
    const { userNames: oldUserNames } = prevProps;
    const { userNames: newUserNames } = this.props;
    if (newUserNames.length !== oldUserNames.length) {
      this.setState({
        userNames: newUserNames,
        selectedUserNames: newUserNames,
      });
    }

    const { selectedUserNames } = this.state;

    let ret = await getVods(selectedUserNames);
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      this.setState({
        pastStreams: ret.resp.vods,
      });
    }
  }

  handleSelectionChange = (selected) => {
    this.setState({
      selectedUserNames: selected,
    });
  };

  render() {
    const { isLoading, userNames, pastStreams, selectedUserNames } = this.state;

    return (
      <React.Fragment>
        {isLoading ? (
          <LoadingSpinner />
        ) : (
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
            <Videos
              videos={pastStreams}
              emptyErrMsg="No past live streams found"
              onVideoClick={() => {}}
            />
          </React.Fragment>
        )}
      </React.Fragment>
    );
  }
}

export default Vods;
