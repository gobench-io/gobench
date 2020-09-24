/* eslint-disable */
import React from 'react'
import { InputNumber } from 'antd'

function onChange(value) {
  console.log('changed', value)
}

class AntdInputNumberExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <InputNumber min={1} max={10} defaultValue={3} onChange={onChange} />
        </div>
      </div>
    )
  }
}

export default AntdInputNumberExample
