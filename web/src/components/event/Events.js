import { Row, Col } from "react-bootstrap";

import EventCard from "components/event/EventCard";

function Events(props) {
  const { events, onVideoClick, displayStartDate } = props;

  return (
    <Row xs={1} sm={2} md={3} lg={5} xl={8} className="g-4">
      {events.map((e, idx) => (
        <Col
          key={idx}
          style={{ display: "flex" } /* make all columns same height */}
        >
          <EventCard
            event={e}
            onVideoClick={onVideoClick}
            displayStartDate={displayStartDate}
          />
        </Col>
      ))}
    </Row>
  );
}

export default Events;
