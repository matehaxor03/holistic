import React from 'react';
import { Provider, useSelector } from 'react-redux'
import { SelectCurrentPage } from './reducers/AppSlice'
import TablePage from './components/TablePage';
import NavBar from "./components/NavBar";
import store from './configureStore'

export default function App() { 
  const pages = {"TablePage": TablePage }
  const state = useSelector(SelectCurrentPage);
  var CurrentPage = null;

  console.log(state);

  if(state) {
    var Zlass = pages[state.type];
    CurrentPage = <Zlass id={state.type} params={state}></Zlass>;
  }

  return (
    <Provider store={store}>
        <div className="App">
          <NavBar id="menu"></NavBar>
          { CurrentPage }
        </div>
    </Provider>
  );
}