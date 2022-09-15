import React, { useState, createContext} from 'react';
import Dispatcher from './processors/Dispatcher';

var dispatcherObj = new Dispatcher();

const themeObj = {
    fg: "yellow",
    bg: "white"
};

var stateObj = {};
//var params = {dispatchers: dispatcherObj, theme: themeObj, state: stateObj};

const AppContext = createContext({});

export class AppContextProvider extends React.Component {
   
    updateState = (state) => {
      console.log(state);
      this.setState({...state});
    };

    state = {
        dispatchers: dispatcherObj, 
        theme: themeObj, 
        state: stateObj,
        updateState: this.updateState};
  
    render() {
      return (
          //passing the state object as a value prop to all children
          <AppContext.Provider value={this.state}>
              {this.props.children}
          </AppContext.Provider>
      )}
  }


export default AppContext;
