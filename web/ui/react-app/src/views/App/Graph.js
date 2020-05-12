import React, { useEffect, useState, lazy, Suspense, useContext } from 'react';
import { get } from 'lodash';
import GoBenchAPI from '../../api/gobench';
import { TIME_RANGE } from '../../realtimeHelpers';
import { AppContext } from '../../context';

const MetricComponent = lazy(() => import('./Metric'));

const loading = () => <p>Loading graph...</p>;

const Graph = ({ graph = {}, timestamp }) => {
  const [metrics, fetchMetrics] = useState([]);
  const [timeRange, setTimeRange] = useState(TIME_RANGE['1h']);
  const appData = useContext(AppContext);
  const appStatus = get(appData, 'status', '');
  useEffect(() => {
    if (graph && graph.id) {
      GoBenchAPI.getMetrics(graph.id).then(res => {
        return fetchMetrics(res);
      })
    }
  }, [graph]);
  const isRealtime = !['finished', 'cancel'].includes(appStatus);

  return (
    <div className="graph">
      <div className="graph-header">
        <h5 title={graph.id || ''}
            className="graph-title">{get(graph, 'title', '')} ({get(graph, 'unit', '')})</h5>
        {
          isRealtime &&
          <div className="options-group">
            <ul className="time-range-options-list">
              <li className="time-range-option">
                <button
                  className={timeRange === TIME_RANGE['5m'] ? 'active' : ''}
                  onClick={() => setTimeRange(TIME_RANGE['5m'])}>5m
                </button>
              </li>
              <li className="time-range-option">
                <button
                  className={timeRange === TIME_RANGE['15m'] ? 'active' : ''}
                  onClick={() => setTimeRange(TIME_RANGE['15m'])}>15m
                </button>
              </li>
              <li className="time-range-option">
                <button
                  className={timeRange === TIME_RANGE['30m'] ? 'active' : ''}
                  onClick={() => setTimeRange(TIME_RANGE['30m'])}>30m
                </button>
              </li>
              <li className="time-range-option">
                <button
                  className={timeRange === TIME_RANGE['1h'] ? 'active' : ''}
                  onClick={() => setTimeRange(TIME_RANGE['1h'])}>1h
                </button>
              </li>
              <li className="time-range-option">
                <button
                  className={timeRange === TIME_RANGE['12h'] ? 'active' : ''}
                  onClick={() => setTimeRange(TIME_RANGE['12h'])}>12h
                </button>
              </li>
              <li className="time-range-option">
                <button
                  className={timeRange === TIME_RANGE['24h'] ? 'active' : ''}
                  onClick={() => setTimeRange(TIME_RANGE['24h'])}>24h
                </button>
              </li>
            </ul>
          </div>
        }
      </div>
      <Suspense fallback={loading()}>
        <MetricComponent
          timeRange={timeRange}
          graph={graph}
          metrics={metrics}
          timestamp={timestamp}
          unit={get(graph, 'unit', '')}/>
      </Suspense>
    </div>
  );
};

export default Graph;
