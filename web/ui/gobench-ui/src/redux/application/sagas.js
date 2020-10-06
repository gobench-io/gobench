import { all, put, call, takeEvery, select } from 'redux-saga/effects'
import { notification } from 'antd'
import actions from './actions'
import { history } from 'index'
import {
  list, detail, create, update, destroy, cancel,
  getGroups, getGraphs, getGraphMetrics, getCounters,
  getHistograms, getGauges, getMetrics, getMetricData,
  getOfflineMetricData, getMetricDataPolling,
  logs, addTag, removeTag, getTags
} from 'services/application'

export function * LIST ({ payload }) {
  const { skip, limit } = payload
  yield loading(true)
  const response = yield call(list, skip, limit)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        list: response,
        total: (response || []).length
      }
    })
  }
  yield loading(false)
}
export function * DETAIL ({ payload }) {
  const { id } = payload
  yield loading(true)
  const response = yield call(detail, id)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        detail: response
      }
    })
  }
  yield loading(false)
}
export function * CREATE ({ payload }) {
  const { name, scenario, gomod, gosum } = payload
  yield loading(true)
  const response = yield call(create, {
    name,
    scenario: window.btoa(unescape(encodeURIComponent(scenario))),
    gomod: window.btoa(unescape(encodeURIComponent(gomod))),
    gosum: window.btoa(unescape(encodeURIComponent(gosum)))
  })
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        detail: response
      }
    })
    notification.success({
      message: 'Application created',
      description: 'You have successfully created an application!'
    })
    history.push(`/applications/${response.id}`)
  }
  // clear clone data
  yield put({
    type: 'application/SET_STATE',
    payload: {
      clone: undefined
    }
  })

  yield loading(false)
}
export function * CLONE ({ payload }) {
  const { data } = payload
  yield put({
    type: 'application/SET_STATE',
    payload: {
      clone: data
    }
  })
  history.push('/applications/create')
}
export function * UPDATE ({ payload }) {
  const { id, data } = payload
  yield loading(true)
  const response = yield call(update, id, data)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        detail: response
      }
    })
  }
  notification.success({
    message: 'Application updated',
    description: 'You have successfully updated an application!'
  })
  yield loading(false)
}
export function * CANCEL ({ payload }) {
  const { id, data } = payload
  const state = yield select()
  const { list } = state.application
  yield loading(true)
  const response = yield call(cancel, id, data)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        list: list.map(x => {
          if (x.id === response.id) {
            return response
          }
          return x
        }),
        detail: response
      }
    })
    notification.success({
      message: 'Application canceled',
      description: 'You have successfully canceled an application!'
    })
  }
  yield loading(false)
}
export function * DELETE ({ payload }) {
  const { id, redirect } = payload
  const state = yield select()
  const { list, total } = state.application
  yield loading(true)
  const response = yield call(destroy, id)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        list: list.filter(x => x.id !== id),
        total: total - 1
      }
    })
    if (redirect) {
      history.push(redirect)
    }
  }
  yield loading(false)
}
export function * GROUPS ({ payload }) {
  const { id } = payload

  yield loading(true)
  const response = yield call(getGroups, id)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        groups: response
      }
    })
  }
  yield loading(false)
}

