import React, { lazy, useEffect, useState, Suspense, useContext } from 'react';
import { get, isArray } from 'lodash';
import {
  useInterval,
  getMetricData,
  getMetricDataInterval,
  makeChartDataByTimeRange,
  makeHistogramSeriesData,
  INTERVAL,
  METRIC_TYPE,
} from '../../realtimeHelpers';
import { AppContext } from '../../context';

const ApexChartComponent = lazy(() => import('./ApexChart'));

const loading = () => <p>Loading chart...</p>;


const Metric = ({ metrics, unit, timeRange }) => {
  const [data, fetchMetricData] = useState([]);
  const appData = useContext(AppContext);
  const appStatus = get(appData, 'status', '');
  const timestamp = get(appData, 'timestamp', '');
  const isRealtime = !['finished', 'cancel'].includes(appStatus);

  useEffect(() => {
    if (metrics.length > 0) {
      getMetricData(metrics, timeRange, timestamp, isRealtime).then(res => {
        fetchMetricData(res);
      });
    }
  }, [metrics]);
  useInterval(() => {
    getMetricDataInterval(metrics, data)
      .then(res => {
        fetchMetricData(res);
      });
  }, isRealtime ? INTERVAL : null);
  const metricType = get(data, '[0].type', '');
  let series;
  if (metricType === METRIC_TYPE.HISTOGRAM) {
    series = [...makeHistogramSeriesData(get(data, '[0].chartData.data', []))];
  } else {
    if (isArray(data)) {
      series = data.map(d => get(d, 'chartData', {
        name: d.title,
        data: []
      }));
    }
  }
  const chartData = isRealtime ? makeChartDataByTimeRange(series, timeRange) : series;
  return <Suspense fallback={loading()}>
    <ApexChartComponent
      height="220"
      series={chartData}
      unit={unit}/>
  </Suspense>;
};

export default Metric;
