/* eslint-disable */
import React from 'react'
import { Switch } from 'antd'

function onChange(checked) {
  console.log(`switch to ${checked}`)
}

class AntdSwitchExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Switch className="component-col" defaultChecked className="mb-3 mr-3" />
          <Switch
            className="component-col"
            checkedChildren="Off"
            unCheckedChildren="On"
            defaultChecked
            className="mb-3 mr-3"
          />
          <Switch className="component-col" loading defaultChecked className="mb-3 mr-3" />
        </div>
      </div>
    )
  }
}

export default AntdSwitchExample
