/* eslint-disable */
import React from 'react'
import { Slider, Switch } from 'antd'

class AntdSliderExample extends React.Component {
  state = {
    disabled: false,
  }

  handleDisabledChange = disabled => {
    this.setState({ disabled })
  }

  render() {
    const { disabled } = this.state
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Slider defaultValue={30} disabled={disabled} />
          <Slider range defaultValue={[20, 50]} disabled={disabled} />
          <br />
          <Switch
            unCheckedChildren="enabled"
            checkedChildren="disabled"
            checked={disabled}
            onChange={this.handleDisabledChange}
          />
        </div>
      </div>
    )
  }
}

export default AntdSliderExample
