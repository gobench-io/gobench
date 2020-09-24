import { get, isArray } from 'lodash'
import { values, METRIC_TYPE, DEFAULT_VALUE } from 'constant'
/***
 * get value of metric by type
 * @param metric
 * @param valueType
 * @returns {*|number}
 */
export const getValue = (metric, valueType) => (metric[valueType]).toFixed(0) || DEFAULT_VALUE

/***
 * Get chart data depend on metric type
 * @param type
 * @param data
 * @returns {*}
 */
export const getChartData = (type, data) => isArray(data) ? data.map(m => ({
  x: m.time,
  y: getValue(m, values[type])
})) : [{
  x: new Date(data.time).getTime(),
  y: getValue(data, values[type])
}]

/***
 * Make chartData by metric type
 * @param data
 * @returns {({data: ([]|*[]), name: string}|{data: ([]|*[]), name: string}|{data: ([]|*[]), name: string}|{data: ([]|*[]), name: string}|{data: ([]|*[]), name: string})[]|{data: *, name: *}}
 */
export const hDataKeys = ['min', 'mean', 'p99', 'max']
export const makeHistogramSeriesData = (data) => {
  console.log('makeHis', data)
  let hData = hDataKeys.reduce((acc, key) => ({
    ...acc,
    [key]: []
  }), {})
  data.forEach(d => {
    hData = hDataKeys.reduce((acc, key) => ({
      ...acc,
      [key]: [...hData[key], { x: d.time, y: (d[key]).toFixed(0) || 0 }]
    }), {})
  })
  return hDataKeys.map(key => ({ name: key, data: hData[key] }))
}

export const getDataByType = (data, type) => {
  return (type === METRIC_TYPE.HISTOGRAM ? data : getChartData(type, data))
}
/**
 * make timestamp without to second
 * @param timestamp
 * @returns {number}
 */
const fixSecond = (timestamp) => Math.round(timestamp / 1000) * 1000

/***
 * Make chart data by time range
 * @param rawData
 * @param timeRange
 * @returns {*}
 */
export const makeChartDataByTimeRange = (rawData = [], timeRange = 3600) => {
  const timeRangeMiliseconds = timeRange * 1000
  return rawData.map((seri) => {
    const seriData = get(seri, 'data', [])
    if (seriData.length === 0) {
      return []
    }
    const seriDataLength = seriData.length || 0
    const firstData = seriData[0]
    const lastData = seriData[seriDataLength - 1]
    const lastDataTime = get(lastData, 'x', 0)
    const firstDataTime = get(firstData, 'x', 0)
    const dataTime = fixSecond(lastDataTime - firstDataTime)

    if (dataTime < timeRangeMiliseconds) {
      const extraTime = timeRangeMiliseconds - dataTime
      return {
        ...seri,
        data: [...seriData.map(s => ({
          x: Math.round(s.x / 1000) * 1000, y: s.y
        })), { x: lastDataTime + extraTime, y: DEFAULT_VALUE }]
      }
    }
    return {
      ...seri,
      data: seriData.filter(d => d.x >= (fixSecond(lastDataTime) - timeRangeMiliseconds))
    }
  })
}
