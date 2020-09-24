import React from 'react'

const General12v1 = () => {
  return (
    <div className="card-body">
      <div className="d-flex mb-1">
        <div className="text-uppercase font-weight-bold mr-auto">Revenue</div>
        <div>+20% Goal Reached</div>
      </div>
      <div className="d-flex mb-2">
        <div className="font-size-24 font-weight-bold mr-auto">+3,125</div>
        <div className="font-size-24">5,000</div>
      </div>
      <div className="progress">
        <div
          className="progress-bar bg-success"
          style={{
            width: '60%',
          }}
          role="progressbar"
          aria-valuenow={60}
          aria-valuemin={0}
          aria-valuemax={100}
        />
      </div>
    </div>
  )
}

export default General12v1
