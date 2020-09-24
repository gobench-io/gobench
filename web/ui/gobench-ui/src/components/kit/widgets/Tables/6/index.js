import React from 'react'
import { Table } from 'antd'
import data from './data.json'
import style from './style.module.scss'

const columns = [
  {
    title: 'USERNAME',
    dataIndex: 'userName',
    className: 'text-gray-6',
    key: 'userName',
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
    className: 'text-gray-6',
    key: 'location',
    render: text => {
      return <a className="text-blue">{text}</a>
    },
  },
  {
    title: 'VALUE',
    dataIndex: 'value',
    key: 'value',
    className: 'text-right text-gray-6',
    render: text => <span className="font-weight-bold">{text}</span>,
  },
  {
    dataIndex: 'action',
    key: 'action',
    className: 'text-right',
    render: () => {
      return (
        <div className="text-nowrap">
          <button type="button" className="btn btn-outline-danger mr-2 mb-2">
            Decline
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

const Table6 = () => {
  return (
    <div>
      <div className={`${style.textDivider} mb-2`}>
        <h4 className={`${style.textDividerContent} font-size-24 font-weight-bold`}>
          Waiting actions
        </h4>
      </div>
      <div className={`${style.table}`}>
        <Table columns={columns} dataSource={data} pagination={false} rowSelection={rowSelection} />
      </div>
    </div>
  )
}

export default Table6
