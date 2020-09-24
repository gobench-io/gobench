/* eslint-disable */
import React from 'react'
import { Row, Col } from 'antd'

class AntdGridExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <Row className="mb-3">
          <Col span={12}>
            <div className="bg-light" style={{ padding: '10px' }}>
              col-12
            </div>
          </Col>
          <Col span={12}>
            <div className="bg-light" style={{ padding: '10px' }}>
              col-12
            </div>
          </Col>
        </Row>
        <Row className="mb-3">
          <Col span={8}>
            <div className="bg-light" style={{ padding: '10px' }}>
              col-8
            </div>
          </Col>
          <Col span={8}>
            <div className="bg-light" style={{ padding: '10px' }}>
              col-8
            </div>
          </Col>
          <Col span={8}>
            <div className="bg-light" style={{ padding: '10px' }}>
              col-8
            </div>
          </Col>
        </Row>
        <Row className="mb-3">
          <Col span={6}>
            <div className="bg-light" style={{ padding: '10px' }}>
              col-6
            </div>
          </Col>
          <Col span={6}>
            <div className="bg-light" style={{ padding: '10px' }}>
              col-6
            </div>
          </Col>
          <Col span={6}>
            <div className="bg-light" style={{ padding: '10px' }}>
              col-6
            </div>
          </Col>
          <Col span={6}>
            <div className="bg-light" style={{ padding: '10px' }}>
              col-6
            </div>
          </Col>
        </Row>
      </div>
    )
  }
}

export default AntdGridExample
