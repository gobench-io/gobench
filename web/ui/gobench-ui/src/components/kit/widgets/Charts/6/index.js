import React from 'react'
import ChartistGraph from 'react-chartist'
import ChartistTooltip from 'chartist-plugin-tooltips-updated'
import data from './data.json'
import style from './style.module.scss'

const options = {
  low: 0,
  chartPadding: {
    right: 0,
    left: 0,
    top: 5,
    bottom: 0,
  },
  fullWidth: true,
  showPoint: true,
  lineSmooth: false,
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
  showArea: true,
  plugins: [
    ChartistTooltip({
      anchorToPoint: false,
      appendToBody: true,
      seriesName: false,
    }),
  ],
}

const Chart6 = () => {
  return (
    <div>
      <div className="card-body">
        <div className="text-dark font-size-18 font-weight-bold mb-1">Income Progress</div>
        <div className="text-gray-6 mb-2">Revenue by location and date</div>
        <div className="font-weight-bold font-size-36 text-dark mb-2">$390,012.01</div>
        <div className="d-flex align-items-center">
          <div className={`${style.progressIcon} bg-gray-4 text-white mr-3`}>
            <i className="fe fe-menu font-size-18" />
          </div>
          <div className="flex-grow-1">
            <div className="text-dark font-size-18 font-weight-bold text-nowrap mb-2">
              78% from $500,000.00
            </div>
            <div className="progress">
              <div className="progress-bar bg-success" style={{ width: '70%' }} />
            </div>
          </div>
        </div>
      </div>
      <ChartistGraph
        className={`height-200 ct-hidden-points ${style.chart}`}
        data={data}
        options={options}
        type="Line"
      />
    </div>
  )
}

export default Chart6
