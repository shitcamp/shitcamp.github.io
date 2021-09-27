import { Spinner } from "react-bootstrap";

function LoadingSpinner() {
  return (
    <Spinner animation="border" role="status" className="mx-auto">
      <span className="sr-only visually-hidden">Loading...</span>
    </Spinner>
  );
}

export default LoadingSpinner;
