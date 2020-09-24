/* eslint-disable */
import React from 'react'
import { EditOutlined, EllipsisOutlined, SettingOutlined } from '@ant-design/icons'
import { Card, Avatar } from 'antd'

const { Meta } = Card

class AntdCardExample extends React.Component {
  render() {
    return (
      <div>
        <div className="row">
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Basic</strong>
            </h5>
            <div className="mb-5">
              <div className="mb-3">
                <Card title="Default size card" extra={<a href="#">More</a>} style={{ width: 300 }}>
                  <p>Card content</p>
                  <p>Card content</p>
                  <p>Card content</p>
                </Card>
              </div>
              <div className="mb-3">
                <Card
                  size="small"
                  title="Small size card"
                  extra={<a href="#">More</a>}
                  style={{ width: 300 }}
                >
                  <p>Card content</p>
                  <p>Card content</p>
                  <p>Card content</p>
                </Card>
              </div>
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Advanced</strong>
            </h5>
            <div className="mb-5">
              <div className="mb-3">
                <Card
                  style={{ width: 300 }}
                  cover={
                    <img
                      alt="example"
                      src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
                    />
                  }
                  actions={[
                    <SettingOutlined key="setting" />,
                    <EditOutlined key="edit" />,
                    <EllipsisOutlined key="ellipsis" />,
                  ]}
                >
                  <Meta
                    avatar={
                      <Avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
                    }
                    title="Card title"
                    description="This is the description"
                  />
                </Card>
              </div>
              <div className="mb-3">
                <Card style={{ width: 300, marginTop: 16 }} loading>
                  <Meta
                    avatar={
                      <Avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
                    }
                    title="Card title"
                    description="This is the description"
                  />
                </Card>
              </div>
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default AntdCardExample
