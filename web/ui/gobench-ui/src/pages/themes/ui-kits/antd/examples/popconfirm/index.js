/* eslint-disable */
import React from 'react'
import { Popconfirm, message, Button } from 'antd'

function confirm(e) {
  console.log(e)
  message.success('Click on Yes')
}

function cancel(e) {
  console.log(e)
  message.error('Click on No')
}

class AntdPopconfirmExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Popconfirm
            title="Are you sure delete this task?"
            onConfirm={confirm}
            onCancel={cancel}
            okText="Yes"
            cancelText="No"
          >
            <Button type="primary">Confirm Deletion</Button>
          </Popconfirm>
        </div>
      </div>
    )
  }
}

export default AntdPopconfirmExample
