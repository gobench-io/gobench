import React, { lazy, Suspense } from 'react'
import ReactDOM from 'react-dom'
import './css/index.css'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  HashRouter
} from 'react-router-dom'

import API from './api/api'
import config from './config/config'
import Layout from './components/Layout'
import Notification from './components/Notification'

const Applications = lazy(() => import('./views/ApplicationsList/Applications'))
const CreateApplicationForm = lazy(() => import('./views/ApplicationsList/CreateApplicationForm'))
const App = lazy(() => import('./views/App/App'))

API.init({
  baseURL: config.apiEndpoint,
  on404: () => {
    console.log('API not found')
  }
})

ReactDOM.render(
  <React.StrictMode>
    <Router>
      <HashRouter>
        <Switch>
          <Notification>
            <Layout>
              <Suspense fallback={<div />}>
                <Route exact path='/application/create' component={CreateApplicationForm} />
                <Route exact path='/application/:appId' component={App} />
                <Route exact path='/' component={Applications} />
              </Suspense>
            </Layout>
          </Notification>
        </Switch>
      </HashRouter>
    </Router>
  </React.StrictMode>,
  document.getElementById('root')
)
