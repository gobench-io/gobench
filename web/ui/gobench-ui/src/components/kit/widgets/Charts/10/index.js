import React, { useState, useEffect } from 'react'
import { Doughnut } from 'react-chartjs-2'

import data from './data.json'
import style from './style.module.scss'

const Chart10 = () => {
  const tooltip = React.createRef()
  const tooltipLabel = React.createRef()
  const tooltipValue = React.createRef()

  const [myRef, setMyRef] = useState(null)
  const [legend, setLegend] = useState(undefined)

  // eslint-disable-next-line react-hooks/exhaustive-deps
  useEffect(() => {
    const leg = generateLegend()
    setLegend(leg)
  })

  const setChartRef = element => {
    setMyRef(element)
  }

  const generateLegend = () => {
    if (!myRef) return null
    return myRef.chartInstance.generateLegend()
  }

  const createMarkup = () => {
    return { __html: legend }
  }

  const options = {
    animation: false,
    responsive: true,
    cutoutPercentage: 70,
    legend: {
      display: false,
    },
    tooltips: {
      enabled: false,
      custom: tooltipData => {
        const tooltipEl = tooltip.current
        tooltipEl.style.opacity = 1
        if (tooltipData.opacity === 0) {
          tooltipEl.style.opacity = 0
        }
      },
      callbacks: {
        label: (tooltipItem, itemData) => {
          const dataset = itemData.datasets[0]
          const value = dataset.data[tooltipItem.index]
          tooltipValue.current.innerHTML = value
          tooltipLabel.current.innerHTML = itemData.labels[tooltipItem.index]
        },
      },
    },
    legendCallback: chart => {
      const { labels } = chart.data
      let legendMarkup = []
      const dataset = chart.data.datasets[0]
      legendMarkup.push('<div class="kit__c9__chartLegend flex-shrink-0">')
      let legends = labels.map((label, index) => {
        const color = dataset.backgroundColor[index]
        return `<div class="d-flex align-items-center flex-nowrap mt-2 mb-2"><div class="tablet mr-3" style="background-color: ${color}"></div>${label}</div>`
      })
      legends = legends.join('')
      legendMarkup.push(legends)
      legendMarkup.push('</div>')
      legendMarkup = legendMarkup.join('')
      return legendMarkup
    },
  }

  return (
    <div>
      <div className="text-dark font-size-18 font-weight-bold mb-1">Profit Change</div>
      <div className="text-gray-6 mb-2">Revenue by location and date</div>
      <div className="d-flex flex-wrap align-items-center">
        <div className="mr-3 mt-3 mb-3 position-relative">
          <Doughnut
            ref={element => setChartRef(element)}
            data={data}
            options={options}
            width={140}
            height={140}
          />
          <div className={`${style.tooltip} text-gray-5 font-size-28 text-center`} ref={tooltip}>
            <div className="font-size-14 font-weight-bold text-dark" ref={tooltipLabel} />
            <div className="font-size-14 text-dark" ref={tooltipValue} />
          </div>
        </div>
        <div dangerouslySetInnerHTML={createMarkup()} />
      </div>
    </div>
  )
}

export default Chart10
