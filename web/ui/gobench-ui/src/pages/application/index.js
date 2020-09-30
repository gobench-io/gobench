import React, { useState, useEffect } from 'react'
import { Helmet } from 'react-helmet'
import { Table, Tag, Button, Popconfirm, Tooltip } from 'antd'
import { connect } from 'react-redux'
import { withRouter, Link, useHistory } from 'react-router-dom'
import { statusColors, formatTag } from 'utils/status'
import { RetweetOutlined } from '@ant-design/icons'
import moment from 'moment'

const mapStateToProps = ({ application, dispatch }) => ({ application, dispatch })

const DefaultPage = ({ application, dispatch }) => {
  const history = useHistory()
  const { list = [], loading, total } = application
  const [fetching, setFetching] = useState(false)
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
    total: 0,
    showTotal: (total, range) => `${range[0]}-${range[1]} of ${total} scenarios`
  })
  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      sorter: (a, b) => a.name.length - b.name.length
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
      sorter: (a, b) => a.name.length - b.name.length,
      render: (text, item) => (
        <>
          <Link key={item.id} to={`/applications/${item.id}`}>
            <Tag color={statusColors[text]} key={item.id}>
              {text.toUpperCase()}
            </Tag>
          </Link>
        </>
      )
    },
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      sorter: (a, b) => a.name.length - b.name.length,
      render: (text, item) => <Link key={item.id} to={`/applications/${item.id}`}>{text}</Link>
    },
    {
      title: 'Tags',
      dataIndex: 'tags',
      key: 'tags',
      render: (text) => {
        return (
          <>
            {formatTag(text).map(x => (
              <Tag color={x.color} key={x.index}>
                {x.text}
              </Tag>

            ))}
          </>
        )
      }
    },
    {
      title: 'Started at',
      dataIndex: 'started_at',
      key: 'started_at',
      sorter: (a, b) => a.name.length - b.name.length,
      render: x => {
        return moment(x).utc().format()
      }
    },
    {
      title: 'Duration',
      dataIndex: 'duration',
      key: 'duration',
      sorter: (a, b) => a.name.length - b.name.length,
      render: (x, item) => {
        const { started_at: startedAt, updated_at: updated } = item
        const start = moment(startedAt).utc() // some random moment in time (in ms)
        if (['provisioning', 'pending', 'error'].includes(item.status)) {
          return <span />
        }
        if (['finished', 'cancel'].includes(item.status)) {
          const end = moment(updated).utc() // some random moment after start (in ms)
          const diff = end.diff(start)
          const duration = moment.utc(diff).format('HH:mm:ss.SSS')
          return <span>{duration}</span>
        }
        const diff = moment.utc().diff(start)
        const duration = moment.utc(diff).format('HH:mm:ss.SSS')
        return <span>{duration}</span>
      }
    },
    {
      title: 'Action',
      dataIndex: 'action',
      key: 'action',
      render: (x, application) => {
        return (
          <div style={{ float: 'right' }} key={application.id}>
            <Button
              style={{ marginLeft: 5 }}
              type='default'
              onClick={() => clone(application)}
            >
              Clone
            </Button>
            {['running', 'pending'].includes(application.status) && (
              <Popconfirm
                title={`Are you sure cancel application ${application.name}?`}
                onConfirm={() => cancel(application.id)}
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
            {['finished', 'pending', 'error', 'cancel'].includes(application.status) && (
              <Popconfirm
                title={`Are you sure delete application ${application.name}?`}
                onConfirm={() => destroy(application.id)}
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
        )
      }
    }
  ]
  useEffect(() => {
    //   fetch data when the first time access
    if (!fetching) {
      dispatch({
        type: 'application/LIST',
        payload: { skip: pagination.current - 1, limit: pagination.pageSize }
      })
      setFetching(true)
    }
  }, [list, total])
  const onTableChange = (pagination, filters, sorter, extra) => {
    setPagination(pagination)
    dispatch({
      type: 'application/LIST',
      payload: { skip: pagination.current - 1, limit: pagination.pageSize, pagination, filters, sorter }
    })
  }
  const onSearch = (e) => {
    if (e.key === 'Enter') {
      dispatch({
        type: 'application/LIST',
        payload: { skip: 0, limit: pagination.pageSize }
      })
      return
    }
    dispatch({
      type: 'application/LIST',
      payload: { skip: 0, limit: pagination.pageSize }
    })
  }
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
      payload: { id }
    })
  }
  return (
    <>
      <div className='application' onKeyUp={onSearch}>
        <Helmet title='Applications' />

        <div className='card'>
          <div className='card-header row'>
            <div className='col-md-6'>
              <div className='cui__utils__heading mb-0'>
                <h2>Benchmark Application</h2>
              </div>
              <div className='text-muted'>A distributed benchmark tool with Golang</div>
            </div>
            <div className='col-md-6'>
              <div className='text-right'>
                <Tooltip title='Refresh'>
                  <Button
                    icon={<RetweetOutlined />}
                    style={{ marginRight: 5 }}
                    onClick={onSearch}
                  />
                </Tooltip>
                <Button type='primary' onClick={() => history.push('/applications/create')}>Create Application</Button>
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
            <Table
              dataSource={list}
              pagination={pagination}
              loading={loading}
              columns={columns}
              onChange={onTableChange}
              ellipsis
            />
          </div>
        </div>
      </div>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))
