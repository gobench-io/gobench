/* eslint-disable */
import React from 'react'
import { Mentions } from 'antd'

const { Option } = Mentions

function onChange(value) {
  console.log('Change:', value)
}

function onSelect(option) {
  console.log('select', option)
}

class AntdMentionsExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <Mentions
          style={{ width: '100%' }}
          onChange={onChange}
          onSelect={onSelect}
          defaultValue="@afc163"
        >
          <Option value="afc163">afc163</Option>
          <Option value="zombieJ">zombieJ</Option>
          <Option value="yesmeck">yesmeck</Option>
        </Mentions>
      </div>
    )
  }
}

export default AntdMentionsExample
