/* eslint-disable */
import React from 'react'
import { Rate } from 'antd'

class AntdRateExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-3">
          <Rate defaultValue={3} />
        </div>
        <div className="mb-3">
          <span>
            <Rate allowHalf defaultValue={2.5} tooltips="good" />
            <span className="ant-rate-text">good</span>
          </span>
        </div>
        <div className="mb-3">
          <Rate defaultValue={3} character="W" />
        </div>
      </div>
    )
  }
}

export default AntdRateExample
