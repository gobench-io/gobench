import React from 'react'
import ChartistGraph from 'react-chartist'
import ChartistTooltip from 'chartist-plugin-tooltips-updated'
import data from './data.json'
import style from './style.module.scss'

const options = {
  fullWidth: true,
  chartPadding: {
    right: 0,
    left: 0,
    top: 5,
    bottom: 0,
  },
  low: 0,
  axisY: {
    showGrid: false,
    showLabel: false,
    offset: 0,
  },
  axisX: {
    showGrid: false,
  },
  seriesBarDistance: 15,
  plugins: [
    ChartistTooltip({
      anchorToPoint: false,
      appendToBody: true,
      seriesName: false,
    }),
  ],
}

const listener = {
  draw: item => {
    if (item.type === 'bar') {
      item.group.elem('line', {
        x1: item.x1,
        x2: item.x2,
        y1: item.y2,
        y2: 0,
        stroke: '#e4e9f0',
        'stroke-width': '10',
      })
    }
  },
}

const Chart3 = () => {
  return (
    <div>
      <ChartistGraph
        className="height-200"
        data={data}
        options={options}
        type="Bar"
        listener={listener}
      />
      <div className="d-flex flex-wrap">
        <div className="mr-5 mb-2">
          <div className="text-nowrap text-uppercase text-gray-4">
            <div className={`${style.donut} ${style.success}`} />
            Gross revenue
          </div>
          <div className="font-weight-bold font-size-18 text-dark">+$125,367.36</div>
        </div>
        <div className="mr-5 mb-2">
          <div className="text-nowrap text-uppercase text-gray-4">
            <div className={`${style.donut} ${style.primary}`} />
            Gross Earnings
          </div>
          <div className="font-weight-bold font-size-18 text-dark">+$125,367.36</div>
        </div>
      </div>
    </div>
  )
}

export default Chart3
