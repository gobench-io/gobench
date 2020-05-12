import axios from 'axios';
import { get } from 'lodash';

let axiosInstance = {};

const API = {
  init({
         baseURL, on401, on404, onNoResponse
       }) {
    axiosInstance = axios.create({
      baseURL,
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json'
      }
    });

    this.on401 = on401 || function noop() {
    };
    this.on404 = on404 || function noop() {
    };
    this.onNoResponse = onNoResponse || function noop() {
    };
    axiosInstance.interceptors.response.use(undefined, (error) => {
      if (!error.response) {
        return this.onNoResponse(error);
      }
      if (error.response.status === 401) {
        this.on401(error.request, error.response);
      }
      if (error.response.status === 404) {
        this.on404(error.request, error.response);
      }
      return Promise.reject(error);
    });
  },
  axios() {
    return axiosInstance;
  }
};

export const APIError = err => get(err, ['response', 'data', 'error'], {});

export default API;
