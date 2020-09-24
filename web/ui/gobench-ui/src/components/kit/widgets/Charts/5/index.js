import React from 'react'
import ChartistGraph from 'react-chartist'
import ChartistTooltip from 'chartist-plugin-tooltips-updated'
import data from './data.json'
import style from './style.module.scss'

const options = {
  stackBars: true,
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
    showLabel: false,
    offset: 0,
  },
  seriesBarDistance: 5,
  plugins: [
    ChartistTooltip({
      anchorToPoint: false,
      appendToBody: true,
      seriesName: false,
    }),
  ],
}

const Chart5 = () => {
  return (
    <div>
      <div className="text-dark font-size-18 font-weight-bold mb-1">Year Profit</div>
      <div className="text-gray-6 mb-2">Revenue by location and date</div>
      <ChartistGraph
        className={`height-200 ${style.chart}`}
        data={data}
        options={options}
        type="Bar"
      />
    </div>
  )
}

export default Chart5
