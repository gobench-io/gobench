import React from 'react'
import { Helmet } from 'react-helmet'
import ChartistGraph from 'react-chartist'
import ChartistTooltip from 'chartist-plugin-tooltips-updated'

const animationData = {
  labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
  series: [
    [1, 2, 2.7, 0, 3, 5, 3, 4, 8, 10, 12, 7],
    [0, 1.2, 2, 7, 2.5, 9, 5, 8, 9, 11, 14, 4],
    [10, 9, 8, 6.5, 6.8, 6, 5.4, 5.3, 4.5, 4.4, 3, 2.8],
  ],
}

const animatonOptions = {
  axisX: {
    labelInterpolationFnc(value, index) {
      return index % 2 !== 0 ? !1 : value
    },
  },
}

const smilData = {
  labels: ['1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12'],
  series: [
    [12, 9, 7, 8, 5, 4, 6, 2, 3, 3, 4, 6],
    [4, 5, 3, 7, 3, 5, 5, 3, 4, 4, 5, 5],
    [5, 3, 4, 5, 6, 3, 3, 4, 5, 6, 3, 4],
    [3, 4, 5, 6, 7, 6, 4, 5, 6, 7, 6, 3],
  ],
}

const smilOptions = {
  low: 0,
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
  seq: 0,
}

const smilListener = {
  created() {
    smilOptions.seq = 0
  },
  draw(data) {
    const delays = 80
    const durations = 500

    if (data.type === 'line') {
      smilOptions.seq += 1
      data.element.animate({
        opacity: {
          begin: smilOptions.seq * delays + 1e3,
          dur: durations,
          from: 0,
          to: 1,
        },
      })
    } else if (data.type === 'label' && data.axis === 'x')
      data.element.animate({
        y: {
          begin: smilOptions.seq * delays,
          dur: durations,
          from: data.y + 100,
          to: data.y,
          easing: 'easeOutQuart',
        },
      })
    else if (data.type === 'label' && data.axis === 'y')
      data.element.animate({
        x: {
          begin: smilOptions.seq * delays,
          dur: durations,
          from: data.x - 100,
          to: data.x,
          easing: 'easeOutQuart',
        },
      })
    else if (data.type === 'point')
      data.element.animate({
        x1: {
          begin: smilOptions.seq * delays,
          dur: durations,
          from: data.x - 10,
          to: data.x,
          easing: 'easeOutQuart',
        },
        x2: {
          begin: smilOptions.seq * delays,
          dur: durations,
          from: data.x - 10,
          to: data.x,
          easing: 'easeOutQuart',
        },
        opacity: {
          begin: smilOptions.seq * delays,
          dur: durations,
          from: 0,
          to: 1,
          easing: 'easeOutQuart',
        },
      })
    else if (data.type === 'grid') {
      const pos1Animation = {
        begin: smilOptions.seq * delays,
        dur: durations,
        from: data[`${data.axis.units.pos}1`] - 30,
        to: data[`${data.axis.units.pos}1`],
        easing: 'easeOutQuart',
      }
      const pos2Animation = {
        begin: smilOptions.seq * delays,
        dur: durations,
        from: data[`${data.axis.units.pos}2`] - 100,
        to: data[`${data.axis.units.pos}2`],
        easing: 'easeOutQuart',
      }
      const ctAnimations = {}
      ctAnimations[`${data.axis.units.pos}1`] = pos1Animation
      ctAnimations[`${data.axis.units.pos}2`] = pos2Animation
      ctAnimations.opacity = {
        begin: smilOptions.seq * delays,
        dur: durations,
        from: 0,
        to: 1,
        easing: 'easeOutQuart',
      }
      data.element.animate(ctAnimations)
    }
  },
}

const lineData = {
  labels: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'],
  series: [
    [12, 9, 7, 8, 5],
    [2, 1, 3.5, 7, 3],
    [1, 3, 4, 5, 6],
  ],
}

const lineOptions = {
  fullWidth: !0,
  chartPadding: {
    right: 40,
  },
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
}

const areaData = {
  labels: [1, 2, 3, 4, 5, 6, 7, 8],
  series: [[5, 9, 7, 8, 5, 3, 5, 4]],
}

const areaOptions = {
  low: 0,
  showArea: true,
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
}

const scatterTimes = function scatter(n) {
  return Array(...new Array(n))
}

const scatterData = scatterTimes(52)
  .map(Math.random)
  .reduce(
    (scatter, rnd, index) => {
      const data = scatter
      data.labels.push(index + 1)
      data.series.forEach(series => {
        series.push(Math.random() * 100)
      })
      return data
    },
    {
      labels: [],
      series: scatterTimes(4).map(() => []),
    },
  )

const scatterOptions = {
  showLine: false,
  axisX: {
    labelInterpolationFnc(value, index) {
      return index % 13 === 0 ? `W${value}` : null
    },
  },
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
}

const scatterResponsiveOptions = [
  [
    'screen and (min-width: 640px)',
    {
      axisX: {
        labelInterpolationFnc(value, index) {
          return index % 4 === 0 ? `W${value}` : null
        },
      },
    },
  ],
]

const horizontalData = {
  labels: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
  series: [
    [5, 4, 3, 7, 5, 10, 3],
    [3, 2, 9, 5, 4, 6, 4],
  ],
}

const horizontalOptions = {
  seriesBarDistance: 10,
  reverseData: !0,
  horizontalBars: !0,
  axisY: {
    offset: 70,
  },
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
}

