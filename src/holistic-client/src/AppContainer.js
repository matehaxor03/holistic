import React, { Fragment } from 'react';
import { ThemeProvider } from 'styled-components';
import App from './App';
import AppContext, { AppContextProvider } from './AppContext';
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

  /*
  componentDidMount() {
    this.context.app = this;
    console.log(this.context);
  }*/
  
  render() {
    return (
      <AppContextProvider> 
   
          <App></App>
    
      </AppContextProvider>
    );
  }
}

//AppContainer.contextType = AppContext;

export default AppContainer;
