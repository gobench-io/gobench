import React, { useMemo } from 'react'
import ReactApexChart from 'react-apexcharts'

const DefaultPage = ({ series = [], unit }) => {
  // console.log(series, unit)
  const options = {
    colors: ['#0000ff', '#00ff00', '#ff9933', '#996600', '#6600cc'],
    chart: {
      type: 'line',
      animations: {
        enabled: false
      },
      brush: {
        enable: true
      },
      dataLabels: {
        enable: true
      },
      toolbar: {
        show: true,
        autoSelected: 'zoom',
        tools: {
          download: false
        }
      },
      zoom: {
        type: 'x',
        enabled: true,
        autoScaleYaxis: true
      }
    },
    dataLabels: {
      enabled: false
    },
    stroke: {
      curve: 'smooth',
      width: 1
    },
    markers: {
      size: 0
    },
    tooltip: {
      enabled: true,
      theme: 'light',
      x: {
        show: true,
        datetimeUTC: true,
        format: 'HH:mm:ss'
      }
    },
    grid: {
      show: true,
      borderColor: '#00000030',
      strokeDashArray: 5,
      yaxis: {
        lines: {
          show: true
        }
      }
    },
    xaxis: {
      type: 'datetime'
    },
    yaxis: {
      show: true,
      showAlways: true,
      forceNiceScale: true,
      axisBorder: {
        show: true
      },
      tickAmount: 4,
      title: {
        text: unit,
        rotate: 0,
        offsetX: 40,
        offsetY: -90,
        style: {
          color: undefined,
          fontSize: '11px'
        }
      },
      labels: {
        show: true,
        style: {
          color: '#000',
          fontSize: 11
        }
      }
    },
    legend: {
      show: true,
      showForSingleSeries: true,
      showForNullSeries: true,
      showForZeroSeries: true,
      horizontalAlign: series.length > 4 ? 'left' : 'center',
      fontSize: '12px',
      labels: {
        colors: '#000'
      },
      markers: {
        width: 8,
        height: 8
      }
    }
  }
  const seriesMemo = useMemo(() => series, [series])
  return (
    <>
      <div className='chart'>
        <ReactApexChart
          height='250'
          options={options}
          series={seriesMemo}
          type='line'
        />
      </div>
    </>
  )
}

export default DefaultPage
