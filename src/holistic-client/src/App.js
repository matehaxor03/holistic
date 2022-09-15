import React from 'react';
import AppContext from './AppContext';
import NavBar from "./components/NavBar";
//import TablePage from './components/TablePage';


class App extends React.Component { 
  data = {"hello":"hello"};

  //state = AppContext;
  //pages = {"TablePage": TablePage};
  //state = Context;

  /*
  viewPage = (pageName, params) => {
    var Zlass = this.pages[pageName];
    var instance = <Zlass id={pageName} params={params}></Zlass>;
    this.setState({...this.state, currentPage: instance});
  }*/
  
  render() {
    if(!this.context.ui) {
      this.context.ui = this;
    }
    
    return (
        <div className="App">
          <NavBar id="menu"></NavBar>
          {this.context.state.currentPage}
        </div>
    );
  }
}

App.contextType = AppContext;

export default App;
