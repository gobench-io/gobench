import React from 'react'
import { Table, Tabs, Dropdown, Menu } from 'antd'
import data from './data.json'
import style from './style.module.scss'

const { TabPane } = Tabs

const dropdownMenu = (
  <Menu>
    <Menu.Item>
      <a>Action</a>
    </Menu.Item>
    <Menu.Item>
      <a>Another action</a>
    </Menu.Item>
    <Menu.Item>
      <a>Something else here</a>
    </Menu.Item>
    <div className="dropdown-divider" />
    <Menu.Item>
      <a>Separated link</a>
    </Menu.Item>
  </Menu>
)

const columns = [
  {
    dataIndex: 'avatar',
    key: 'avatar',
    className: 'bg-transparent text-gray-6 width-50',
    render: text => {
      return (
        <div>
          <div className="kit__utils__avatar">
            <img src={text} alt="User avatar" />
          </div>
        </div>
      )
    },
  },
  {
    title: 'USERNAME',
    dataIndex: 'userName',
    key: 'userName',
    className: 'bg-transparent text-gray-6',
    render: text => {
      return (
        <div>
          <div>{text.name}</div>
          <div className="text-gray-4">{text.position}</div>
        </div>
      )
    },
  },
  {
    title: 'LOCATION',
    dataIndex: 'location',
    key: 'location',
    className: 'bg-transparent text-gray-6',
    render: text => {
      return <a className="text-blue">{text}</a>
    },
  },
  {
    dataIndex: 'action',
    key: 'action',
    className: 'bg-transparent text-right',
    render: () => {
      return (
        <div>
          <button type="button" className="btn btn-primary mr-2">
            <i className="fe fe-plus-circle" />
          </button>
          <button type="button" className="btn btn-light">
            <i className="fe fe-settings text-blue" />
          </button>
        </div>
      )
    },
  },
]

const Table3 = () => {
  return (
    <div>
      <Tabs className={style.tabs} defaultActiveKey="1">
        <TabPane tab="History" key="1" />
        <TabPane
          tab={
            <Dropdown overlay={dropdownMenu} placement="bottomRight">
              <a className="nav-link dropdown-toggle" role="button">
                Dropdown
              </a>
            </Dropdown>
          }
          key="2"
        />
        <TabPane tab="Actions" key="3" />
      </Tabs>
      <div className="pb-4 px-4">
        <div className={style.table}>
          <Table columns={columns} dataSource={data} pagination={false} />
        </div>
      </div>
    </div>
  )
}

export default Table3
