/* eslint-disable */
import React from 'react'
import { Pagination } from 'antd'

class AntdPaginationExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Pagination Usage</strong>
        </h5>
        <div className="mb-5">
          <Pagination defaultCurrent={1} total={50} />
        </div>
        <div className="mb-5">
          <Pagination showSizeChanger defaultCurrent={3} total={500} />
        </div>
        <div className="mb-5">
          <Pagination total={50} showSizeChanger showQuickJumper />
        </div>
        <div className="mb-5">
          <Pagination simple defaultCurrent={2} total={50} />
        </div>
        <div className="mb-5">
          <Pagination size="small" total={50} />
        </div>
        <div className="mb-5">
          <Pagination size="small" total={50} showSizeChanger showQuickJumper />
        </div>
        <div className="mb-5">
          <Pagination size="small" total={50} />
        </div>
      </div>
    )
  }
}

export default AntdPaginationExample
