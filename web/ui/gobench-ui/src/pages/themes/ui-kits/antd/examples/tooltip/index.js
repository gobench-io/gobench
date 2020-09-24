/* eslint-disable */
import React from 'react'
import { Tooltip, Button } from 'antd'

class AntdTooltipExample extends React.Component {
  render() {
    return (
      <div>
        <div className="row">
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Basic</strong>
            </h5>
            <div className="mb-5">
              <Tooltip placement="topLeft" title="Prompt Text">
                <Button>Align edge</Button>
              </Tooltip>
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Aligned</strong>
            </h5>
            <div className="mb-5">
              <Tooltip placement="topLeft" title="Prompt Text" arrowPointAtCenter>
                <Button>Arrow points to center</Button>
              </Tooltip>
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default AntdTooltipExample
