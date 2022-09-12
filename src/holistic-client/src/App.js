import React from 'react';
import NavBar from "./components/NavBar";

export default class App extends React.Component { 
  render() {
    return (
      <div className="App">
        <NavBar id="menu"></NavBar>
        <div id="content">hi4
        </div>
      </div>
    );
  }
}
