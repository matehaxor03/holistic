import React from 'react';
import NavBar from "./components/NavBar";
import ContentPage from "./components/ContentPage";
import TablePage from './components/TablePage';

export default class App extends React.Component { 
  pages = {"TablePage": TablePage};
  state = {currentPage: <ContentPage id="content" context={this}></ContentPage>};

  viewPage = (pageName, params) => {
    var Zlass = this.pages[pageName];
    var instance = <Zlass id={pageName} context={this} params={params}></Zlass>;
    this.setState({...this.state, currentPage: instance});
  }
  
  render() {
    return (
      <div className="App">
        <NavBar id="menu" context={this}></NavBar>
        {this.state.currentPage}
      </div>
    );
  }
}
