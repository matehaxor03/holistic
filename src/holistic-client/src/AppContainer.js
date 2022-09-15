import React from 'react';
import App from './App';
import store from './configureStore'
import { Provider } from 'react-redux'

class AppContainer extends React.Component { 
  
  render() {
    return (
      <Provider store={store}> 
      
          <App></App>
    
      </Provider>
    );
  }

  
}

export default AppContainer;
