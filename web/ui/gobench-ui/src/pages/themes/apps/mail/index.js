import React from 'react'
import { SearchOutlined } from '@ant-design/icons'
import { Input, Tabs, Table } from 'antd'
import PerfectScrollbar from 'react-perfect-scrollbar'
import { Helmet } from 'react-helmet'
import mails from './data.json'
import style from './style.module.scss'

const { TabPane } = Tabs

const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows)
  },
}

const AppsMail = () => {
  const columns = [
    {
      title: '',
      dataIndex: 'favorites',
      key: 'favorites',
      render: text => (
        <i
          className={text === true ? 'icmn-star-full text-warning' : 'icmn-star-full text-default'}
        />
      ),
    },
    {
      title: 'From',
      dataIndex: 'from',
      key: 'from',
      render: text => (
        <a href="#" onClick={e => e.preventDefault()}>
          {text}
        </a>
      ),
      sorter: (a, b) => a.from.length - b.from.length,
    },
    {
      title: 'Message',
      dataIndex: 'message',
      key: 'message',
      sorter: (a, b) => a.message.length - b.message.length,
    },
    {
      title: '',
      dataIndex: 'attachments',
      key: 'attachments',
      render: text => {
        if (text === true) {
          return <i className="icmn-attachment text-default" />
        }
        return ''
      },
    },
    {
      title: '',
      dataIndex: 'time',
      key: 'time',
    },
  ]

  return (
    <div>
      <Helmet title="Mail" />
      <div className="row">
        <div className="col-12 col-md-3">
          <div className="mb-4">
            <Input
              prefix={<SearchOutlined style={{ color: 'rgba(0,0,0,.25)' }} />}
              placeholder="Search mail..."
            />
          </div>
          <div className={style.categories}>
            <PerfectScrollbar>
              <div className="d-flex flex-column">
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className={`${style.category} ${style.current} text-dark font-size-18 font-weight-bold`}
                >
                  <span className="text-truncate">Inbox</span>
                  <span>(2)</span>
                </a>
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className={`${style.category} text-dark font-size-18`}
                >
                  <span className="text-truncate">Snoozed</span>
                </a>
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className={`${style.category} text-dark font-size-18`}
                >
                  <span className="text-truncate">Sent</span>
                </a>
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className={`${style.category} text-dark font-size-18 font-weight-bold`}
                >
                  <span className="text-truncate">Drafts</span>
                  <span>(1)</span>
                </a>
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className={`${style.category} text-dark font-size-18`}
                >
                  <span className="text-truncate">Spam</span>
                </a>
              </div>
            </PerfectScrollbar>
          </div>
        </div>
        <div className="col-12 col-md-9">
          <div className="card">
            <div className="card-header card-header-flex">
              <Tabs defaultActiveKey="1" className="mr-auto kit-tabs-bold">
                <TabPane tab="Notifications" key="1" />
                <TabPane
                  tab={
                    <span>
                      Social
                      <span className=" ml-2 badge badge-primary text-uppercase">4 new</span>
                    </span>
                  }
                  key="2"
                />
                <TabPane tab="Primary" key="3" />
              </Tabs>
              <div className="d-inline-flex align-items-center">
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className="btn btn-sm btn-light mr-2"
                  data-toggle="tooltip"
                  data-placement="top"
                  title=""
                  data-original-title="Unlock Account"
                >
                  <i className="fe fe-unlock" />
                </a>
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className="btn btn-sm btn-light mr-2"
                  data-toggle="tooltip"
                  data-placement="top"
                  title=""
                  data-original-title="Mark as important"
                >
                  <i className="fe fe-star" />
                </a>
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className="btn btn-sm btn-light"
                  data-toggle="tooltip"
                  data-placement="top"
                  title=""
                  data-original-title="Delete user"
                >
                  <i className="fe fe-trash" />
                </a>
              </div>
            </div>
            <div className="card-body">
              <div className="kit__utils__table">
                <Table columns={columns} rowSelection={rowSelection} dataSource={mails} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default AppsMail
