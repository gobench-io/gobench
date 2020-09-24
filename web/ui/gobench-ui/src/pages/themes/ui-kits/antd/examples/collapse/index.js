/* eslint-disable */
import React from 'react'
import { Collapse } from 'antd'

const { Panel } = Collapse

function callback(key) {
  console.log(key)
}

const text = `
  A dog is a type of domesticated animal.
  Known for its loyalty and faithfulness,
  it can be found as a welcome guest in many households across the world.
`

class AntdCollapseExample extends React.Component {
  render() {
    return (
      <div>
        <div className="mb-5">
          <h5 className="mb-3">
            <strong>Basic</strong>
          </h5>
          <Collapse defaultActiveKey={['1']} onChange={callback}>
            <Panel header="This is panel header 1" key="1">
              <p>{text}</p>
            </Panel>
            <Panel header="This is panel header 2" key="2">
              <p>{text}</p>
            </Panel>
            <Panel header="This is panel header 3" key="3" disabled>
              <p>{text}</p>
            </Panel>
          </Collapse>
        </div>
        <div className="mb-5">
          <h5 className="mb-3">
            <strong>Borderless</strong>
          </h5>
          <Collapse defaultActiveKey={['1']} onChange={callback} bordered={false}>
            <Panel header="This is panel header 1" key="1">
              <p>{text}</p>
            </Panel>
            <Panel header="This is panel header 2" key="2">
              <p>{text}</p>
            </Panel>
            <Panel header="This is panel header 3" key="3" disabled>
              <p>{text}</p>
            </Panel>
          </Collapse>
        </div>
      </div>
    )
  }
}

export default AntdCollapseExample
