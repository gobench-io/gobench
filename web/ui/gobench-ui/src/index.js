
// Bootstrap CSS 
import 'bootstrap/dist/css/bootstrap.min.css'
import 'antd/dist/reset.css'
import './global.scss' // app & third-party component styles 

import React from 'react'
import { createRoot } from 'react-dom/client'
import { createHashHistory } from 'history'
import { createStore, applyMiddleware, compose } from 'redux'
import { Provider } from 'react-redux'
// import { logger } from 'redux-logger'
import createSagaMiddleware from 'redux-saga'
import { routerMiddleware } from 'connected-react-router'

import reducers from './redux/reducers'
import sagas from './redux/sagas'
import Localization from './localization'
import Router from './router'
import * as serviceWorker from './serviceWorker'
import './polifill'
import { ConfigProvider } from 'antd'

// middlewared
const history = createHashHistory()
const sagaMiddleware = createSagaMiddleware()
const routeMiddleware = routerMiddleware(history)
const middlewares = [sagaMiddleware, routeMiddleware]
if (process.env.NODE_ENV === 'development') {
  // middlewares.push(logger)
}
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose

const store = createStore(reducers(history), composeEnhancers(applyMiddleware(...middlewares)))
sagaMiddleware.run(sagas)

const root = createRoot(
  document.getElementById('root')
)

root.render(
  <ConfigProvider
    theme={{
      token: {
        colorPrimary: '#1677ff',
        borderRadius: 2,
        fontFamily: 'Mukta'
      },
    }}
  >
    <Provider store={store}>
      <Localization>
        <Router history={history} />
      </Localization>
    </Provider>
  </ConfigProvider>
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister()
export { store, history }
