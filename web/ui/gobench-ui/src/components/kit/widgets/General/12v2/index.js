import React from 'react'

const General12v2 = () => {
  return (
    <div>
      <div className="h-100 pt-5 pb-5 text-center">
        <i className="fe fe-users font-size-50" />
      </div>
      <div className="d-flex mb-1">
        <div className="text-uppercase font-weight-bold text-dark mr-auto">New users</div>
        <div>+20% Goal Reached</div>
      </div>
      <div className="d-flex mb-2">
        <div className="font-size-24 font-weight-bold text-success mr-auto">+3,125</div>
        <div className="font-size-24 text-gray-4">5,000</div>
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

export default General12v2
