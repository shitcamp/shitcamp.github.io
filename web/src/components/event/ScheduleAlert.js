import { Alert } from "react-bootstrap";

function ScheduleAlert(props) {
  const { isLatestSchedule } = props;

  return !Boolean(isLatestSchedule) ? (
    <Alert variant="warning">
      This is the latest schedule based on the available info. It will be
      updated as more details are released.
    </Alert>
  ) : null;
}

export default ScheduleAlert;
