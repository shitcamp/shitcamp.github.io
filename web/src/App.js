import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { Container, Navbar, Nav } from "react-bootstrap";

import Home from "pages/home/Home";
import AboutPage from "pages/about/AboutPage";
import Schedule from "pages/schedule/Schedule";
import ClipsPage from "pages/clips/ClipsPage";

import { getRelUrl } from "utils";

import "App.css";
import peepoShy from "assets/peepoShy.gif";
import { ReactComponent as ShitcampSvg } from "assets/logo.svg";

const SHITCAMP_STREAMERS = [
  "AdeptTheBest",
  "AustinShow",
  "Cyr",
  "EsfandTV",
  "HasanAbi",
  "JustaMinx",
  "Jschlatt",
  "Kaceytron",
  "Ludwig",
  "Malena",
  "Myth",
  "Nmplol",
  "QTCinderella",
  "Sodapoppin",
  "WillNeff",
  "xQcOW",
];

class App extends React.PureComponent {
  constructor(props) {
    super(props);

    this.state = {
      userNames: SHITCAMP_STREAMERS,
    };
  }

  async componentDidMount() {
    // let ret = await getUsers();
    // if (ret.error != null) {
    //   console.error(ret.error);
    // } else {
    //   this.setState({
    //     userNames: ret.resp.users,
    //   });
    // }
  }

  render() {
    let { userNames } = this.state;

    return (
      <Router basename={"/" + process.env.PUBLIC_URL}>
        <div className="push">
          {/* <Router> */}
          <Navbar
            collapseOnSelect
            expand="sm"
            fixed="top"
            bg="shitcamp"
            variant="dark"
          >
            <Container>
              <Navbar.Brand href={getRelUrl("/")}>
                <ShitcampSvg className="shitcamp-logo" />
              </Navbar.Brand>

              <Navbar.Toggle aria-controls="responsive-navbar-nav" />

              <Navbar.Collapse id="responsive-navbar-nav">
                <Nav className="me-auto">
                  <Nav.Link href={getRelUrl("/")}>Home</Nav.Link>
                  <Nav.Link href={getRelUrl("/schedule")}>Schedule</Nav.Link>
                  <Nav.Link href={getRelUrl("/about")}>About</Nav.Link>
                  <Nav.Link href={getRelUrl("/clips")}>Top Clips</Nav.Link>
                </Nav>
                <Nav>
                  <Nav.Link
                    href="https://shitcamp.gg/"
                    target="_blank"
                    rel="noreferrer"
                  >
                    Get Merch{" "}
                    <img
                      src={peepoShy}
                      alt="peepoShy"
                      className="peepo-shy-gif"
                    />
                  </Nav.Link>
                </Nav>
              </Navbar.Collapse>
            </Container>
          </Navbar>

          {/* look through the children <Route>s and render the first one that matches the current URL */}
          <Switch>
            <Route
              path={process.env.PUBLIC_URL + "/schedule"}
              component={Schedule}
            />
            <Route
              path={process.env.PUBLIC_URL + "/about"}
              render={() => <AboutPage userNames={userNames} />}
            />
            <Route
              path={process.env.PUBLIC_URL + "/clips"}
              render={() => <ClipsPage userNames={userNames} />}
            />
            <Route path="/" render={() => <Home userNames={userNames} />} />
            <Route path="" render={() => <Home userNames={userNames} />} />
          </Switch>
        </div>
        <div>
          <Navbar bg="shitcamp" variant="dark" className="bottom-navbar">
            <Container>
              <Nav />
              <ShitcampSvg className="shitcamp-logo" />
              <Nav />
            </Container>
          </Navbar>
        </div>
      </Router>
    );
  }
}

export default App;
