// define some action
const actions = {
  SET_STATE: 'application/SET_STATE',
  SET_GRAPH_STATE: 'application/SET_GRAPH_STATE',
  SET_GRAPH_METRIC_STATE: 'application/SET_GRAPH_METRIC_STATE',
  SET_METRIC_STATE: 'application/SET_METRIC_STATE',
  SET_METRIC_DATA_REALTIME: 'application/SET_METRIC_DATA_REALTIME',
  SET_GRAPH_METRIC_DATA: 'application/SET_GRAPH_METRIC_DATA',
  LIST: 'application/LIST',
  DETAIL: 'application/DETAIL',
  CREATE: 'application/CREATE',
  UPDATE: 'application/UPDATE',
  DELETE: 'application/DELETE',
  // other
  CANCEL: 'application/CANCEL',
  CLONE: 'application/CLONE',
  GROUPS: 'application/GROUPS',
  GRAPHS: 'application/GRAPHS',
  GRAPH_METRICS: 'application/GRAPH_METRICS',
  COUNTER: 'application/COUNTER',
  HISTOGRAMS: 'application/HISTOGRAMS',
  GAUGES: 'application/GAUGES',
  METRICS: 'application/METRICS',
  // complex
  METRIC_INTERVAL: 'application/METRIC_INTERVAL',
  METRIC_DATA: 'application/METRIC_DATA',
  GRAPH_METRIC_DATA: 'application/GRAPH_METRIC_DATA',
  METRIC_DATA_POLLING: 'application/METRIC_DATA_POLLING'

}

export default actions
