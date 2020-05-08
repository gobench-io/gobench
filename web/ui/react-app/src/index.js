import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './views/App/App';
import API from './api/api';
import config from './config/config';
// import * as serviceWorker from './serviceWorker';

API.init({
  baseURL: config.apiEndpoint,
  on404: () => {
    console.log('API not found');
  }
});

ReactDOM.render(
  <React.StrictMode>
    <App/>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
// serviceWorker.unregister();
