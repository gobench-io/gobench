/* eslint-disable */
import React from 'react'
import { Spin, Alert } from 'antd'

class AntdSpinExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Spin tip="Loading...">
            <Alert
              message="Alert message title"
              description="Further details about the context of this alert."
              type="info"
            />
          </Spin>
        </div>
      </div>
    )
  }
}

export default AntdSpinExample
