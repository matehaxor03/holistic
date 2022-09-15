import React, { createContext} from 'react';
import Dispatcher from './processors/Dispatcher';



const themeObj = {
    fg: "yellow",
    bg: "white"
};

//var params = {dispatchers: dispatcherObj, theme: themeObj, state: stateObj};

const AppContext = createContext({});

export class AppContextProvider extends React.Component {   
    updateState = (state) => {
      this.setState({...state});
    };

    getDispatcher = () => {
        if(!this.dispatcher) {
            this.dispatcher = new Dispatcher(this);
            this.setState({...this.state, dispatcher: this.dispatcher});
        }
        return this.dispatcher;
    };

    state = {updateState: this.updateState, 
        getDispatcher: this.getDispatcher,
        theme: themeObj};
  
    render() {
      return (
          //passing the state object as a value prop to all children
          <AppContext.Provider value={this.state}>
              {this.props.children}
          </AppContext.Provider>
      )}
  }


export default AppContext;