export function * COUNTERS ({ payload }) {
  const { id, from, end } = payload
  yield loading(true)
  const response = yield call(getCounters, id, from, end)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        counters: response
      }
    })
  }
  yield loading(false)
}
export function * HISTOGRAMS ({ payload }) {
  const { id, from, end } = payload
  yield loading(true)
  const response = yield call(getHistograms, id, from, end)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        histograms: response
      }
    })
  }
  yield loading(false)
}
export function * GAUGES ({ payload }) {
  const { id, from, end } = payload
  yield loading(true)
  const response = yield call(getGauges, id, from, end)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        gauges: response
      }
    })
  }
  yield loading(false)
}
export function * METRICS ({ payload }) {
  const { id, from, end } = payload
  yield loading(true)
  const response = yield call(getMetrics, id, from, end)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        metrics: response
      }
    })
  }
  yield loading(false)
}
export function * METRIC_DATA ({ payload }) {
  const { id, type, fromTime, toTime } = payload
  yield loading(true)
  const response = yield call(getMetricData, id, type, fromTime, toTime)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        metricDatas: response
      }
    })
  }
  yield loading(false)
}
export function * GRAPHS ({ payload }) {
  const { id } = payload
  yield loading(true)
  const response = yield call(getGraphs, id)
  if (response) {
    const graphs = yield response.map(x => {
      x.groupId = id
      return x
    })
    yield put({
      type: 'application/SET_GRAPH_STATE',
      payload: {
        graphs
      }
    })
  }
  yield loading(false)
}
export function * GRAPH_METRICS ({ payload }) {
  const { id } = payload

  yield loading(true)
  const response = yield call(getGraphMetrics, id)
  if (response) {
    yield put({
      type: 'application/SET_GRAPH_METRIC_STATE',
      payload: {
        graphId: id,
        metrics: response
      }
    })
  }
  yield loading(false)
}
export function * GRAPH_METRIC_DATA ({ payload }) {
  const { id, metrics, timeRange, timestamp, isRealtime } = payload
  yield loading(true)
  const response = yield call(getOfflineMetricData, metrics, timeRange, timestamp, isRealtime)
  if (response) {
    yield put({
      type: 'application/SET_GRAPH_METRIC_DATA',
      payload: {
        graphId: id,
        metrics: response
      }
    })
  }
  yield loading(false)
}
export function * METRIC_DATA_POLLING ({ payload }) {
  const { id, metrics, data } = payload
  yield loading(true)
  const response = yield call(getMetricDataPolling, metrics, data)
  if (response) {
    yield put({
      type: 'application/SET_GRAPH_METRIC_DATA',
      payload: {
        graphId: id,
        metrics: response
      }
    })
  }
  yield loading(false)
}
export function * SYSLOG ({ payload }) {
  const { id } = payload
  yield loading(true)
  const response = yield call(logs, id, 'system')
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        logs: response
      }
    })
  }
  yield loading(false)
}
export function * LOG ({ payload }) {
  const { id } = payload
  yield loading(true)
  const response = yield call(logs, id, 'user')
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        logs: response
      }
    })
  }
  yield loading(false)
}
export function * TAGS ({ payload }) {
  const { id } = payload
  yield loading(true)
  const response = yield call(getTags, id)
  if (response) {
    yield put({
      type: 'application/SET_STATE',
      payload: {
        tags: response
      }
    })
  }
  yield loading(false)
}
export function * TAG_ADD ({ payload }) {
  const { id, name } = payload
  yield loading(true)
  const response = yield call(addTag, id, name)
  if (response) {
    yield put({
      type: 'application/SET_TAG_STATE',
      payload: [
        response
      ]
    })
  }
  yield loading(false)
}
export function * TAG_REMOVE ({ payload }) {
  const { id, tagId } = payload
  yield loading(true)
  const response = yield call(removeTag, id, tagId)
  if (response) {
    yield put({
      type: 'application/SET_TAG_STATE',
      payload: [
        response
      ]
    })
  }
  notification.success({
    message: 'Application Tag deleted',
    description: 'You have successfully delete a tag for this application!'
  })
  yield loading(false)
}
function * loading (isLoading = false) {
  yield put({
    type: 'application/SET_STATE',
    payload: {
      loading: isLoading
    }
  })
}
export default function * rootSaga () {
  yield all([
    takeEvery(actions.LIST, LIST),
    takeEvery(actions.DETAIL, DETAIL),
    takeEvery(actions.CREATE, CREATE),
    takeEvery(actions.UPDATE, UPDATE),
    takeEvery(actions.DELETE, DELETE),
    takeEvery(actions.LOG, LOG),
    takeEvery(actions.SYSLOG, SYSLOG),
    takeEvery(actions.TAGS, TAGS),
    takeEvery(actions.TAG_ADD, TAG_ADD),
    takeEvery(actions.TAG_REMOVE, TAG_REMOVE),

    takeEvery(actions.CLONE, CLONE),
    takeEvery(actions.CANCEL, CANCEL),
    takeEvery(actions.GROUPS, GROUPS),
    takeEvery(actions.GRAPHS, GRAPHS),
    takeEvery(actions.GRAPH_METRICS, GRAPH_METRICS),
    // takeEvery(actions.COUNTERS, COUNTERS),
    takeEvery(actions.HISTOGRAMS, HISTOGRAMS),
    takeEvery(actions.GAUGES, GAUGES),
    takeEvery(actions.METRICS, METRICS),
    takeEvery(actions.METRIC_DATA, METRIC_DATA),
    takeEvery(actions.GRAPH_METRIC_DATA, GRAPH_METRIC_DATA),
    takeEvery(actions.METRIC_DATA_POLLING, METRIC_DATA_POLLING)
  ])
}
