import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Container from "react-bootstrap/Container";
import Button from "react-bootstrap/Button";
import ButtonToolbar from "react-bootstrap/ButtonToolbar";
import { LinkContainer } from "react-router-bootstrap";

import TwitchEmbed from "./TwitchEmbed";

const Home = () => <span>Home</span>;

const About = () => <span>About</span>;

const Users = () => <span>Users</span>;

function App() {
  return (
    <React.Fragment>
      <Router>
        <Container className="p-3">
          <h1 className="header">Welcome To React-Bootstrap</h1>
          <h2>
            Current Page is{" "}
            {/* look through the children <Route>s and render the first one that matches the current URL */}
            <Switch>
              <Route path="/about" component={About} />
              <Route path="/users" component={Users} />
              <Route path="/" component={Home} />
            </Switch>
          </h2>

          <h2>
            Navigate to{" "}
            <ButtonToolbar className="custom-btn-toolbar">
              <LinkContainer to="/">
                <Button>Home</Button>
              </LinkContainer>

              <LinkContainer to="/about">
                <Button>About</Button>
              </LinkContainer>

              <LinkContainer to="/users">
                <Button>Users</Button>
              </LinkContainer>
            </ButtonToolbar>
          </h2>
        </Container>
      </Router>

      <div className="App">
        <header className="App-header">
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>

        <TwitchEmbed
          id="qtcinderella-1"
          channel="ludwig" // name of the chat room and channel to stream
          // video="1144988822" // ID of VOD to play
          // chat="mobile" // TODO: detect mobile devices
        />
      </div>
    </React.Fragment>
  );
}

export default App;
