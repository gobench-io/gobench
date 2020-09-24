import React from 'react'
import { Table } from 'antd'
import data from './data.json'
import style from './style.module.scss'

const columns = [
  {
    title: 'ACTION NAME',
    dataIndex: 'actionName',
    key: 'actionName',
    className: 'text-gray-6',
  },
  {
    title: 'LOCATION',
    dataIndex: 'location',
    key: 'location',
    className: 'text-gray-6',
    render: text => {
      return <a className="text-blue">{text}</a>
    },
  },
  {
    title: 'VALUE',
    dataIndex: 'value',
    key: 'value',
    className: 'text-gray-6',
    render: text => <span className="font-weight-bold">{text}</span>,
  },
  {
    title: 'DESCRIPTION',
    dataIndex: 'description',
    key: 'description',
    className: 'text-gray-6',
  },
  {
    dataIndex: 'users',
    key: 'users',
    render: users => {
      return (
        <div className={`kit__utils__avatarGroup ${users.length ? '' : 'd-none'}`}>
          {users.map(user => (
            <div key={Math.random()} className="kit__utils__avatar kit__utils__avatar--rounded">
              <img src={user} alt="User Avatar" />
            </div>
          ))}
          <button type="button" className="kit__utils__avatarGroupAdd">
            <i className="fe fe-plus" />
          </button>
        </div>
      )
    },
  },
  {
    dataIndex: 'action',
    key: 'action',
    className: 'text-right',
    render: () => {
      return (
        <div className="text-nowrap">
          <button type="button" className="btn btn-outline-success mr-2 mb-2">
            Accept
          </button>
        </div>
      )
    },
  },
]

const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows)
  },
}

const Table7 = () => {
  return (
    <div>
      <div className={`${style.textDivider} mb-2`}>
        <h4 className={`${style.textDividerContent} font-size-24 font-weight-bold`}>
          Active Users
        </h4>
      </div>
      <div className={style.table}>
        <Table columns={columns} dataSource={data} pagination={false} rowSelection={rowSelection} />
      </div>
    </div>
  )
}

export default Table7
