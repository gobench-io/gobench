/* eslint-disable */
import React from 'react'
import { HomeOutlined, UserOutlined } from '@ant-design/icons'
import { Breadcrumb } from 'antd'

class AntdBreadcrumbExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Breadcrumb>
            <Breadcrumb.Item>Home</Breadcrumb.Item>
            <Breadcrumb.Item>
              <a href="">Application Center</a>
            </Breadcrumb.Item>
            <Breadcrumb.Item>
              <a href="">Application List</a>
            </Breadcrumb.Item>
            <Breadcrumb.Item>An Application</Breadcrumb.Item>
          </Breadcrumb>
        </div>
        <h5 className="mb-3">
          <strong>With an Icon</strong>
        </h5>
        <div className="mb-5">
          <Breadcrumb>
            <Breadcrumb.Item href="">
              <HomeOutlined />
            </Breadcrumb.Item>
            <Breadcrumb.Item href="">
              <UserOutlined />
              <span>Application List</span>
            </Breadcrumb.Item>
            <Breadcrumb.Item>Application</Breadcrumb.Item>
          </Breadcrumb>
        </div>
      </div>
    )
  }
}

export default AntdBreadcrumbExample
