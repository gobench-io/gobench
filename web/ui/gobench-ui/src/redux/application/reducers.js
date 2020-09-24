import actions from './actions'

const initialState = {
  list: [],
  detail: {},
  total: 0,
  groups: [],
  graphs: [],
  graphMetrics: [],
  counters: [],
  histograms: [],
  gauges: [],
  metrics: [],
  metricDatas: [],
  metricDataRealtime: [],
  loading: false
}

export default (state = initialState, action) => {
  switch (action.type) {
    case actions.SET_STATE:
      return { ...state, ...action.payload }
    case actions.SET_GRAPH_STATE: {
      if (action.payload.graphs.some(x => state.graphs.some(a => a.id === x.id))) {
        const oldGraph = state.graphs.filter(x => action.payload.graphs.every(a => a.id !== x.id))
        return { ...state, graphs: [...oldGraph, ...action.payload.graphs] }
      }
      return { ...state, graphs: [...state.graphs, ...action.payload.graphs] }
    }
    case actions.SET_GRAPH_METRIC_STATE: {
      if (state.graphMetrics.some(x => x.graphId === action.payload.graphId)) {
        return {
          ...state,
          graphMetrics: state.graphMetrics.map(x => {
            if (x.graphId === action.payload.graphId) {
              return action.payload
            }
            return x
          })
        }
      }
      return { ...state, graphMetrics: [...state.graphMetrics, action.payload] }
    }
    case actions.SET_GRAPH_METRIC_DATA: {
      if (state.metricDatas.some(x => x.graphId === action.payload.graphId)) {
        return {
          ...state,
          metricDatas: state.metricDatas.map(x => {
            if (x.graphId === action.payload.graphId) {
              return action.payload
            }
            return x
          })
        }
      }
      return { ...state, metricDatas: [...state.metricDatas, action.payload] }
    }
    case actions.SET_METRIC_STATE:

      if (state.graphs.some(x => x.id === action.payload.id)) {
        return {
          ...state,
          graphs: state.graphs.map(x => {
            if (x.id === action.payload.id) {
              return { ...x, metrics: action.payload.metrics }
            }
            return x
          })
        }
      }
      return { ...state, graphMetrics: [...state.graphMetrics, action.payload] }
    case actions.SET_METRIC_DATA_REALTIME: {
      // return { ...state, metricDataRealtime: [...state.metricDataRealtime, ...action.payload] }
      let metricDataRealtime
      if (state.metricDataRealtime.some(x => action.payload.some(a => a.id === x.id))) {
        metricDataRealtime = state.metricDataRealtime.map(x => {
          if (x.id === action.payload.id) {
            return action.payload
          }
          return x
        })
      } else {
        metricDataRealtime = [...state.metricDataRealtime, ...action.payload]
      }

      return { ...state, metricDataRealtime }
    }
    default:
      return state
  }
}
