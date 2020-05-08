import * as configDefault from './config.default';

console.log('process.env.NODE_ENV :', process.env.NODE_ENV);
const config = {
  env: process.env.NODE_ENV,
  apiEndpoint: process.env.NODE_ENV === "production" ? '/' : configDefault.apiEndpoint,
};
export default config;
