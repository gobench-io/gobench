import React from 'react'
import { Line, Bar, Radar, Polar, Pie, Doughnut } from 'react-chartjs-2'
import { Helmet } from 'react-helmet'

const lineData = {
  labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
  datasets: [
    {
      label: 'My First dataset',
      fill: false,
      lineTension: 0.1,
      backgroundColor: 'rgba(75,192,192,0.4)',
      borderColor: 'rgba(75,192,192,1)',
      borderCapStyle: 'butt',
      borderDash: [],
      borderDashOffset: 0.0,
      borderJoinStyle: 'miter',
      pointBorderColor: 'rgba(75,192,192,1)',
      pointBackgroundColor: '#fff',
      pointBorderWidth: 1,
      pointHoverRadius: 5,
      pointHoverBackgroundColor: 'rgba(75,192,192,1)',
      pointHoverBorderColor: 'rgba(220,220,220,1)',
      pointHoverBorderWidth: 2,
      pointRadius: 1,
      pointHitRadius: 10,
      data: [65, 59, 80, 81, 56, 55, 40],
    },
  ],
}

const lineOptions = {
  scales: {
    xAxes: [
      {
        display: false,
      },
    ],
  },
}

const barData = {
  labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
  datasets: [
    {
      label: 'My First dataset',
      backgroundColor: [
        'rgba(255, 99, 132, 0.2)',
        'rgba(54, 162, 235, 0.2)',
        'rgba(255, 206, 86, 0.2)',
        'rgba(75, 192, 192, 0.2)',
        'rgba(153, 102, 255, 0.2)',
        'rgba(255, 159, 64, 0.2)',
      ],
      borderColor: [
        'rgba(255,99,132,1)',
        'rgba(54, 162, 235, 1)',
        'rgba(255, 206, 86, 1)',
        'rgba(75, 192, 192, 1)',
        'rgba(153, 102, 255, 1)',
        'rgba(255, 159, 64, 1)',
      ],
      borderWidth: 1,
      data: [65, 59, 80, 81, 56, 55, 40],
    },
  ],
}

const barOptions = {
  scales: {
    xAxes: [
      {
        stacked: true,
      },
    ],
    yAxes: [
      {
        stacked: true,
      },
    ],
  },
}

const radarData = {
  labels: ['Eating', 'Drinking', 'Sleeping', 'Designing', 'Coding', 'Cycling', 'Running'],
  datasets: [
    {
      label: 'My First dataset',
      backgroundColor: 'rgba(179,181,198,0.2)',
      borderColor: 'rgba(179,181,198,1)',
      pointBackgroundColor: 'rgba(179,181,198,1)',
      pointBorderColor: '#fff',
      pointHoverBackgroundColor: '#fff',
      pointHoverBorderColor: 'rgba(179,181,198,1)',
      data: [65, 59, 90, 81, 56, 55, 40],
    },
    {
      label: 'My Second dataset',
      backgroundColor: 'rgba(255,99,132,0.2)',
      borderColor: 'rgba(255,99,132,1)',
      pointBackgroundColor: 'rgba(255,99,132,1)',
      pointBorderColor: '#fff',
      pointHoverBackgroundColor: '#fff',
      pointHoverBorderColor: 'rgba(255,99,132,1)',
      data: [28, 48, 40, 19, 96, 27, 100],
    },
  ],
}

const radarOptions = {
  scale: {
    reverse: true,
    ticks: {
      beginAtZero: true,
    },
  },
}

const polarData = {
  datasets: [
    {
      data: [11, 16, 7, 3, 14],
      backgroundColor: ['#FF6384', '#4BC0C0', '#FFCE56', '#E7E9ED', '#36A2EB'],
      label: 'My dataset', // for legend
    },
  ],
  labels: ['Red', 'Green', 'Yellow', 'Grey', 'Blue'],
}

const polarOptions = {
  elements: {
    arc: {
      borderColor: '#4BC0C0',
    },
  },
}

const pieData = {
  labels: ['Red', 'Blue', 'Yellow'],
  datasets: [
    {
      data: [300, 50, 100],
      backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56'],
      hoverBackgroundColor: ['#FF6384', '#36A2EB', '#FFCE56'],
    },
  ],
}

const pieOptions = {}

const doughnutData = {
  labels: ['Red', 'Blue', 'Yellow'],
  datasets: [
    {
      data: [300, 50, 100],
      backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56'],
      hoverBackgroundColor: ['#FF6384', '#36A2EB', '#FFCE56'],
    },
  ],
}

const doughnutOptions = {}

const ChartsChartjs = () => {
  return (
    <div>
      <Helmet title="Charts / Chartjs" />
      <div className="kit__utils__heading">
        <h5>
          <span className="mr-3">Charts / Chartjs</span>
          <a
            href="http://www.chartjs.org/"
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
                <strong>Line Chart</strong>
              </h5>
              <div className="mb-5">
                <Line data={lineData} options={lineOptions} width={400} height={200} />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Bar Chart</strong>
              </h5>
              <div className="mb-5">
                <Bar data={barData} options={barOptions} width={400} height={200} />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Radar Chart</strong>
              </h5>
              <div className="mb-5">
                <Radar data={radarData} options={radarOptions} width={400} height={200} />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Polar Area Chart</strong>
              </h5>
              <div className="mb-5">
                <Polar data={polarData} options={polarOptions} width={400} height={200} />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Pie Chart</strong>
              </h5>
              <div className="mb-5">
                <Pie data={pieData} options={pieOptions} width={400} height={200} />
              </div>
            </div>
            <div className="col-xl-6 col-lg-12">
              <h5 className="mb-4">
                <strong>Doughnut Chart</strong>
              </h5>
              <div className="mb-5">
                <Doughnut data={doughnutData} options={doughnutOptions} width={400} height={200} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ChartsChartjs
