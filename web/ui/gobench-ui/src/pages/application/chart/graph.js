import React, { lazy, useEffect, useState, Suspense } from 'react'
import { connect } from 'react-redux'
import { withRouter } from 'react-router-dom'
import { get } from 'lodash'
import { TIME_RANGE } from 'constant'

const Metric = lazy(() => import('./metric'))

const mapStateToProps = ({ application, dispatch }) => {
  const { detail, graphMetrics } = application
  return {
    detail,
    graphMetrics,
    dispatch
  }
}

const loading = () => <p>Loading graph...</p>
const DefaultPage = ({ detail, graph, graphMetrics, timestamp, dispatch }) => {
  const [timeRange, setTimeRange] = useState(TIME_RANGE['1h'])
  const appStatus = get(detail, 'status', '')
  const isRealtime = !['finished', 'cancel'].includes(appStatus)
  useEffect(() => {
    if (graph) {
      if (graphMetrics.some(x => x.graphId === graph.id)) {
        return
      }
      dispatch({
        type: 'application/GRAPH_METRICS',
        payload: { id: graph.id }
      })
    }
  }, [graph.id])
  return (
    <>
      <div className='graph'>
        <div className='graph-header'>
          <h5
            title={graph.id || ''}
            className='graph-title'
          >{get(graph, 'title', '')} ({get(graph, 'unit', '')})
          </h5>
          {
            isRealtime &&
              <div className='options-group'>
                <ul className='time-range-options-list'>
                  <li className='time-range-option'>
                    <button
                      className={timeRange === TIME_RANGE['5m'] ? 'active' : ''}
                      onClick={() => setTimeRange(TIME_RANGE['5m'])}
                    >5m
                    </button>
                  </li>
                  <li className='time-range-option'>
                    <button
                      className={timeRange === TIME_RANGE['15m'] ? 'active' : ''}
                      onClick={() => setTimeRange(TIME_RANGE['15m'])}
                    >15m
                    </button>
                  </li>
                  <li className='time-range-option'>
                    <button
                      className={timeRange === TIME_RANGE['30m'] ? 'active' : ''}
                      onClick={() => setTimeRange(TIME_RANGE['30m'])}
                    >30m
                    </button>
                  </li>
                  <li className='time-range-option'>
                    <button
                      className={timeRange === TIME_RANGE['1h'] ? 'active' : ''}
                      onClick={() => setTimeRange(TIME_RANGE['1h'])}
                    >1h
                    </button>
                  </li>
                  <li className='time-range-option'>
                    <button
                      className={timeRange === TIME_RANGE['12h'] ? 'active' : ''}
                      onClick={() => setTimeRange(TIME_RANGE['12h'])}
                    >12h
                    </button>
                  </li>
                  <li className='time-range-option'>
                    <button
                      className={timeRange === TIME_RANGE['24h'] ? 'active' : ''}
                      onClick={() => setTimeRange(TIME_RANGE['24h'])}
                    >24h
                    </button>
                  </li>
                </ul>
              </div>
          }
        </div>
        <Suspense fallback={loading()}>
          <Metric
            timeRange={timeRange}
            graph={graph}
            timestamp={timestamp}
            unit={get(graph, 'unit', '')}
          />
        </Suspense>
      </div>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))
