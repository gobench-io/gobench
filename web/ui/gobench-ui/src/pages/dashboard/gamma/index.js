import React from 'react'
import { Helmet } from 'react-helmet'
import ChartistGraph from 'react-chartist'
import ChartistTooltip from 'chartist-plugin-tooltips-updated'
import { Table } from 'antd'
import Chart12 from 'components/kit/widgets/Charts/12'
import Chart12v1 from 'components/kit/widgets/Charts/12v1'
import General5v1 from 'components/kit/widgets/General/5v1'
import General2 from 'components/kit/widgets/General/2'
import General2v1 from 'components/kit/widgets/General/2v1'
import General2v2 from 'components/kit/widgets/General/2v2'
import General13v1 from 'components/kit/widgets/General/13v1'
import List10 from 'components/kit/widgets/Lists/10'
import List11 from 'components/kit/widgets/Lists/11'

import {
  inboundBandwidthData,
  outboundBandwidthData,
  supportCasesTableData,
  supportCasesPieData
} from './data.json'

import styles from './style.module.scss'

const boundChartistOptions = {
  showPoint: true,
  showLine: true,
  showArea: true,
  fullWidth: true,
  showLabel: false,
  axisX: {
    showGrid: false,
    showLabel: false,
    offset: 0
  },
  axisY: {
    showGrid: false,
    showLabel: false,
    offset: 0
  },
  chartPadding: 0,
  low: 0,
  plugins: [
    ChartistTooltip({
      anchorToPoint: false,
      appendToBody: true,
      seriesName: false
    })
  ]
}

const supportCasesPieOptions = {
  donut: true,
  donutWidth: 35,
  showLabel: false,
  plugins: [
    ChartistTooltip({
      anchorToPoint: false,
      appendToBody: true,
      seriesName: false
    })
  ]
}

const supportCasesTableColumns = [
  {
    title: 'Type',
    dataIndex: 'type',
    key: 'type'
  },
  {
    title: 'Amount',
    key: 'amount',
    dataIndex: 'amount',
    render: amount => {
      if (amount === 'Negative') {
        return <span className='text-danger font-weight-bold'>{amount}</span>
      }
      return <span className='text-primary font-weight-bold'>{amount}</span>
    }
  }
]

const DashboardGamma = () => {
  return (
    <div>
      <Helmet title='Dashboard Gamma' />
      <div className='row'>
        <div className='col-xl-12'>
          <div className='row'>
            <div className='col-lg-6'>
              <div className='card'>
                <div className='card-body'>
                  <Chart12 />
                </div>
              </div>
            </div>
            <div className='col-lg-6'>
              <div className='card'>
                <div className='card-body'>
                  <Chart12v1 />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className='row'>
        <div className='col-xl-4'>
          <div className='card'>
            <General13v1 />
          </div>
          <div className='card'>
            <div className='card-body'>
              <General5v1 />
            </div>
          </div>
        </div>
        <div className='col-xl-4'>
          <div className='card'>
            <div className='card-header border-0 pb-0'>
              <div className='cui__utils__heading mb-0'>
                <strong>RECENT INVITES</strong>
              </div>
            </div>
            <div className='card-body'>
              <List10 />
            </div>
          </div>
          <div className='card'>
            <div className='card-body'>
              <General2 />
            </div>
          </div>
          <div className='card'>
            <div className='card-body'>
              <General2v1 />
            </div>
          </div>
          <div className='card'>
            <div className='card-body'>
              <General2v2 />
            </div>
          </div>
        </div>
        <div className='col-xl-4'>
          <div className='row'>
            <div className='col-xl-12'>
              <div className='card'>
                <div className='card-header border-0 pb-0'>
                  <div className='cui__utils__heading mb-0'>
                    <strong>Inbound Bandwidth</strong>
                  </div>
                </div>
                <div className='card-body'>
                  <strong className='font-size-36 text-dark'>246Gb</strong>
                </div>
                <ChartistGraph
                  data={inboundBandwidthData}
                  options={boundChartistOptions}
                  type='Line'
                  className='height-250'
                />
              </div>
            </div>
            <div className='col-xl-12'>
              <div className='graphCard card'>
                <div className='card-header border-0 pb-0'>
                  <div className='cui__utils__heading mb-0'>
                    <strong>Outbound Bandwidth</strong>
                  </div>
                </div>
                <div className='card-body'>
                  <strong className='font-size-36 text-dark'>52Gb</strong>
                </div>
                <div className='utils__chartist utils__chartist--success'>
                  <ChartistGraph
                    data={outboundBandwidthData}
                    options={boundChartistOptions}
                    type='Line'
                    className='height-250'
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className='row'>
        <div className='col-lg-6'>
          <div className='card'>
            <div className='card-header border-0 pb-0'>
              <div className='cui__utils__heading mb-0'>
                <strong className='text-uppercase font-size-16'>Support cases</strong>
              </div>
            </div>
            <div className='card-body'>
              <div className='row'>
                <div className='col-xl-6'>
                  <div className='mb-3'>
                    <Table
                      dataSource={supportCasesTableData}
                      columns={supportCasesTableColumns}
                      pagination={false}
                    />
                  </div>
                </div>
                <div className='col-xl-6'>
                  <div
                    className={`h-100 d-flex flex-column justify-content-center align-items-center ${styles.chartPieExample}`}
                  >
                    <div className='mb-4'>
                      <ChartistGraph
                        data={supportCasesPieData}
                        type='Pie'
                        options={supportCasesPieOptions}
                      />
                    </div>
                    <div className='text-center mb-4'>
                      <span className='mr-2'>
                        <span className='kit__utils__donut kit__utils__donut--success' />
                        Ready
                      </span>
                      <span className='mr-2'>
                        <span className='kit__utils__donut kit__utils__donut--primary' />
                        In Progress
                      </span>
                      <span className='mr-2'>
                        <span className='kit__utils__donut kit__utils__donut--danger' />
                        Defected
                      </span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className='col-lg-6'>
          <div className='card'>
            <div className='card-header border-0 pb-0'>
              <div className='cui__utils__heading mb-0'>
                <strong className='text-uppercase font-size-16'>Finance Stats</strong>
              </div>
            </div>
            <div className='card-body'>
              <List11 />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default DashboardGamma
