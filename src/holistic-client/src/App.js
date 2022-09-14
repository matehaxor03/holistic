import React from 'react';
import NavBar from "./components/NavBar";
import TablePage from './components/TablePage';
import Context from './Context';
import { ThemeProvider } from 'styled-components';

class App extends React.Component { 
  data = {"hello":"hello"};
  pages = {"TablePage": TablePage};
  state = {currentPage: null};

  viewPage = (pageName, params) => {
    var Zlass = this.pages[pageName];
    var instance = <Zlass id={pageName} context={this} params={params}></Zlass>;
    this.setState({...this.state, currentPage: instance});
  }
  
  render() {
    return (
      <ThemeProvider theme={theme}>
      <Context.Provider value={this}>
        <div className="App">
          <NavBar id="menu" context={this}></NavBar>
          {this.state.currentPage}
        </div>
      </Context.Provider>
      </ThemeProvider>
    );
  }
}

const theme = {
  fg: "yellow",
  bg: "white"
};

export default App;
