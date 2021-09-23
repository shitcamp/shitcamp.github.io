import { Alert } from "react-bootstrap";

function ScheduleAlert(props) {
  const { isLatestSchedule } = props;

  return !Boolean(isLatestSchedule) ? (
    <Alert variant="warning">
      There have been some changes to the schedule. It will be reflected here
      once it's released.
    </Alert>
  ) : null;
}

export default ScheduleAlert;
