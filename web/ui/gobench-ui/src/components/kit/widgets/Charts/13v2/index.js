import React from 'react'
import ChartistGraph from 'react-chartist'
import ChartistTooltip from 'chartist-plugin-tooltips-updated'
import Chartist from 'chartist'
import chartData from './data.json'

const Chart13v2 = () => {
  const chartOptions = {
    lineSmooth: Chartist.Interpolation.none({
      fillHoles: false,
    }),
    showPoint: true,
    showLine: true,
    showArea: true,
    fullWidth: true,
    showLabel: false,
    axisX: {
      showGrid: false,
      showLabel: false,
      offset: 0,
    },
    axisY: {
      showGrid: false,
      showLabel: false,
      offset: 0,
    },
    chartPadding: 0,
    low: 0,
    plugins: [
      ChartistTooltip({
        anchorToPoint: false,
        appendToBody: true,
        seriesName: false,
      }),
    ],
  }

  return (
    <div>
      <div className="card-body">
        <div className="font-weight-bold font-size-36 font-weight-bold text-pink">13,846$</div>
      </div>
      <div className="position-relative">
        <ChartistGraph
          data={chartData}
          options={chartOptions}
          type="Line"
          className="height-200 ct-hidden-points"
        />
      </div>
    </div>
  )
}

export default Chart13v2
