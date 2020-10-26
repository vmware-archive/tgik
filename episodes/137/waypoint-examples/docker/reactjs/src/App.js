import React from 'react';
import logo from './logo.svg';
import language from './language.svg';
import hashi from './hashi.svg';
import './App.css';

function App() {
  return ( 
    <div class="container">
      <header>
        <a href="https://waypointproject.io" class="logo">
          <img src={logo} />
        </a>
      </header>
      <section class="content">
        <div class="language-icon">
          <img src={language} />
        </div>
        <h1>This React app was deployed with Waypoint.</h1>
        <p>
          Try making a change to this text locally and run <code>waypoint up</code> again to see it.
        </p>
        <p>
          Read the <a href="https://waypointproject.io/docs">documentation</a> for more about Waypoint.
        </p>
      </section>
      <footer>
        <a href="https://hashicorp.com" class="hashi">
          <img src={hashi} />
        </a>
      </footer>
    </div>
  );
}

export default App;