const biPolarLineData = {
  labels: [1, 2, 3, 4, 5, 6, 7, 8],
  series: [
    [1, 2, 3, 1, -2, 0, 1, 0],
    [-2, -1, -2, -1, -2.5, -1, -2, -1],
    [0, 0, 0, 1, 2, 2.5, 2, 1],
    [2.5, 2, 1, 0.5, 1, 0.5, -1, -2.5],
  ],
}

const biPolarLineOptions = {
  high: 3,
  low: -3,
  showArea: !0,
  showLine: !1,
  showPoint: !1,
  fullWidth: !0,
  axisX: {
    showLabel: false,
    showGrid: false,
  },
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
}

const biPolarBarData = {
  labels: ['W1', 'W2', 'W3', 'W4', 'W5', 'W6', 'W7', 'W8', 'W9', 'W10'],
  series: [[1, 2, 4, 8, 6, -2, -1, -4, -6, -2]],
}

const biPolarBarOptions = {
  high: 10,
  low: -10,
  axisX: {
    labelInterpolationFnc(value, index) {
      return index % 2 === 0 ? value : null
    },
  },
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
}

const stackedBarData = {
  labels: ['Q1', 'Q2', 'Q3', 'Q4'],
  series: [
    [8e5, 12e5, 14e5, 13e5],
    [2e5, 4e5, 5e5, 3e5],
    [1e5, 2e5, 4e5, 6e5],
  ],
}

const stackedBarOptions = {
  stackBars: !0,
  axisY: {
    labelInterpolationFnc(value) {
      return `${value / 1e3}k`
    },
  },
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
}

const overlappingBarData = {
  labels: ['Jan', 'Feb', 'Mar', 'Apr', 'Mai', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
  series: [
    [5, 4, 3, 7, 5, 10, 3, 4, 8, 10, 6, 8],
    [3, 2, 9, 5, 4, 6, 4, 6, 7, 8, 7, 4],
  ],
}

const overlappingBarOptions = {
  seriesBarDistance: 10,
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
}

const overlappingResponsiveOptions = [
  [
    '',
    {
      seriesBarDistance: 5,
      axisX: {
        labelInterpolationFnc(value) {
          return value[0]
        },
      },
    },
  ],
]

const labelsPieData = {
  labels: ['Bananas', 'Apples', 'Grapes'],
  series: [20, 15, 40],
}
const labelsPieOptions = {
  labelInterpolationFnc(value) {
    return value[0]
  },
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
}
const labelsPieResponsiveOptions = [
  [
    'screen and (min-width: 640px)',
    {
      chartPadding: 30,
      labelOffset: 100,
      labelDirection: 'explode',
      labelInterpolationFnc(value) {
        return value
      },
    },
  ],
  [
    'screen and (min-width: 1024px)',
    {
      labelOffset: 80,
      chartPadding: 20,
    },
  ],
]

const simplePieData = {
  series: [5, 3, 4],
}

const simplePieSum = function sum(a, b) {
  return a + b
}

const simplePieOptions = {
  labelInterpolationFnc(value) {
    return `${Math.round((value / simplePieData.series.reduce(simplePieSum)) * 100)}%`
  },
  plugins: [ChartistTooltip({ anchorToPoint: false, appendToBody: true, seriesName: false })],
}

const ChartsChartistjs = () => {
  return (
    <div>
      <Helmet title="Charts / Chartist.js" />
      <div className="kit__utils__heading">
        <h5>
          <span className="mr-3">Chartist.js</span>
          <a
            href="https://gionkunz.github.io/chartist-js/"
            rel="noopener noreferrer"
            target="_blank"
            className="btn btn-sm btn-light"
          >
            Official Documentation
            <i className="fe fe-corner-right-up" />
          </a>
        </h5>
      </div>
      <div className="card">
        <div className="card-body">
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>CSS Styling & Animations</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300 chart-css-animations chartist-theme-dark chartist-animated"
                  data={animationData}
                  options={animatonOptions}
                  type="Line"
                />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>SMIL Animations</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300 chart-smil-animations"
                  data={smilData}
                  options={smilOptions}
                  type="Line"
                  listener={smilListener}
                />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Line</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300"
                  data={lineData}
                  options={lineOptions}
                  type="Line"
                />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Area</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300"
                  data={areaData}
                  options={areaOptions}
                  type="Line"
                />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Scatter</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300"
                  data={scatterData}
                  options={scatterOptions}
                  responsive-options={scatterResponsiveOptions}
                  type="Line"
                />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Horizontal Bar</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300"
                  type="Bar"
                  data={horizontalData}
                  options={horizontalOptions}
                />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Bi-polar Line</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300"
                  data={biPolarLineData}
                  options={biPolarLineOptions}
                  type="Line"
                />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Bi-polar Bar</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300"
                  data={biPolarBarData}
                  options={biPolarBarOptions}
                  type="Bar"
                />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Stacked Bar</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300"
                  data={stackedBarData}
                  options={stackedBarOptions}
                  type="Bar"
                />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Overlapping Bar</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300"
                  data={overlappingBarData}
                  options={overlappingBarOptions}
                  responsive-options={overlappingResponsiveOptions}
                  type="Bar"
                />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Simple Pie</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300"
                  data={simplePieData}
                  options={simplePieOptions}
                  type="Pie"
                />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Pie w/ Labels</strong>
              </h5>
              <div className="mb-5">
                <ChartistGraph
                  className="height-300"
                  data={labelsPieData}
                  options={labelsPieOptions}
                  responsive-options={labelsPieResponsiveOptions}
                  type="Pie"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ChartsChartistjs
