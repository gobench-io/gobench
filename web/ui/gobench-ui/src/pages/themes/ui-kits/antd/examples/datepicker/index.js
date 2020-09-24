/* eslint-disable */
import React from 'react'
import { DatePicker } from 'antd'

const { MonthPicker, RangePicker, WeekPicker } = DatePicker

function onChange(date, dateString) {
  console.log(date, dateString)
}

class AntdDatePickerExample extends React.Component {
  render() {
    return (
      <div>
        <div className="row">
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>DatePicker</strong>
            </h5>
            <div className="mb-5">
              <DatePicker onChange={onChange} className="mb-2" />
            </div>
            <h5 className="mb-3">
              <strong>Month Picker</strong>
            </h5>
            <div className="mb-5">
              <MonthPicker onChange={onChange} className="mb-2" placeholder="Select month" />
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Range Picker</strong>
            </h5>
            <div className="mb-5">
              <RangePicker onChange={onChange} className="mb-2" />
            </div>
            <h5 className="mb-3">
              <strong>Week Picker</strong>
            </h5>
            <div className="mb-5">
              <WeekPicker onChange={onChange} className="mb-2" placeholder="Select week" />
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default AntdDatePickerExample
