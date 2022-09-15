import React from 'react';
import Dispatcher from './processors/Dispatcher';

var dispatcherObj = new Dispatcher();

const themeObj = {
    fg: "yellow",
    bg: "white"
};

var stateObj = {};
var params = {controllers: dispatcherObj, theme: themeObj, state: stateObj};
var AppContext = React.createContext(params);

export default AppContext;
