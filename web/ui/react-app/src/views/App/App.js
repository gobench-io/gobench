import React, { useEffect, useState } from 'react'
import { get } from 'lodash'
import GoBenchAPI from '../../api/gobench'
import { AppContext } from '../../context'
import { useParams, useHistory } from 'react-router-dom'

import { useInterval, INTERVAL } from '../../realtimeHelpers'
import { statusColors } from '../../components/Status'
import Dashboard from './Dashboard'
import Scenario from './Scenario'
import Logs from './Logs'

const App = (props) => {
  const [appData, fetchAppData] = useState(null)
  const [fetching, setFetching] = useState(false)
  const [activeTab, setActiveTab] = useState('dashboard')
  const { appId } = useParams()
  const history = useHistory()
  const appStatus = get(appData, 'status', '')

  useEffect(() => {
    setFetching(true)
    GoBenchAPI.getAppInfo(appId).then(res => {
      setFetching(false)
      fetchAppData(res)
    })
  }, [appId])

  useInterval(() => {
    GoBenchAPI.getAppInfo(appId).then(res => {
      setFetching(false)
      fetchAppData(res)
    })
  }, appStatus === 'running' ? INTERVAL : null)

  return <div className='card'>
    {(!appData && !fetching)
      ? <div className='app'>
        <p>Loading...</p>
        </div>
      : <div key={props.match.params.appId}>
        <div className='app-header'>
          <div className='app-header-left'>
            <h2 className='application-title'>
              {get(appData, 'name', '') || ''} application benchmark
            </h2>
            <span
              className='application-status'
              style={{
                color: '#FFFFFF',
                background: statusColors[appStatus] || '#bfbfbf'
              }}
            >{get(appData, 'status', '')}
            </span>
          </div>
          <button className='btn btn-cancel' onClick={() => history.goBack()}>&lt; Back to Applications</button>
        </div>
        <div className=''>
          <ul className='tabs'>
            <li
              className={`tab-nav-item ${activeTab === 'dashboard' ? 'tab-active' : ''}`}
              onClick={() => setActiveTab('dashboard')}
            >
              Dashboard
            </li>
            <li
              className={`tab-nav-item ${activeTab === 'scenario' ? 'tab-active' : ''}`}
              onClick={() => setActiveTab('scenario')}
            >
              Scenario
            </li>
            <li
              className={`tab-nav-item ${activeTab === 'logs' ? 'tab-active' : ''}`}
              onClick={() => setActiveTab('logs')}
            >
             Logs
            </li>
          </ul>
          <AppContext.Provider value={appData}>
          <div className='tab-content'>
            {
              activeTab === 'dashboard' &&
                <div className='tab-item'>
                  <Dashboard />
                </div>
            }
            {
              activeTab === 'scenario' &&
                <div className='tab-item'>
                  <Scenario />
                </div>
            }
            {
              activeTab === 'logs' &&
                <div className='tab-item'>
                  <Logs />
                </div>
            }
          </div>
          </AppContext.Provider>
        </div>
      </div>}
  </div>
}

export default App
