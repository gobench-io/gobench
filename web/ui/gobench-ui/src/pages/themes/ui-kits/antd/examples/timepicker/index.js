/* eslint-disable */
import React from 'react'
import { TimePicker } from 'antd'
import moment from 'moment'

function onChange(time, timeString) {
  console.log(time, timeString)
}

class AntdTimePickerExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <TimePicker onChange={onChange} defaultOpenValue={moment('00:00:00', 'HH:mm:ss')} />
        </div>
      </div>
    )
  }
}

export default AntdTimePickerExample
