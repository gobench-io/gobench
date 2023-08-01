import React, { useState, useEffect, lazy } from 'react'
import { Helmet } from 'react-helmet'
import { Tabs, Tag, Button, Popconfirm } from 'antd'
import { connect } from 'react-redux'
import { withRouter, useParams, useHistory } from 'react-router-dom'
import Dashboard from './dashboard'
import Scenario from './scenario'
import { statusColors } from 'utils/status'
import { INTERVAL } from 'constant'
import 'css/index.css'
import './style.scss'
import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc'
import duration from 'dayjs/plugin/duration'
dayjs.extend(utc)
dayjs.extend(duration)

const UserLog = lazy(() => import('./user-log'))
const SystemLog = lazy(() => import('./system-log'))
const Tags = lazy(() => import('./tags'))

const mapStateToProps = ({ application, dispatch }) => {
  const { detail } = application
  return { detail, dispatch }
}
const DefaultPage = ({ detail, dispatch }) => {
  const [fetching, setFetching] = useState(false)
  const [tab, setTab] = useState(1)
  const history = useHistory()
  const { id } = useParams()
  const { name, created_at: created, status, started_at: startedAt, updated_at: finishedAt } = detail
  const start = dayjs(startedAt).utc()
  let end = dayjs(finishedAt).utc()
  if (status === 'running') {
    end = dayjs().utc()
  }
  if (['pending', 'error', 'provisioning'].includes(status)) {
    end = start
  }
  const diff = startedAt ? end.diff(start) : 0
  // execution
  const duration = dayjs.duration(diff).format('HH:mm:ss.SSS')
  useEffect(() => {
    if (!fetching) {
      dispatch({
        type: 'application/DETAIL',
        payload: { id }
      })
      setFetching(true)
    }
  }, [detail])
  useEffect(() => {
    if (['pending', 'provisioning'].includes(status)) {
      const interval = setInterval(() => {
        dispatch({
          type: 'application/DETAIL',
          payload: { id }
        })
      }, INTERVAL / 3)
      // destroy interval on unmount
      return () => clearInterval(interval)
    }
  })
  const clone = (data) => {
    dispatch({
      type: 'application/CLONE',
      payload: { data }
    })
  }
  const cancel = (id) => {
    dispatch({
      type: 'application/CANCEL',
      payload: { id }
    })
  }
  const destroy = (id) => {
    dispatch({
      type: 'application/DELETE',
      payload: { id, redirect: '/' }
    })
  }
  const onChange = (value) => {
    setTab(value)
  }
  return (
    <>
      <div className='application-detail'>
        <Helmet title='Application|Detail' />
        <div className='card'>
          <div className='card-header'>
            <div className='row'>
              <div className='col-md-6'>
                <div className='cui__utils__heading mb-0'>
                  <h2>{name}</h2>
                  <Tag color={statusColors[status]}>
                    {(status || '').toUpperCase()}
                  </Tag>
                </div>
                <div className='text-muted'>Created: <strong>{dayjs(created).utc().format()} UTC</strong></div>
                <div className='text-muted'>Started: <strong>{start.format()} UTC</strong></div>
                <div className='text-muted'>Ended: <strong>{['pending', 'error', 'provisioning', 'running'].includes(status) ? <i>not finish yet</i> : `${end.format()} UTC`}</strong></div>
                <div className='text-muted'>Duration: <strong>{duration}</strong></div>
              </div>
              <div className='col-md-6'>
                <div className='text-end'>
                  <div style={{ float: 'right' }} key={detail.id}>
                    <Button
                      style={{ marginLeft: 5 }}
                      type='default'
                      onClick={() => clone(detail)}
                    >
                      Clone
                    </Button>
                    {['running', 'pending'].includes(detail.status) && (
                      <Popconfirm
                        title={`Are you sure cancel application ${detail.name}?`}
                        onConfirm={() => cancel(detail.id)}
                        okText='Yes'
                        cancelText='No'
                      >
                        <Button
                          type='dashed'
                          style={{ marginLeft: 5 }}
                          danger
                        >
                          Cancel
                        </Button>
                      </Popconfirm>
                    )}
                    {['finished', 'pending', 'error', 'cancel'].includes(detail.status) && (
                      <Popconfirm
                        title={`Are you sure delete application ${detail.name}?`}
                        onConfirm={() => destroy(detail.id)}
                        okText='Yes'
                        cancelText='No'
                      >
                        <Button
                          type='primary'
                          className='delete-button'
                          style={{ marginLeft: 5, color: 'white', backgroundColor: '#f5222d!important' }}
                          danger
                        >
                          Delete
                        </Button>
                      </Popconfirm>
                    )}
                  </div>
                  <Button type='default' onClick={() => history.push('/applications')}>Back</Button>
                </div>
              </div>
            </div>
            <div className='row'>
              <div className='application-tag'>
                <Tags />
              </div>
            </div>
          </div>
          <div className='card-body'>
            <div className='application-header'>
              <div className='search-bar'>
                <div className='row'>
                  <div className='col-md-3 col-xs-12'>
                    {/* <Search placeholder='input application name or tags' onSearch={value => onSearch(value)}>Search</Search> */}
                  </div>
                </div>
              </div>
            </div>
            <Tabs activeKey={tab} defaultActiveKey='1' size='large' onTabClick={onChange} items={[
              { key: 1, label: 'Dashboard', children: <Dashboard />, forceRender: true },
              { key: 2, label: 'Scenario', children: <Scenario /> },
              { key: 3, label: 'User Log', children: <UserLog /> },
              { key: 4, label: 'System Log', children: <SystemLog /> }
            ]}>
            </Tabs>
          </div>
        </div>
      </div>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))
