/* eslint-disable */
import React from 'react'
import { Progress } from 'antd'

class AntdProgressExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <div className="mb-3">
            <Progress percent={30} />
          </div>
          <div className="mb-3">
            <Progress percent={50} status="active" />
          </div>
          <div className="mb-3">
            <Progress percent={70} status="exception" />
          </div>
          <div className="mb-3">
            <Progress percent={100} />
          </div>
          <div className="mb-3">
            <Progress percent={50} showInfo={false} />
          </div>
        </div>
      </div>
    )
  }
}

export default AntdProgressExample
