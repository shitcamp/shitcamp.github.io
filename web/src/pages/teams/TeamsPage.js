import React from "react";
import { Card, Container, Col, Row } from "react-bootstrap";

import { getTeamsInfo } from "apis";

import "pages/teams/TeamsPage.css";

class TeamsPage extends React.PureComponent {
  constructor(props) {
    super(props);

    this.state = {
      teams: [],
    };
  }

  async componentDidMount() {
    let ret = await getTeamsInfo();
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      const { teams } = ret.resp;

      this.setState({
        teams: teams,
      });
    }
  }

  render() {
    const { teams } = this.state;

    return (
      <Container className="teams-page">
        <h3>Teams</h3>

        {teams.length === 0 ? (
          <h5>Teams are not available yet</h5>
        ) : (
          <Row className="g-4">
            {teams.map((team) => (
              <Col xs={12} sm={6} lg={4}>
                <Card>
                  <Card.Header as="h5">{team.name}</Card.Header>
                  <Card.Body>
                    <Card.Text>
                      <ul className="no-bullets">
                        {team.users.map((user) => (
                          <li>
                            <a
                              href={`https://www.twitch.tv/${user}`}
                              target="_blank"
                              rel="noreferrer"
                              className="link"
                            >
                              {user}
                            </a>
                          </li>
                        ))}
                      </ul>
                    </Card.Text>
                  </Card.Body>
                </Card>
              </Col>
            ))}
          </Row>
        )}
      </Container>
    );
  }
}

export default TeamsPage;
