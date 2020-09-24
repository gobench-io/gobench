/* eslint-disable */
import React from 'react'
import { Calendar } from 'antd'

class AntdCalendarExample extends React.Component {
  onPanelChange = (value, mode) => {
    console.log(value, mode)
  }

  render() {
    return (
      <div>
        <Calendar onPanelChange={this.onPanelChange} />
      </div>
    )
  }
}

export default AntdCalendarExample
