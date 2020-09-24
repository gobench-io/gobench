import React from 'react'
import { Table } from 'antd'
import data from './data.json'
import style from './style.module.scss'

const columns = [
  {
    title: 'ACTION NAME',
    dataIndex: 'actionName',
    key: 'actionName',
    className: 'bg-transparent text-gray-6',
  },
  {
    title: 'PROGRESS',
    dataIndex: 'progress',
    key: 'progress',
    className: 'text-right bg-transparent text-gray-6',
    render: text => {
      return (
        <div className="progress">
          <div
            className={`progress-bar ${text.color}`}
            style={{ width: `${text.value}%` }}
            role="progressbar"
          />
        </div>
      )
    },
  },
  {
    title: 'VALUE',
    dataIndex: 'value',
    key: 'value',
    className: 'text-right bg-transparent text-gray-6',
    render: text => <span className="font-weight-bold">{text}</span>,
  },
]

const Table1 = () => {
  return (
    <div>
      <div className={style.table}>
        <Table columns={columns} dataSource={data} pagination={false} />
      </div>
    </div>
  )
}

export default Table1
