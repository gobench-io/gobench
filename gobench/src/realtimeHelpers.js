import { useEffect, useRef } from 'react';
import { get, isArray, maxBy, orderBy } from 'lodash';
import GoBenchAPI from './api/gobench';

const values = {
  counter: 'count',
  gauge: 'value'
};
export const METRIC_TYPE = {
  HISTOGRAM: 'histogram',
  COUNTER: 'counter',
  GAUGE: 'gauge'
};

export const INTERVAL = 10000; // realtime data inteval in miliseconds
export const DEFAULT_VALUE = null; // default value for empty data 

export const TIME_RANGE = {
  '5m': 5 * 60,
  '15m': 15 * 60,
  '30m': 30 * 60,
  '1h': 60 * 60,
  '12h': 12 * 60 * 60,
  '24h': 24 * 60 * 60
};

/***
 * useInterval
 * hook API written by Dan Abramov
 * @param callback
 * @param delay
 */
export const useInterval = (callback, delay) => {
  const savedCallback = useRef();
  useEffect(() => {
    savedCallback.current = callback;
  }, [callback]);

  useEffect(() => {
    function tick() {
      savedCallback.current();
    }

    if (delay !== null) {
      let id = setInterval(tick, delay);
      return () => clearInterval(id);
    }
  }, [delay]);
};

/***
 * get value of metric by type
 * @param metric
 * @param valueType
 * @returns {*|number}
 */
const getValue = (metric, valueType) => (metric[valueType]).toFixed(0) || DEFAULT_VALUE;

/***
 * Get chart data depend on metric type
 * @param type
 * @param data
 * @returns {*}
 */
const getChartData = (type, data) => isArray(data) ? data.map(m => ({
  x: m.time,
  y: getValue(m, values[type])
})) : [{
  x: new Date(data.time).getTime(),
  y: getValue(data, values[type])
}];


/***
 * Make chartData by metric type
 * @param data
 * @returns {({data: ([]|*[]), name: string}|{data: ([]|*[]), name: string}|{data: ([]|*[]), name: string}|{data: ([]|*[]), name: string}|{data: ([]|*[]), name: string})[]|{data: *, name: *}}
 */
export const makeHistogramSeriesData = (data) => {
  let hData = { min: [], max: [], p75: [], p95: [], p99: [] };
  data.forEach(d => {
    hData = {
      min: [...hData.min, { x: d.time, y: (d.min).toFixed(0) || 0 }],
      max: [...hData.max, { x: d.time, y: (d.max).toFixed(0) || 0 }],
      p75: [...hData.p75, { x: d.time, y: (d.p75).toFixed(0) || 0 }],
      p95: [...hData.p95, { x: d.time, y: (d.p95).toFixed(0) || 0 }],
      p99: [...hData.p99, { x: d.time, y: (d.p99).toFixed(0) || 0 }],
    }
  });

  return [
    { name: 'min', data: hData.min },
    { name: 'max', data: hData.max },
    { name: 'p75', data: hData.p75 },
    { name: 'p95', data: hData.p95 },
    { name: 'p99', data: hData.p99 }
  ]
};

/***
 * get metrics data (first fetch)
 * @param metrics
 * @param timeRange
 * @param timestamp
 * @param isRealtime
 * @returns {Promise<unknown[]>}
 */
export const getMetricData = async (metrics, timeRange = 3600, timestamp, isRealtime) => {
  const now = new Date().getTime();
  const fromTime = Math.round((now - timestamp) / 1000) < timeRange ? timestamp : (now - (timeRange * 1000));
  const metricsData = metrics.map(async m => {
    let mData;
    if (isRealtime) {
      mData = await GoBenchAPI.getMetricData(m.id, m.type, fromTime, now).then(rs => rs);
    } else {
      mData = await GoBenchAPI.getMetricData(m.id, m.type).then(rs => rs);
    }
    if (mData.length === 0) {
      return {
        ...m,
        lastTimestamp: timestamp,
        chartData: {
          name: m.title,
          data: []
        }
      };
    }
    const lastTimestamp = get(maxBy(mData, m => m.time), 'time');
    return {
      ...m,
      lastTimestamp,
      chartData: {
        name: m.title,
        data: m.type === METRIC_TYPE.HISTOGRAM ? mData : getChartData(m.type, mData)
      }
    }
  });
  return await Promise.all(metricsData)
    .then(rs => rs)
    .catch(err => err);
};

const getDataByType = (data, type) => {
  return (type === METRIC_TYPE.HISTOGRAM ? data : getChartData(type, data));
};

/***
 * get metrics data interval
 * @param metrics
 * @param oldData
 * @returns {Promise<unknown[]>}
 */
export const getMetricDataInterval = async (metrics, oldData = null) => {
  return await Promise.all(metrics.map(mtr => {
    const oldMetricData = oldData.find(o => mtr.id === get(o, ['id'], ''));
    const timestamp = get(oldMetricData, 'lastTimestamp', '');
    if (timestamp) {
      return GoBenchAPI.getMetricData(mtr.id, mtr.type, timestamp)
        .then(mData => {
          if (mData.length > 0) {
            const dataByType = getDataByType(mData, mtr.type);
            const oldMetricChartData = get(oldMetricData, ['chartData', 'data'], []);
            const newData = [...oldMetricChartData, ...dataByType];
            return {
              ...oldMetricData,
              lastTimestamp: get(orderBy(mData, ['time'], 'desc'), '[0].time'),
              chartData: {
                name: mtr.title,
                data: newData
              }
            }
          }
          return oldMetricData;
        });
    }
    return oldMetricData;
  }))
    .then(rs => rs)
    .catch(err => err);
};

/**
 * make timestamp without to second
 * @param timestamp
 * @returns {number}
 */
const fixSecond = (timestamp) => Math.round(timestamp / 1000) * 1000;

/***
 * Make chart data by time range
 * @param rawData
 * @param timeRange
 * @returns {*}
 */
export const makeChartDataByTimeRange = (rawData = [], timeRange = 3600) => {
  const timeRangeMiliseconds = timeRange * 1000;
  return rawData.map((seri) => {
    const seriData = get(seri, 'data', []);
    if (seriData.length === 0) {
      return [];
    }
    const seriDataLength = seriData.length || 0;
    const firstData = seriData[0];
    const lastData = seriData[seriDataLength - 1];
    const lastDataTime = get(lastData, 'x', 0);
    const firstDataTime = get(firstData, 'x', 0);
    const dataTime = fixSecond(lastDataTime - firstDataTime);

    if (dataTime < timeRangeMiliseconds) {
      const extraTime = timeRangeMiliseconds - dataTime;
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
  });
}; 