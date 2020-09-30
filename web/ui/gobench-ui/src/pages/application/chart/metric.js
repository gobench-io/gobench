import React, { lazy, useEffect, Suspense, useRef } from 'react'
import { connect } from 'react-redux'
import { withRouter } from 'react-router-dom'
import { get, isArray } from 'lodash'
import { INTERVAL, METRIC_TYPE } from 'constant'
import { makeHistogramSeriesData, makeChartDataByTimeRange } from 'utils/chart'

const ApexChart = lazy(() => import('./apex-chart'))

const mapStateToProps = ({ application, dispatch }) => {
  const { detail, graphMetrics, metricDatas } = application
  return {
    detail,
    graphMetrics,
    metricDatas,
    dispatch
  }
}

const loading = () => <p>Loading graph...</p>

const DefaultPage = ({ detail, graph, graphMetrics, metricDatas, unit, timeRange, dispatch }) => {
  let series
  const appStatus = get(detail, 'status', '')
  const timestamp = get(detail, 'timestamp', '')
  const isRealtime = appStatus === 'running'
  const { graphId, metrics } = graphMetrics.find(x => x.graphId === graph.id) || { metrics: [] }
  const metricData = metricDatas.find(x => x.graphId === graph.id) || { metrics: [] }
  const metricType = get(metrics, '[0].type', '')
  useEffect(() => {
    if (metrics.length > 0) {
      if (metricData.metrics.length > 0) {
      // no need get data again
        return
      }
      dispatch({
        type: 'application/GRAPH_METRIC_DATA',
        payload: { id: graphId, metrics, timeRange, timestamp, isRealtime }
      })
    }
  }, [graphId, timestamp, metricData.metrics.length, dispatch])
  useInterval(() => {
    dispatch({
      type: 'application/METRIC_DATA_POLLING',
      payload: { id: graph.id, metrics, data: metricData.metrics }
    })
  }, isRealtime ? INTERVAL : null)

  if (metricType === METRIC_TYPE.HISTOGRAM) {
    series = [...makeHistogramSeriesData(get(metricData.metrics, '[0].chartData.data', []))]
  } else {
    if (isArray(metricData.metrics)) {
      series = metricData.metrics.map(d => get(d, 'chartData', {
        name: d.title,
        data: []
      }))
    }
  }
  const chartData = isRealtime ? makeChartDataByTimeRange(series, timeRange) : series
  return (
    <>
      <Suspense fallback={loading()}>
        <ApexChart
          series={chartData}
          unit={unit}
        />
      </Suspense>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))

export const useInterval = (callback, delay) => {
  const savedCallback = useRef()
  useEffect(() => {
    savedCallback.current = callback
  }, [callback])

  useEffect(() => {
    function tick () {
      savedCallback.current()
    }

    if (delay !== null) {
      const id = setInterval(tick, delay)
      return () => clearInterval(id)
    }
  }, [delay])
}
