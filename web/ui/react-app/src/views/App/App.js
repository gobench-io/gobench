import React, { useEffect, useState, useCallback } from 'react'
import { get } from 'lodash'
import GoBenchAPI from '../../api/gobench'
import { AppContext } from '../../context'
import { useParams, useHistory } from 'react-router-dom'
import { Modal, Button, Tag } from 'antd'
import Tags from './Tag'
import { useInterval, INTERVAL } from '../../realtimeHelpers'
import { statusColors,colorFull } from '../../components/Status'
import Dashboard from './Dashboard'
import Scenario from './Scenario'
import Logs from './Logs'
import './style.scss'

const App = (props) => {
  const [appData, fetchAppData] = useState({})
  const [tags, setTags] = useState([])
  const [fetching, setFetching] = useState(false)
  const [showModal, setShowModal] = useState(false)
  const [activeTab, setActiveTab] = useState('dashboard')
  const { appId } = useParams()
  const history = useHistory()
  const appStatus = get(appData, 'status', '')

  const setApplicationTags = useCallback(() => {
    const appTags = tags.join(',')
    GoBenchAPI.setApplicationTags(appId, appTags).then((res) => {
      setFetching(false)
      fetchAppData(res)
      setShowModal(false)
    })
  })

  useEffect(() => {
    setFetching(true)
    GoBenchAPI.getAppInfo(appId).then((res) => {
      setFetching(false)
      fetchAppData(res)
    })
  }, [appId])

  useInterval(
    () => {
      GoBenchAPI.getAppInfo(appId).then((res) => {
        setFetching(false)
        fetchAppData(res)
      })
    },
    appStatus === 'running' ? INTERVAL : null
  )

  return (
    <div className='card app-detail'>
    <AppContext.Provider value={appData} >
      {!appData && !fetching ? (
        <div className='app'>
          <p>Loading...</p>
        </div>
      ) : (
        <div key={props.match.params.appId}>
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
              >
                {get(appData, 'status', '')}
              </span>
            </div>
            <div className='header-right'>
              <Button type='primary' style={{ marginRight: '10px' }} onClick={() => setShowModal(true)}>
                Manage Tags
              </Button>
              <Button type='ghost' onClick={() => history.goBack()}>
                &lt; Back to Applications
              </Button>
            </div>
          </div>
          <div className='app-small-timestamp'>
            {appData.tags && appData.tags.split(',').filter(x => x).map((item, index) => (
              <Tag key={index} color={colorFull()}>{item}</Tag>
            ))}
            <small>{get(appData, 'created_at', '')}</small>
          </div>
          <div className=''>
            <ul className='tabs'>
              <li
                className={`tab-nav-item ${
                    activeTab === 'dashboard' ? 'tab-active' : ''
                }`}
                onClick={() => setActiveTab('dashboard')}
              >
                  Dashboard
              </li>
              <li
                className={`tab-nav-item ${
                    activeTab === 'scenario' ? 'tab-active' : ''
                }`}
                onClick={() => setActiveTab('scenario')}
              >
                  Scenario
              </li>
              <li
                className={`tab-nav-item ${
                    activeTab === 'logs' ? 'tab-active' : ''
                }`}
                onClick={() => setActiveTab('logs')}
              >
                  Logs
              </li>
            </ul>
            <div className='tab-content'>
              {activeTab === 'dashboard' && (
                <div className='tab-item'>
                  <Dashboard />
                </div>
              )}
              {activeTab === 'scenario' && (
                <div className='tab-item'>
                  <Scenario />
                </div>
              )}
              {activeTab === 'logs' && (
                <div className='tab-item'>
                  <Logs {...props} />
                </div>
              )}
            </div>
          </div>
        </div>
      )}
      <Modal
        title='Manage Application Tags'
        visible={showModal}
        onOk={() => setApplicationTags(appId, tags)}
        onCancel={() => setShowModal(false)}
      >
        <Tags tags={tags} setTags={setTags} />
      </Modal>
    </AppContext.Provider>
    </div>
  )
}

export default App
