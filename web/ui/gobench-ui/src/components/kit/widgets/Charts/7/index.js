import React from 'react'
import { Table } from 'antd'
import { VectorMap } from 'react-jvectormap'
import data from './data.json'
import style from './style.module.scss'

const columns = [
  {
    title: 'Action name',
    dataIndex: 'actionName',
    key: 'actionName',
    className: 'bg-transparent text-gray-6',
  },
  {
    title: 'Location',
    dataIndex: 'location',
    key: 'location',
    className: 'bg-transparent text-gray-6',
    render: text => {
      return <a className="text-blue">{text}</a>
    },
  },
  {
    title: 'Phone',
    dataIndex: 'phone',
    key: 'phone',
    className: 'text-left text-gray-6 bg-transparent',
  },
  {
    title: 'Value',
    dataIndex: 'value',
    key: 'value',
    className: 'text-right bg-transparent text-gray-6',
    render: text => <span className="font-weight-bold">{text}</span>,
  },
]

const Chart7 = () => {
  return (
    <div>
      <div className="height-300 position-relative mb-3">
        <VectorMap
          map="us_aea"
          backgroundColor="transparent"
          containerStyle={{
            width: '100%',
            height: '100%',
          }}
          containerClassName="map"
          regionStyle={{
            initial: {
              fill: '#d1e6fa',
              'fill-opacity': 0.9,
              stroke: '#fff',
              'stroke-width': 2,
              'stroke-opacity': 0.05,
            },
            hover: {
              'fill-opacity': 0.8,
              fill: '#1b55e3',
              cursor: 'pointer',
            },
          }}
          series={{
            regions: [
              {
                attribute: 'fill',
                values: {
                  'US-CA': '#69b2f8',
                  'US-MO': '#69b2f8',
                  'US-FL': '#69b2f8',
                  'US-OR': '#69b2f8',
                  'US-TX': '#69b2f8',
                },
              },
            ],
          }}
        />
      </div>
      <div className={style.table}>
        <Table columns={columns} dataSource={data.table} pagination={false} />
      </div>
    </div>
  )
}

export default Chart7
