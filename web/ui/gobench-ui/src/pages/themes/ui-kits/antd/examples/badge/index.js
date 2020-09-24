/* eslint-disable */
import React from 'react'
import { ClockCircleOutlined } from '@ant-design/icons'
import { Badge } from 'antd'

class AntdBadgeExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <div className="d-inline-block mr-4">
            <Badge count={5}>
              <div
                style={{
                  width: '30px',
                  height: '30px',
                  borderRadius: '3px',
                  border: '1px solid rgba(150, 150, 150, .2)',
                }}
              />
            </Badge>
          </div>
          <div className="d-inline-block mr-4">
            <Badge count={0} showZero>
              <div
                style={{
                  width: '30px',
                  height: '30px',
                  borderRadius: '3px',
                  border: '1px solid rgba(150, 150, 150, .2)',
                }}
              />
            </Badge>
          </div>
          <div className="d-inline-block mr-4">
            <Badge count={<ClockCircleOutlined style={{ color: '#f5222d' }} />}>
              <div
                style={{
                  width: '30px',
                  height: '30px',
                  borderRadius: '3px',
                  border: '1px solid rgba(150, 150, 150, .2)',
                }}
              />
            </Badge>
          </div>
          <div className="d-inline-block mr-4">
            <Badge count={99}>
              <div
                style={{
                  width: '30px',
                  height: '30px',
                  borderRadius: '3px',
                  border: '1px solid rgba(150, 150, 150, .2)',
                }}
              />
            </Badge>
          </div>
          <div className="d-inline-block mr-4">
            <Badge count={100}>
              <div
                style={{
                  width: '30px',
                  height: '30px',
                  borderRadius: '3px',
                  border: '1px solid rgba(150, 150, 150, .2)',
                }}
              />
            </Badge>
          </div>
          <div className="d-inline-block mr-4">
            <Badge count={99} overflowCount={10}>
              <div
                style={{
                  width: '30px',
                  height: '30px',
                  borderRadius: '3px',
                  border: '1px solid rgba(150, 150, 150, .2)',
                }}
              />
            </Badge>
          </div>
          <div className="d-inline-block mr-4">
            <Badge count={1000} overflowCount={999}>
              <div
                style={{
                  width: '30px',
                  height: '30px',
                  borderRadius: '3px',
                  border: '1px solid rgba(150, 150, 150, .2)',
                }}
              />
            </Badge>
          </div>
          <div className="d-inline-block mr-4">
            <Badge dot>
              <div
                style={{
                  width: '30px',
                  height: '30px',
                  borderRadius: '3px',
                  border: '1px solid rgba(150, 150, 150, .2)',
                }}
              />
            </Badge>
          </div>
        </div>
        <h5 className="mb-3">
          <strong>Standalone</strong>
        </h5>
        <div className="mb-5">
          <div className="d-inline-block mr-4">
            <Badge count={25} />
          </div>
          <div className="d-inline-block mr-4">
            <Badge
              count={4}
              style={{
                backgroundColor: '#fff',
                color: '#999',
                boxShadow: '0 0 0 1px #d9d9d9 inset',
              }}
            />
          </div>
          <div className="d-inline-block mr-4">
            <Badge count={109} style={{ backgroundColor: '#52c41a' }} />
          </div>
        </div>
        <h5 className="mb-3">
          <strong>Badge</strong>
        </h5>
        <div className="mb-5">
          <div className="d-inline-block mr-4">
            <Badge status="success" text="Success" />
          </div>
          <div className="d-inline-block mr-4">
            <Badge status="error" text="Error" />
          </div>
          <div className="d-inline-block mr-4">
            <Badge status="default" text="Default" />
          </div>
          <div className="d-inline-block mr-4">
            <Badge status="processing" text="Processing" />
          </div>
          <div className="d-inline-block mr-4">
            <Badge status="warning" text="Warning" />
          </div>
        </div>
      </div>
    )
  }
}

export default AntdBadgeExample
