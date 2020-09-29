import apiClient from 'services/axios'
import { init, listApi, detailApi, createApi, updateApi, destroyApi } from '../index'
import { METRIC_TYPE } from 'constant'
import { mapLimit } from 'async'
import { get, maxBy, orderBy } from 'lodash'
import { getChartData, getDataByType } from 'utils/chart'
const API = {
  list: '/applications',
  detail: '/applications/{0}',
  create: '/applications',
  update: '/applications/{0}',
  delete: '/applications/{0}',
  log: '/applications/{0}/logs/{1}',
  // other apis
  cancel: '/applications/{0}/cancel',
  groups: '/applications/{0}/groups',
  graphs: '/groups/{0}/graphs',
  graphMetrics: '/graphs/{0}/metrics',
  counters: '/metrics/{0}/counters',
  histograms: '/metrics/{0}/histograms',
  gauges: '/metrics/{0}/gauges',
  metrics: '/metrics/{0}'

}
export const list = async (limit = 10, skip = 0, name, zone, dtms) => {
  init(API)
  return listApi(limit, skip, name, zone, dtms)
}
export const detail = async (id) => {
  init(API)
  return detailApi(id)
}
export const create = async (data) => {
  init(API)
  return createApi(data)
}
export const update = async (id, data) => {
  init(API)
  return updateApi(id, data)
}
export const destroy = async (id) => {
  init(API)
  return destroyApi(id)
}
export const cancel = async (id) => {
  const response = await apiClient.put(API.cancel.format(id))
  if (response) {
    return response.data
  }
}
export const getGroups = async (id) => {
  const response = await apiClient.get(API.groups.format(id))
  if (response) {
    return response.data
  }
}
export const getGraphs = async (id) => {
  const response = await apiClient.get(API.graphs.format(id))
  if (response) {
    return response.data
  }
}
export const getGraphMetrics = async (id) => {
  const response = await apiClient.get(API.graphMetrics.format(id))
  if (response) {
    return response.data
  }
}
export const getCounters = async (id, from = '', end = '') => {
  const response = await apiClient.get(`${API.counters.format(id)}${from}${end}`)
  if (response) {
    return response.data
  }
}
export const getHistograms = async (id, from = '', end = '') => {
  const response = await apiClient.get(`${API.histograms.format(id)}${from}${end}`)
  if (response) {
    return response.data
  }
}
export const getGauges = async (id, from = '', end = '') => {
  const response = await apiClient.get(`${API.gauges.format(id)}${from}${end}`)
  if (response) {
    return response.data
  }
}
export const getMetrics = async (id, from = '', end = '') => {
  const response = await apiClient.get(`${API.metrics.format(id)}${from}${end}`)
  if (response) {
    return response.data
  }
}
export const getMetricData = async (id = 0, type = METRIC_TYPE.COUNTER, fromTime = null, toTime = null) => {
  const from = fromTime ? `?from=${fromTime}` : ''
  const end = toTime ? `&end=${toTime}` : ''
  switch (type) {
    case METRIC_TYPE.COUNTER:
      return getCounters(id, from, end)
    case METRIC_TYPE.HISTOGRAM:
      return getHistograms(id, from, end)
    case METRIC_TYPE.GAUGE:
      return getGauges(id, from, end)
    default:
      return getMetrics(id, from, end)
  }
}

/***
 * get metrics data (first fetch)
 * @param metrics
 * @param timeRange
 * @param timestamp
 * @param isRealtime
 * @returns {Promise<unknown[]>}
 */

export const getMetricDataRealtime = async (metrics, timeRange = 3600, timestamp, isRealtime) => {
  const now = new Date().getTime()
  const fromTime = Math.round((now - timestamp) / 1000) < timeRange ? timestamp : (now - (timeRange * 1000))
  return await mapLimit(metrics, 5, async (m) => {
    let response
    if (isRealtime) {
      response = await getMetricData(m.id, m.type, fromTime, now)
    } else {
      response = await getMetricData(m.id, m.type)
    }
    if (response.length === 0) {
      return {
        ...m,
        lastTimestamp: timestamp,
        chartData: {
          name: m.title,
          data: []
        }
      }
    }
    const lastTimestamp = get(maxBy(response, m => m.time), 'time')
    return {
      ...m,
      lastTimestamp,
      chartData: {
        name: m.title,
        data: m.type === METRIC_TYPE.HISTOGRAM ? response : getChartData(m.type, response)
      }
    }
  })
}
/***
 * get metrics data interval
 * @param metrics
 * @param oldData
 * @returns {Promise<unknown[]>}
 */
export const getMetricDataPolling = async (metrics, oldData = []) => {
  return await mapLimit(metrics, 5, async (mtr) => {
    const oldMetricData = oldData.find(o => mtr.id === get(o, ['id'], ''))
    const timestamp = get(oldMetricData, 'lastTimestamp', '')
    if (!timestamp) {
      return oldMetricData
    }
    const response = await getMetricData(mtr.id, mtr.type, timestamp)
    if (!response || response.length === 0) {
      return oldMetricData
    }
    const dataByType = getDataByType(response, mtr.type)
    const oldMetricChartData = get(oldMetricData, ['chartData', 'data'], [])
    const newData = [...oldMetricChartData, ...dataByType]
    return {
      ...oldMetricData,
      lastTimestamp: get(orderBy(response, ['time'], 'desc'), '[0].time'),
      chartData: {
        name: mtr.title,
        data: newData
      }
    }
  })
}

export const logs = async (id, type) => {
  const response = await apiClient.get(API.log.format(id, type))
  if (response) {
    return response.data
  }
}
