import React from 'react';
import { ThemeProvider } from 'styled-components';
import App from './App';
import AppContext from './AppContext';
//import TablePage from './components/TablePage';


class AppContainer extends React.Component { 
 

  //state = AppContext;
  //pages = {"TablePage": TablePage};
  //state = Context;

  /*
  viewPage = (pageName, params) => {
    var Zlass = this.pages[pageName];
    var instance = <Zlass id={pageName} params={params}></Zlass>;
    this.setState({...this.state, currentPage: instance});
  }*/
  componentDidMount() {
    this.context.app = this;
    console.log(this.context);
  }
  
  render() {
    return (
      <AppContext.Provider value={this.context}>
      <ThemeProvider theme={this.context.theme}>
    
      <App></App>
      </ThemeProvider>
      </AppContext.Provider>
    );
  }
}

AppContainer.contextType = AppContext;


export default AppContainer;
