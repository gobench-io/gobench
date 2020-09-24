/* eslint-disable */
import React from 'react'
import { Tag } from 'antd'

function log(e) {
  console.log(e)
}

function preventDefault(e) {
  e.preventDefault()
  console.log('Clicked! But prevent default.')
}

class AntdTagExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-3">
          <Tag className="mb-3 mr-3">Tag 1</Tag>
          <Tag className="mb-3 mr-3">
            <a href="https://github.com/ant-design/ant-design/issues/1862">Link</a>
          </Tag>
          <Tag className="mb-3 mr-3" closable>
            Tag 2
          </Tag>
          <Tag className="mb-3 mr-3" closable>
            Prevent Default
          </Tag>
        </div>
        <h5 className="mb-3">
          <strong>Presets</strong>
        </h5>
        <div className="mb-3">
          <Tag className="mb-3 mr-3" color="magenta">
            magenta
          </Tag>
          <Tag className="mb-3 mr-3" color="red">
            red
          </Tag>
          <Tag className="mb-3 mr-3" color="volcano">
            volcano
          </Tag>
          <Tag className="mb-3 mr-3" color="orange">
            orange
          </Tag>
          <Tag className="mb-3 mr-3" color="gold">
            gold
          </Tag>
          <Tag className="mb-3 mr-3" color="lime">
            lime
          </Tag>
          <Tag className="mb-3 mr-3" color="green">
            green
          </Tag>
          <Tag className="mb-3 mr-3" color="cyan">
            cyan
          </Tag>
        </div>
        <h5 className="mb-3">
          <strong>Custom</strong>
        </h5>
        <div className="mb-3">
          <Tag color="#f50">#f50</Tag>
          <Tag color="#2db7f5">#2db7f5</Tag>
          <Tag color="#87d068">#87d068</Tag>
          <Tag color="#108ee9">#108ee9</Tag>
        </div>
      </div>
    )
  }
}

export default AntdTagExample
