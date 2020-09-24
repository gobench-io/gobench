import React from 'react'
import ChartistGraph from 'react-chartist'
import ChartistTooltip from 'chartist-plugin-tooltips-updated'
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
    title: 'LAST WEEK PROFIT',
    dataIndex: 'chart',
    key: 'chart',
    className: 'text-right text-gray-6',
    render: chart => {
      return (
        <ChartistGraph className="ct-hidden-points" data={chart} options={options} type="Line" />
      )
    },
  },
]

const options = {
  width: '110px',
  height: '50px',
  chartPadding: {
    right: 0,
    left: 0,
    top: 5,
    bottom: 5,
  },
  fullWidth: true,
  showPoint: true,
  lineSmooth: true,
  axisY: {
    showGrid: false,
    showLabel: false,
    offset: 0,
  },
  axisX: {
    showGrid: false,
    showLabel: false,
    offset: 0,
  },
  showArea: false,
  plugins: [
    ChartistTooltip({
      anchorToPoint: false,
      appendToBody: true,
      seriesName: false,
    }),
  ],
}

const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows)
  },
}

const Chart8 = () => {
  return (
    <div>
      <div className={`${style.textDivider} mb-2`}>
        <h4 className={`${style.textDividerContent} font-size-24 font-weight-bold`}>
          Waiting actions
        </h4>
      </div>
      <div className={style.table}>
        <Table
          columns={columns}
          dataSource={data.table}
          pagination={false}
          rowSelection={rowSelection}
        />
      </div>
    </div>
  )
}

export default Chart8
