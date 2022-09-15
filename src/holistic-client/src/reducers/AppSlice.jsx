import { createSlice } from '@reduxjs/toolkit'

export const AppSlice = createSlice({
  name: 'app',
  initialState: {
    currentPage: null,
    test: "test"
  },
  reducers: {
    viewPage: (state, action) => {
        var payload = action.payload;
        state.currentPage = payload;
        return state;
    }
  } 
})

export const { viewPage } = AppSlice.actions

export const SelectCurrentPage = (reducers) => reducers.app.currentPage

export default AppSlice.reducer