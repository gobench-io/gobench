import React, { lazy, Suspense } from 'react';
import ReactDOM from 'react-dom';
import './css/index.css';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";

import API from './api/api';
import config from './config/config';
import Title from './components/Title';

const Applications = lazy(() => import('./views/ApplicationsList/Applications'));
const CreateApplicationForm = lazy(() => import('./views/ApplicationsList/CreateApplicationForm'));
const App = lazy(() => import('./views/App/App'));

API.init({
  baseURL: config.apiEndpoint,
  on404: () => {
    console.log('API not found');
  }
});

ReactDOM.render(
  <React.StrictMode>
    <div className="gobench-container">
      <Title />
      <Router>
        <Switch>
          <Suspense fallback={<div />}>
            <Route exact={true} path="/application/:appId" component={App} />
            <Route exact={true} path="/application/create" component={CreateApplicationForm} />
            <Route exact={true} path="/" component={Applications} />
          </Suspense>
        </Switch>
      </Router>
    </div>
  </React.StrictMode >,
  document.getElementById('root')
);
