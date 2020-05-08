import API from './api';
import { get } from 'lodash';
import { METRIC_TYPE } from '../realtimeHelpers';

const GoBenchAPI = {
  getAppInfo() {
    return new Promise((resolve, reject) => {
      API.axios().get(`api/application`)
        .then(res => resolve(get(res, 'data', null)))
        .catch(reject);
    });
  },
  getGroups() {
    return new Promise((resolve, reject) => {
      API.axios().get(`api/groups`)
        .then(res => resolve(get(res, 'data', [])))
        .catch(reject);
    });
  },
  getGraphs(id) {
    return new Promise((resolve, reject) => {
      API.axios().get(`api/groups/${id}/graphs`)
        .then(res => resolve(get(res, 'data', [])))
        .catch(reject);
    });
  },
  getMetrics(id) {
    return new Promise((resolve, reject) => {
      API.axios().get(`api/graphs/${id}/metrics`)
        .then(res => resolve(get(res, 'data', [])))
        .catch(reject);
    });
  },
  getMetricData(id = 0, type = 'counter', fromTime = null, toTime = null) {
    const from = fromTime ? `?from=${fromTime}` : '';
    const end = toTime ? `&end=${toTime}` : '';
    switch (type) {
      case METRIC_TYPE.COUNTER:
        return new Promise((resolve, reject) => {
          API.axios().get(`api/metrics/${id}/counters${from}${end}`)
            .then(res => resolve(get(res, 'data', [])))
            .catch(reject);
        });

      case METRIC_TYPE.HISTOGRAM:
        return new Promise((resolve, reject) => {
          API.axios().get(`api/metrics/${id}/histograms${from}${end}`)
            .then(res => resolve(get(res, 'data', [])))
            .catch(reject);
        });

      case METRIC_TYPE.GAUGE:
        return new Promise((resolve, reject) => {
          API.axios().get(`api/metrics/${id}/gauges${from}${end}`)
            .then(res => resolve(get(res, 'data', [])))
            .catch(reject);
        });
      default:
        return new Promise((resolve, reject) => {
          API.axios().get(`api/metrics/${id}`)
            .then(res => resolve(get(res, 'data', [])))
            .catch(reject);
        });
    }
  }
};

export default GoBenchAPI;
