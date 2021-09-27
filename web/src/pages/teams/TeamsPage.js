import React from "react";
import { Card, Container, Col, Row } from "react-bootstrap";

import { getTeamsInfo } from "apis";

import "pages/teams/TeamsPage.css";
import LoadingSpinner from "components/spinner/Spinner";

class TeamsPage extends React.PureComponent {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: true,
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

    this.setState({
      isLoading: false,
    });
  }

  render() {
    const { isLoading, teams } = this.state;

    return (
      <Container className="teams-page">
        <h3>Teams</h3>

        {isLoading ? (
          <LoadingSpinner />
        ) : teams.length === 0 ? (
          <h5>Teams are not available</h5>
        ) : (
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
        )}
      </Container>
    );
  }
}

export default TeamsPage;
