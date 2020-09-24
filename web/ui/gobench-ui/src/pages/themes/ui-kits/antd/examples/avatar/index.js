/* eslint-disable */
import React from 'react'
import { UserOutlined } from '@ant-design/icons'
import { Avatar } from 'antd'

class AntdAvatarExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <div className="d-inline-block mr-4">
            <Avatar icon={<UserOutlined />} />
          </div>
          <div className="d-inline-block mr-4">
            <Avatar shape="square" icon={<UserOutlined />} />
          </div>
        </div>
        <h5 className="mb-3">
          <strong>Type</strong>
        </h5>
        <div className="mb-5">
          <div className="d-inline-block mr-4">
            <Avatar icon={<UserOutlined />} />
          </div>
          <div className="d-inline-block mr-4">
            <Avatar>U</Avatar>
          </div>
          <div className="d-inline-block mr-4">
            <Avatar>USER</Avatar>
          </div>
          <div className="d-inline-block mr-4">
            <Avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
          </div>
          <div className="d-inline-block mr-4">
            <Avatar style={{ color: '#f56a00', backgroundColor: '#fde3cf' }}>U</Avatar>
          </div>
          <div className="d-inline-block mr-4">
            <Avatar style={{ backgroundColor: '#87d068' }} icon={<UserOutlined />} />
          </div>
        </div>
      </div>
    )
  }
}

export default AntdAvatarExample
