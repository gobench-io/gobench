/* eslint-disable */
import React from 'react'
import { Radio } from 'antd'

class AntdRadioExample extends React.Component {
  render() {
    return (
      <div>
        <div className="row">
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Basic</strong>
            </h5>
            <div className="mb-5">
              <Radio.Group defaultValue={1}>
                <Radio value={1}>Apple</Radio>
                <Radio value={2}>Pear</Radio>
                <Radio value={3}>Banana</Radio>
                <Radio value={4}>Strawberry</Radio>
              </Radio.Group>
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Buttons</strong>
            </h5>
            <div className="mb-5">
              <Radio.Group defaultValue="a">
                <Radio.Button value="a">Hangzhou</Radio.Button>
                <Radio.Button value="b">Shanghai</Radio.Button>
                <Radio.Button value="c">Beijing</Radio.Button>
                <Radio.Button value="d">Chengdu</Radio.Button>
              </Radio.Group>
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default AntdRadioExample
