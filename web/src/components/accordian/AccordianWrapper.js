import { Accordion } from "react-bootstrap";

import "components/accordian/AccordianWrapper.css";

function AccordianWrapper(props) {
  const { title, children } = props;

  return (
    <Accordion defaultActiveKey="0" className="accordian">
      <Accordion.Item eventKey="0">
        <Accordion.Header>
          <h5>{title}</h5>
        </Accordion.Header>
        <Accordion.Body>{children}</Accordion.Body>
      </Accordion.Item>
    </Accordion>
  );
}

export default AccordianWrapper;
