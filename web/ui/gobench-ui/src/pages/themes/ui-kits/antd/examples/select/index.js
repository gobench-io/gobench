/* eslint-disable */
import React from 'react'
import { Select } from 'antd'

const { Option, OptGroup } = Select

const children = []
for (let i = 10; i < 36; i++) {
  children.push(<Option key={i.toString(36) + i}>{i.toString(36) + i}</Option>)
}

class AntdSelectExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <div className="d-inline-block mb-3 mr-3">
            <Select defaultValue="lucy" style={{ width: 160 }}>
              <Option value="jack">Jack</Option>
              <Option value="lucy">Lucy</Option>
              <Option value="disabled" disabled>
                Disabled
              </Option>
              <Option value="Yiminghe">yiminghe</Option>
            </Select>
          </div>
          <div className="d-inline-block mb-3 mr-3">
            <Select defaultValue="lucy" style={{ width: 160 }}>
              <OptGroup label="Manager">
                <Option value="jack">Jack</Option>
                <Option value="lucy">Lucy</Option>
              </OptGroup>
              <OptGroup label="Engineer">
                <Option value="Yiminghe">yiminghe</Option>
              </OptGroup>
            </Select>
          </div>
        </div>
        <h5 className="mb-3">
          <strong>Multiple</strong>
        </h5>
        <div className="mb-5">
          <Select
            mode="multiple"
            style={{ width: '100%' }}
            placeholder="Please select"
            defaultValue={['a10', 'c12']}
          >
            {children}
          </Select>
        </div>
      </div>
    )
  }
}

export default AntdSelectExample
