import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";

import TwitchEmbed from "./components/TwitchEmbed/TwitchEmbed";
import Vods from "./components/Vods/Vods";
import { getVods } from "./apis/api";

import "./App.css";
import peepoShy from "./assets/peepoShy.gif";
import { ReactComponent as ShitcampSvg } from "./assets/logo.svg";

const Home = () => <span>Home</span>;

const About = () => <span>About</span>;

const Users = () => <span>Users</span>;

function getRelUrl(path) {
  // use prefix for GitHub Pages homepage, and query param to redirect to 404.html to fix routing.
  return process.env.PUBLIC_URL + "/?" + path;
}

function App() {
  console.log("v4");

  return (
    <Router basename={"/" + process.env.PUBLIC_URL}>
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
            Shitcamp <h6>unofficial</h6>
          </Navbar.Brand>

          <Navbar.Toggle aria-controls="responsive-navbar-nav" />

          <Navbar.Collapse id="responsive-navbar-nav">
            <Nav className="me-auto">
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
                <img src={peepoShy} alt="" className="peepo-shy-gif" />
              </Nav.Link>
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>

      {"Shit's gettin' litty, come thru <3"}

      {/* look through the children <Route>s and render the first one that matches the current URL */}
      <Switch>
        <Route path={process.env.PUBLIC_URL + "/schedule"} component={Users} />
        <Route path={process.env.PUBLIC_URL + "/about"} component={About} />
        <Route path={process.env.PUBLIC_URL + "/clips"} component={Users} />
        <Route path="/" component={Home} />
        <Route path="" component={Home} />
      </Switch>

      <Vods vods={getVods()} />
      <TwitchEmbed
        id="qtcinderella-1"
        channel="ludwig" // name of the chat room and channel to stream
        // video="1144988822" // ID of VOD to play
        // chat="mobile" // TODO: detect mobile devices
      />

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

export default App;
