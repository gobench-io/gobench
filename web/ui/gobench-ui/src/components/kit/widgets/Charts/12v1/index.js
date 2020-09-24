import React from 'react'
import { Progress } from 'antd'

const Chart12v1 = () => {
  return (
    <div>
      <div className="d-flex justify-content-between align-items-center">
        <div className="pr-3">
          <h2 className="font-size-18 font-weight-bold mb-1 text-dark">Feedbacks</h2>
          <p className="font-size-15 mb-3">Profit</p>
        </div>
        <div className="text-success font-weight-bold font-size-24">160,100</div>
      </div>
      <div className="mb-3">
        <Progress
          type="line"
          percent={55}
          showInfo={false}
          strokeWidth={12}
          strokeColor="#46be8a"
        />
      </div>
      <div className="d-flex text-gray-5 justify-content-between font-size-14">
        <span className="text-uppercase">Change</span>
        <span className="text-uppercase">55%</span>
      </div>
    </div>
  )
}

export default Chart12v1
