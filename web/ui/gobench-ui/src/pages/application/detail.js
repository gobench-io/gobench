import React, { useState, useEffect } from 'react'
import { Helmet } from 'react-helmet'
import { Tabs, Tag, Button } from 'antd'
import { connect } from 'react-redux'
import { withRouter, useParams, useHistory } from 'react-router-dom'
import Dashboard from './dashboard'
import Scenario from './scenario'
import { statusColors } from 'utils/status'
import { INTERVAL } from 'constant'
import 'css/index.css'

const { TabPane } = Tabs

const mapStateToProps = ({ application, dispatch }) => {
  const { detail } = application
  return { detail, dispatch }
}
const DefaultPage = ({ detail, dispatch }) => {
  const [fetching, setFetching] = useState(false)
  const history = useHistory()
  const { id } = useParams()
  const { name, created_at: created, status, started_at: beginAt, ended_at: endAt } = detail
  const duration = (new Date(endAt)) - (new Date(beginAt)) | 0
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
    const interval = setInterval(() => {
      console.log(Date.now())
    }, INTERVAL)
    // destroy interval on unmount
    return () => clearInterval(interval)
  })
  return (
    <>
      <div className='application-detail'>
        <Helmet title='Application|Detail' />
        <div className='card'>
          <div className='card-header row'>
            <div className='col-md-6'>
              <div className='cui__utils__heading mb-0'>
                <h2>{name}</h2>
                <Tag color={statusColors[status]}>
                  {(status || '').toUpperCase()}
                </Tag>
              </div>
              <div className='text-muted'>Created: <strong>{created}</strong></div>
              <div className='text-muted'>Started: <strong>{beginAt}</strong></div>
              <div className='text-muted'>Ended: <strong>{endAt}</strong></div>
              <div className='text-muted'>Duration: <strong>{duration}</strong></div>
            </div>
            <div className='col-md-6'>
              <div className='text-right'>
                <Button type='default' onClick={() => history.push('/applications')}>Back</Button>
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
            <Tabs defaultActiveKey='1' size='large'>
              <TabPane tab='Dashboard' key='1'>
                <Dashboard />
              </TabPane>
              <TabPane tab='Scenario' key='2'>
                <Scenario />
              </TabPane>
              <TabPane tab='Log' key='3'>
                Come in soon
              </TabPane>
            </Tabs>
          </div>
        </div>
      </div>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))
