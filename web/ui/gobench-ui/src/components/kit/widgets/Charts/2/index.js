import React from 'react'
import ChartistGraph from 'react-chartist'
import ChartistTooltip from 'chartist-plugin-tooltips-updated'
import data from './data.json'

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

const Chart2 = () => {
  return (
    <div className="position-relative">
      <div className="card-body">
        <div className="text-dark font-size-18 font-weight-bold mb-1">Year Profit</div>
        <div className="text-gray-6 mb-2">Revenue by location and date</div>
        <div className="font-weight-bold font-size-36 text-dark">$437,246.00</div>
      </div>
      <ChartistGraph
        className="height-200 ct-hidden-points"
        data={data}
        options={options}
        type="Line"
      />
    </div>
  )
}

export default Chart2
