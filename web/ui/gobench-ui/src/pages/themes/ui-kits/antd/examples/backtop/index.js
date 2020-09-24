/* eslint-disable */
import React from 'react'
import { BackTop } from 'antd'

class AntdBackTopExample extends React.Component {
  render() {
    return (
      <div className="height-700">
        <BackTop>
          <div className="ant-back-top-inner">UP</div>
        </BackTop>
        Scroll down to see the bottom-right
        <strong style={{ color: '#1088e9' }}> blue </strong>
        button.
        <style>
          {`
            .ant-back-top {
              bottom: 100px;
              z-index: 10000;
            }
            .ant-back-top-inner {
              height: 40px;
              width: 40px;
              line-height: 40px;
              border-radius: 4px;
              background-color: #1088e9;
              color: #fff;
              text-align: center;
              font-size: 20px;
            }
          `}
        </style>
      </div>
    )
  }
}

export default AntdBackTopExample
