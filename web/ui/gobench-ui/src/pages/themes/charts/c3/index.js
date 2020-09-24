import React from 'react'
import { Helmet } from 'react-helmet'
import C3Chart from 'react-c3js'

const ChartsC3 = () => {
  const colors = {
    primary: '#01a8fe',
    def: '#acb7bf',
    success: '#46be8a',
    danger: '#fb434a',
  }

  const line = {
    data: {
      columns: [
        ['Primary', 100, 165, 140, 270, 200, 140, 220],
        ['Success', 110, 80, 100, 85, 125, 90, 100],
      ],
    },
    color: {
      pattern: [colors.primary, colors.success],
    },
    axis: {
      x: {
        tick: {
          outer: !1,
        },
      },
      y: {
        max: 300,
        min: 0,
        tick: {
          outer: !1,
          count: 7,
          values: [0, 50, 100, 150, 200, 250, 300],
        },
      },
    },
    grid: {
      x: {
        show: !1,
      },
      y: {
        show: !0,
      },
    },
  }
  const spline = {
    data: {
      columns: [
        ['Primary', 100, 165, 140, 270, 200, 140, 220],
        ['Danger', 110, 80, 100, 85, 125, 90, 100],
      ],
      type: 'spline',
    },
    color: {
      pattern: [colors.primary, colors.danger],
    },
    axis: {
      x: {
        tick: {
          outer: !1,
        },
      },
      y: {
        max: 300,
        min: 0,
        tick: {
          outer: !1,
          count: 7,
          values: [0, 50, 100, 150, 200, 250, 300],
        },
      },
    },
    grid: {
      x: {
        show: !1,
      },
      y: {
        show: !0,
      },
    },
  }
  const scatter = {
    data: {
      xs: {
        Danger: 'Danger_x',
        Primary: 'Primary_x',
      },
      columns: [
        [
          'Danger_x',
          3.5,
          3,
          3.2,
          3.1,
          3.6,
          3.9,
          3.4,
          3.4,
          2.9,
          3.1,
          3.7,
          3.4,
          3,
          3,
          4,
          4.2,
          3.9,
          3.5,
          3.8,
          3.8,
          3.4,
          3.7,
          3.6,
          3.3,
          3.4,
          3,
          3.4,
          3.5,
          3.4,
          3.2,
          3.1,
          3.4,
          4.1,
          4.2,
          3.1,
          3.2,
          3.5,
          3.6,
          3,
          3.4,
          3.5,
          2.3,
          3.2,
          3.5,
          3.8,
          3,
          3.8,
          3.2,
          3.7,
          3.3,
        ],
        [
          'Primary_x',
          3.2,
          3.2,
          3.1,
          2.3,
          2.8,
          2.8,
          3.3,
          2.4,
          2.9,
          2.7,
          2,
          3,
          2.2,
          2.9,
          2.9,
          3.1,
          3,
          2.7,
          2.2,
          2.5,
          3.2,
          2.8,
          2.5,
          2.8,
          2.9,
          3,
          2.8,
          3,
          2.9,
          2.6,
          2.4,
          2.4,
          2.7,
          2.7,
          3,
          3.4,
          3.1,
          2.3,
          3,
          2.5,
          2.6,
          3,
          2.6,
          2.3,
          2.7,
          3,
          2.9,
          2.9,
          2.5,
          2.8,
        ],
        [
          'Danger',
          0.2,
          0.2,
          0.2,
          0.2,
          0.2,
          0.4,
          0.3,
          0.2,
          0.2,
          0.1,
          0.2,
          0.2,
          0.1,
          0.1,
          0.2,
          0.4,
          0.4,
          0.3,
          0.3,
          0.3,
          0.2,
          0.4,
          0.2,
          0.5,
          0.2,
          0.2,
          0.4,
          0.2,
          0.2,
          0.2,
          0.2,
          0.4,
          0.1,
          0.2,
          0.2,
          0.2,
          0.2,
          0.1,
          0.2,
          0.2,
          0.3,
          0.3,
          0.2,
          0.6,
          0.4,
          0.3,
          0.2,
          0.2,
          0.2,
          0.2,
        ],
        [
          'Primary',
          1.4,
          1.5,
          1.5,
          1.3,
          1.5,
          1.3,
          1.6,
          1,
          1.3,
          1.4,
          1,
          1.5,
          1,
          1.4,
          1.3,
          1.4,
          1.5,
          1,
          1.5,
          1.1,
          1.6,
          1.3,
          1.5,
          1.2,
          1.3,
          1.4,
          1.4,
          1.2,
          1.5,
          1,
          1.1,
          1,
          1.2,
          1.6,
          1.5,
          1.6,
          1.5,
          1.3,
          1.3,
          1.3,
          1.2,
          1.4,
          1.2,
          1,
          1.3,
          1.2,
          1.3,
          1.3,
          1.1,
          1.3,
        ],
      ],
      type: 'scatter',
    },
    color: {
      pattern: [colors.danger, colors.primary],
    },
    axis: {
      x: {
        label: 'Width',
        tick: {
          outer: !1,
          fit: !1,
        },
      },
      size: {
        height: 400,
      },
      padding: {
        right: 40,
      },
      y: {
        label: 'Height',
        tick: {
          outer: !1,
          count: 5,
          values: [0, 0.4, 0.8, 1.2, 1.6],
        },
      },
    },
    grid: {
      x: {
        show: !1,
      },
      y: {
        show: !0,
      },
    },
  }
  const bar = {
    data: {
      columns: [
        ['Danger', 30, 200, 100, 400, 150, 250],
        ['Default', 130, 100, 140, 200, 150, 50],
        ['Primary', 130, -150, 200, 300, -200, 100],
      ],
      type: 'bar',
    },
    bar: {
      width: {
        max: 20,
      },
    },
    color: {
      pattern: [colors.danger, colors.def, colors.primary],
    },
    grid: {
      y: {
        show: !0,
      },
      x: {
        show: !1,
      },
    },
  }
  const stackedBar = {
    data: {
      columns: [
        ['Primary', -30, 200, 300, 400, -150, 250, -30, 200, 300, 400, -150, 250],
        ['Default', 130, 100, -400, 100, -150, 50, 130, 100, -400, 100, -150, 50],
        ['Danger', -230, 200, 200, -300, 250, 250, -230, 200, 200, -300, 250, 250],
        ['Success', 100, -250, 150, 200, -300, -100, 100, -250, 150, 200, -300, -100],
      ],
      type: 'bar',
      groups: [['Primary', 'Default', 'Danger', 'Success']],
    },
    color: {
      pattern: [colors.primary, colors.def, colors.danger, colors.success],
    },
    bar: {
      width: {
        max: 45,
      },
    },
    grid: {
      y: {
        show: !0,
        lines: [
          {
            value: 0,
          },
        ],
      },
    },
  }
  const combination = {
    data: {
      columns: [
        ['Primary', 30, 20, 50, 40, 60, 50, 30, 20, 50, 40, 60, 50],
        ['Default', 200, 130, 90, 240, 130, 220, 200, 130, 90, 240, 130, 220],
        ['Success', 300, 200, 160, 400, 250, 250, 300, 200, 160, 400, 250, 250],
        ['Danger', 200, 130, 90, 240, 130, 220, 200, 130, 90, 240, 130, 220],
        ['Primary', 130, 120, 150, 140, 160, 150, 130, 120, 150, 140, 160, 150],
        ['Danger2', 90, 70, 20, 50, 60, 120, 90, 70, 20, 50, 60, 120],
      ],
      type: 'bar',
      types: {
        Success: 'spline',
        Danger: 'line',
        Danger2: 'area',
      },
      groups: [['Primary', 'Default']],
    },
    color: {
      pattern: [colors.primary, colors.def, colors.success, colors.danger, colors.danger],
    },
    grid: {
      x: {
        show: !1,
      },
      y: {
        show: !0,
      },
    },
  }
  const subChart = {
    data: {
      columns: [
        ['Primary', 100, 165, 140, 270, 200, 140, 220, 210, 190, 100, 170, 250],
        ['Success', 110, 80, 100, 85, 125, 90, 100, 130, 120, 90, 100, 115],
      ],
      type: 'spline',
    },
    color: {
      pattern: [colors.primary, colors.danger],
    },
    subchart: {
      show: true,
    },
  }
  const zoom = {
    data: {
      columns: [
        [
          'Sample',
          30,
          200,
          100,
          400,
          150,
          250,
          150,
          200,
          170,
          240,
          350,
          150,
          100,
          400,
          150,
          250,
          150,
          200,
          170,
          240,
          100,
          150,
          250,
          150,
          200,
          170,
          240,
          30,
          200,
          100,
          400,
          150,
          250,
          150,
          200,
          170,
          240,
          350,
          150,
          100,
          400,
          350,
          220,
          250,
          300,
          270,
          140,
          150,
          90,
          150,
          50,
          120,
          70,
          40,
        ],
      ],
      colors: {
        Sample: colors.primary,
      },
    },
    zoom: {
      enabled: !0,
    },
  }
  const pie = {
    data: {
      columns: [
        ['Primary', 30],
        ['Success', 120],
      ],
      type: 'pie',
    },
    color: {
      pattern: [colors.primary, colors.success],
    },
  }
  const donut = {
    data: {
      columns: [
        ['Danger', 30],
        ['Success', 120],
      ],
      type: 'donut',
    },
    color: {
      pattern: [colors.danger, colors.success],
    },
    donut: {
      title: 'Connections',
    },
  }
  const step = {
    data: {
      columns: [
        ['Primary', 300, 350, 300, 0, 0, 100],
        ['Success', 130, 100, 140, 200, 150, 50],
      ],
      types: {
        Primary: 'step',
        Success: 'area-step',
      },
    },
    color: {
      pattern: [colors.primary, colors.success],
    },
  }
  const area = {
    data: {
      columns: [
        ['Primary', 300, 350, 300, 0, 0, 0],
        ['Success', 130, 100, 140, 200, 150, 50],
      ],
      types: {
        Primary: 'area',
        Success: 'area-spline',
      },
    },
    color: {
      pattern: [colors.primary, colors.success],
    },
  }

  return (
    <div>
      <Helmet title="Charts / C3" />
      <div className="kit__utils__heading">
        <h5>
          <span className="mr-3">C3</span>
          <a
            href="http://c3js.org/"
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
                <strong>Simple Line</strong>
              </h5>
              <div className="mb-5">
                <C3Chart data={line.data} color={line.color} axis={line.axis} grid={line.grid} />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Spline</strong>
              </h5>
              <div className="mb-5">
                <C3Chart
                  data={spline.data}
                  color={spline.color}
                  axis={spline.axis}
                  grid={spline.grid}
                />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Step</strong>
              </h5>
              <div className="mb-5">
                <C3Chart data={step.data} color={step.color} />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Area</strong>
              </h5>
              <div className="mb-5">
                <C3Chart data={area.data} color={area.color} />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Scatter</strong>
              </h5>
              <div className="mb-5">
                <C3Chart
                  data={scatter.data}
                  color={scatter.color}
                  axis={scatter.axis}
                  grid={scatter.grid}
                />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Bar</strong>
              </h5>
              <div className="mb-5">
                <C3Chart data={bar.data} color={bar.color} grid={bar.grid} bar={bar.bar} />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Stacked Bar</strong>
              </h5>
              <div className="mb-5">
                <C3Chart
                  data={stackedBar.data}
                  color={stackedBar.color}
                  bar={stackedBar.bar}
                  grid={stackedBar.grid}
                />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Combination</strong>
              </h5>
              <div className="mb-5">
                <C3Chart
                  data={combination.data}
                  color={combination.color}
                  grid={combination.grid}
                />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Sub Chart</strong>
              </h5>
              <p className="text-muted">
                Element: read{' '}
                <a href="http://c3js.org/" target="_blank" rel="noopener noreferrer">
                  official documentation
                  <small>
                    <i className="fe fe-link ml-1" />
                  </small>
                </a>
              </p>
              <div className="mb-5">
                <C3Chart data={subChart.data} color={subChart.color} subchart={subChart.subchart} />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Zoom</strong>
              </h5>
              <div className="mb-5">
                <C3Chart data={zoom.data} color={zoom.color} zoom={zoom.zoom} />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Pie Chart</strong>
              </h5>
              <div className="mb-5">
                <C3Chart data={pie.data} color={pie.color} />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Donut Chart</strong>
              </h5>
              <div className="mb-5">
                <C3Chart data={donut.data} color={donut.color} donut={donut.donut} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ChartsC3
