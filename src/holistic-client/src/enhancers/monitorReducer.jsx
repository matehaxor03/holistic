const floor = number => Math.floor(number * 1000)

const monitorReducerEnhancer =
  createStore => (reducer, initialState, enhancer) => {
    const monitoredReducer = (state, action) => {
      const start = performance.now()
      const newState = reducer(state, action)
      const end = performance.now()
      const diff = floor(end - start)

      console.log('reducer process time: ', diff, ' µs')

      return newState
    }

    return createStore(monitoredReducer, initialState, enhancer)
  }

export default monitorReducerEnhancer