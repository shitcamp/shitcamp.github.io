import { Alert } from "react-bootstrap";

function ScheduleAlert(props) {
  const { isLatestSchedule } = props;

  return !Boolean(isLatestSchedule) ? (
    <Alert variant="danger">
      This is the latest schedule based on the available info. It will be
      updated as more details are released officially.
    </Alert>
  ) : null;
}

export default ScheduleAlert;
