import { combineReducers } from 'redux'
import AppReducer from './AppSlice'

const rootReducer = combineReducers({
    app: AppReducer
  })
  
  export default rootReducer