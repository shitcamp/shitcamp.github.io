import React from "react";
import { Alert } from "react-bootstrap";

import * as utils from "utils";

function ScheduleAlert(props) {
  const { isLatestSchedule, lastUpdated } = props;

  const d = new Date(lastUpdated);
  let updatedTimeStr = d.toDateString().split(/ (.+)/)[1] + ", ";
  updatedTimeStr += utils.formatAMPM(d);

  return !Boolean(isLatestSchedule) ? (
    <Alert variant="danger">
      This is the latest schedule based on the available info. It will be
      updated as more details are released officially.
    </Alert>
  ) : (
    <Alert variant="success">
      The schedule has been updated to reflect the latest info.{" "}
      {lastUpdated && lastUpdated !== "" && (
        <React.Fragment>Last updated at: {updatedTimeStr}</React.Fragment>
      )}
    </Alert>
  );
}

export default ScheduleAlert;
