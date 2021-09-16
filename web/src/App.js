import logo from './logo.svg';
import './App.css';
import TwitchEmbed from './TwitchEmbed';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
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
  );
}

export default App;
