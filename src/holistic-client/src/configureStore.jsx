import { applyMiddleware, compose, createStore } from 'redux'
import thunkMiddleware from 'redux-thunk'
import loggerMiddleware from 'redux-logger'
import monitorReducersEnhancer from './enhancers/monitorReducer'
import rootReducer from './reducers'

const middlewares = [loggerMiddleware, thunkMiddleware]
const middlewareEnhancer = applyMiddleware(...middlewares)

const enhancers = [middlewareEnhancer, monitorReducersEnhancer]
const composedEnhancers = compose(...enhancers)

const store = createStore(rootReducer, {}, composedEnhancers)

export default store
