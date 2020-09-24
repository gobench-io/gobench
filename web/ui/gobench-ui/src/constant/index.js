export const values = {
  counter: 'count',
  gauge: 'value'
}
export const METRIC_TYPE = {
  HISTOGRAM: 'histogram',
  COUNTER: 'counter',
  GAUGE: 'gauge'
}

export const INTERVAL = 10000 // realtime data inteval in miliseconds
export const DEFAULT_VALUE = null // default value for empty data

export const TIME_RANGE = {
  '5m': 5 * 60,
  '15m': 15 * 60,
  '30m': 30 * 60,
  '1h': 60 * 60,
  '12h': 12 * 60 * 60,
  '24h': 24 * 60 * 60
}
