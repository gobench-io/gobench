import React from 'react'
import { Table } from 'antd'
import data from './data.json'
import style from './style.module.scss'

const columns = [
  {
    title: 'DESCRIPTION',
    dataIndex: 'description',
    key: 'description',
    className: 'bg-transparent text-gray-6',
    render: text => {
      return (
        <div className="text-wrap width-300">
          <div className="text-dark mb-3">{text.title}</div>
          <div>{text.content}</div>
        </div>
      )
    },
  },
  {
    title: 'LOCATION',
    dataIndex: 'location',
    key: 'location',
    className: 'text-right bg-transparent text-gray-6',
    render: text => {
      return <a className="text-blue">{text}</a>
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

const Table2 = () => {
  return (
    <div>
      <div className={style.table}>
        <Table columns={columns} dataSource={data} pagination={false} />
      </div>
      <div className="mt-4 d-flex align-items-center flex-wrap">
        <button type="button" className="btn btn-primary mr-2 mb-2">
          Save
        </button>
        <button type="button" className="btn btn-link mb-2">
          Cancel
        </button>
      </div>
    </div>
  )
}

export default Table2
