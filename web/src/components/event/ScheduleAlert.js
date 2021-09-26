import { Alert } from "react-bootstrap";

function ScheduleAlert(props) {
  const { isLatestSchedule } = props;

  return !Boolean(isLatestSchedule) ? (
    <Alert variant="danger">
      This is the latest schedule based on the available info. It will be
      updated as more details are released officially.
    </Alert>
  ) : (
    <Alert variant="success">
      The schedule has been updated to reflect the latest timings.
    </Alert>
  );
}

export default ScheduleAlert;
