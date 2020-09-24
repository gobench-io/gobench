/* eslint-disable */
import React from 'react'
import {
  HomeOutlined,
  LoadingOutlined,
  SettingFilled,
  SmileOutlined,
  SyncOutlined,
} from '@ant-design/icons'

class AntdIconExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Icons Usage</strong>
        </h5>
        <div className="mb-5">
          <HomeOutlined className="mr-3 mb-3 font-size-24" />
          <SettingFilled className="mr-3 mb-3 font-size-24" />
          <SmileOutlined className="mr-3 mb-3 font-size-24" />
          <SyncOutlined spin className="mr-3 mb-3 font-size-24" />
          <SmileOutlined rotate={180} className="mr-3 mb-3 font-size-24" />
          <LoadingOutlined className="mr-3 mb-3 font-size-24" />
        </div>
      </div>
    )
  }
}

export default AntdIconExample
