/* eslint-disable */
import React from 'react'
import { Input } from 'antd'

const { Search, TextArea } = Input

class AntdInputExample extends React.Component {
  render() {
    return (
      <div>
        <div className="row">
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Basic</strong>
            </h5>
            <div className="mb-5">
              <Input placeholder="Basic usage" />
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Search</strong>
            </h5>
            <div className="mb-5">
              <Search placeholder="input search text" enterButton />
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Textarea</strong>
            </h5>
            <div className="mb-5">
              <TextArea rows={4} />
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default AntdInputExample
