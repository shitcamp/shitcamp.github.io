import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";

import Home from "pages/home/Home";
import About from "pages/about/About";
import Schedule from "pages/schedule/Schedule";
import Clips from "pages/clips/Clips";

import "App.css";
import peepoShy from "assets/peepoShy.gif";
import { ReactComponent as ShitcampSvg } from "assets/logo.svg";

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

      {/* look through the children <Route>s and render the first one that matches the current URL */}
      <Switch>
        <Route
          path={process.env.PUBLIC_URL + "/schedule"}
          component={Schedule}
        />
        <Route path={process.env.PUBLIC_URL + "/about"} component={About} />
        <Route path={process.env.PUBLIC_URL + "/clips"} component={Clips} />
        <Route path="/" component={Home} />
        <Route path="" component={Home} />
      </Switch>

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
