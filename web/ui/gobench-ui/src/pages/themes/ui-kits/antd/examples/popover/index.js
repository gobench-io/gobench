/* eslint-disable */
import React from 'react'
import { Popover, Button } from 'antd'

const content = (
  <div>
    <p>Content</p>
    <p>Content</p>
  </div>
)

class AntdPopoverExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Popover content={content} title="Title" className="mr-3 mb-3">
            <Button type="primary">Default Popover</Button>
          </Popover>
          <Popover
            placement="topLeft"
            title="Title"
            content={content}
            arrowPointAtCenter
            className="mr-3 mb-3"
          >
            <Button>Arrow points to center</Button>
          </Popover>
        </div>
      </div>
    )
  }
}

export default AntdPopoverExample
