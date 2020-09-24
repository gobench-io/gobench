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
    title: 'LOCATION',
    dataIndex: 'location',
    key: 'location',
    className: 'bg-transparent text-gray-6',
    render: text => {
      return <a className="text-blue">{text}</a>
    },
  },
  {
    title: 'PTHONE',
    dataIndex: 'phone',
    key: 'phone',
    className: 'text-left text-gray-6 bg-transparent',
  },
  {
    title: 'VALUE',
    dataIndex: 'value',
    key: 'value',
    className: 'text-right bg-transparent text-gray-6',
    render: text => <span className="font-weight-bold">{text}</span>,
  },
]

const Table4 = () => {
  return (
    <div>
      <div className={style.head}>
        <div className={`${style.headItem} mb-3 pr-3`}>
          <div className={`${style.headIcon} bg-light text-dark mr-3`}>
            <i className="fe fe-menu font-size-18" />
          </div>
          <div>
            <div className="text-uppercase text-muted text-nowrap">Cross Earnings</div>
            <div className="font-weight-bold text-dark">+125,367.36</div>
          </div>
        </div>
        <div className={`${style.headItem} mb-3`}>
          <div className={`${style.headIcon} bg-light text-dark mr-3`}>
            <i className="fe fe-cloud font-size-18" />
          </div>
          <div>
            <div className="text-uppercase text-muted text-nowrap">Tax witheld</div>
            <div className="font-weight-bold text-dark">-$12,350.00</div>
          </div>
        </div>
      </div>
      <div className={style.table}>
        <Table columns={columns} dataSource={data} pagination={false} />
      </div>
    </div>
  )
}

export default Table4
