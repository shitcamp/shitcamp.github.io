import React from "react";
import { Card, Container, Col, Row, ListGroup } from "react-bootstrap";

import { getTeamsInfo } from "apis";

import "pages/teams/TeamsPage.css";
import LoadingSpinner from "components/spinner/Spinner";

class TeamsPage extends React.PureComponent {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: true,
      teams: [],
      teamResults: [],
    };
  }

  async componentDidMount() {
    let ret = await getTeamsInfo();
    if (ret.error != null) {
      console.error(ret.error);
    } else {
      const { teams, team_results } = ret.resp;

      this.setState({
        teams: teams,
        teamResults: team_results,
      });
    }

    this.setState({
      isLoading: false,
    });
  }

  render() {
    const { isLoading, teams, teamResults } = this.state;

    return (
      <Container className="teams-page">
        <h3>Teams</h3>

        {isLoading ? (
          <LoadingSpinner />
        ) : teams.length === 0 ? (
          <h5>Teams are not available</h5>
        ) : (
          <React.Fragment>
            <Row className="g-4">
              {teams.map((team) => (
                <Col xs={12} sm={6} lg={4}>
                  <Card>
                    <Card.Header as="h5">{team.name}</Card.Header>
                    <Card.Img
                      variant="top"
                      alt="thumbnail"
                      src={team.thumbnail_url}
                    />
                    <Card.Body>
                      <Card.Text>
                        <ul className="no-bullets">
                          <li>
                            <Card.Subtitle>
                              <a
                                href={`https://www.twitch.tv/${team.captain}`}
                                target="_blank"
                                rel="noreferrer"
                                className="link"
                              >
                                {team.captain} (captain)
                              </a>
                            </Card.Subtitle>
                          </li>
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

            <br />
            <h4>Results</h4>
            <Row className="g-4">
              {teamResults.map((teamResult) => (
                <Col xs={12} sm={6} lg={4}>
                  <Card>
                    <Card.Header as="h5">{teamResult.team_name}</Card.Header>
                    <Card.Body>
                      <Card.Subtitle className="mb-2">
                        Total points: {teamResult.total_points}
                      </Card.Subtitle>
                      <Card.Text>
                        {teamResult.results.map((result) => (
                          <ListGroup>
                            <ListGroup.Item
                              variant="warning"
                              action
                              href={result.vod_url}
                              target="_blank"
                              rel="noreferrer"
                            >
                              <b>{result.contest_name}</b>: {result.points}{" "}
                              points
                              <br />
                              {result.description}
                            </ListGroup.Item>
                          </ListGroup>
                        ))}
                      </Card.Text>
                    </Card.Body>
                  </Card>
                </Col>
              ))}
            </Row>
          </React.Fragment>
        )}
      </Container>
    );
  }
}

export default TeamsPage;
