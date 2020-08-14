import React, { useEffect, useState, lazy, Suspense } from 'react'
import { get } from 'lodash'
import GoBenchAPI from '../../api/gobench'
import { Row, Col } from 'antd'
import { CaretDownOutlined, CaretUpOutlined } from '@ant-design/icons'

const GraphComponent = lazy(() => import('./Graph'))

const loading = () => <p>Loading group...</p>

function Group({ group, timestamp, expandDefault = false }) {
  const [graphs, fetchGraphs] = useState([])
  const [isCollapse, toggleCollapse] = useState(!expandDefault)
  useEffect(() => {
    if (group && group.id && !isCollapse) {
      GoBenchAPI.getGraphs(group.id).then(res => {
        return fetchGraphs(res)
      })
    }
  }, [group, isCollapse])
  return (
    <div className="group">
      <div className="group-header clickable"
        onClick={() => toggleCollapse(!isCollapse)}>
        <h3 title={graphs.id || ''} className="group-title">{get(group, 'name', '')}</h3>
        <span className="collapse-button">
          {isCollapse ? <CaretDownOutlined /> : <CaretUpOutlined />}
        </span>
      </div>
      <div className={`group-graphs ${isCollapse ? 'collapse' : ''}`}>
        {
          !isCollapse && <Suspense fallback={loading()}>
            <Row gutter={[16, 16]}>
              {
                graphs.length > 0 ?
                  graphs.map((graph, index) => {
                    return <Col key={graph.id || index} xs={24} sm={24} md={24} lg={12} xl={8}>
                      <GraphComponent graph={graph} timestamp={timestamp} />
                    </Col>
                  })
                  : <p className="text-center">Cannot load graphs.</p>
              }
            </Row>
          </Suspense>
        }
      </div>
    </div>
  )
}

export default Group
