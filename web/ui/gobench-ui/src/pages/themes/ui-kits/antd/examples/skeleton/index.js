/* eslint-disable */
import React from 'react'
import { Skeleton } from 'antd'

class AntdSkeletonExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Skeleton active />
        </div>
      </div>
    )
  }
}

export default AntdSkeletonExample
