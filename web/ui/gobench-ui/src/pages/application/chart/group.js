import React, { lazy, useEffect, useState } from 'react'
import { connect } from 'react-redux'
import { CaretDownOutlined, CaretUpOutlined } from '@ant-design/icons'
import { Row, Col } from 'antd'
import { withRouter } from 'react-router-dom'
import { get } from 'lodash'

const Graph = lazy(() => import('./graph'))

const mapStateToProps = ({ application, dispatch }) => {
  const { detail, graphs } = application
  return {
    detail,
    graphs,
    dispatch
  }
}
const DefaultPage = ({ group, graphs = [], timestamp, expandDefault = false, dispatch }) => {
  const [collapsed, toggleCollapse] = useState(!expandDefault)
  const [_graphs, setGraphs] = useState([])
  const ag = graphs.some(x => x.groupId === group.id)
  useEffect(() => {
    if (group && !collapsed) {
      if (graphs.every(x => x.groupId !== group.id)) {
        dispatch({
          type: 'application/GRAPHS',
          payload: { id: group.id }
        })
      }
    }
  }, [group, collapsed])
  useEffect(() => {
    if (ag) {
      setGraphs(graphs.filter(x => x.groupId === group.id))
    }
  }, [graphs])
  return (
    <>
      <div className='application-group'>
        <div className='group'>
          <div
            className='group-header clickable'
            onClick={() => toggleCollapse(!collapsed)}
          >
            <h3 title={_graphs.id || ''} className='group-title'>{get(group, 'name', '')}</h3>
            <span className='collapse-button'>
              {collapsed ? <CaretDownOutlined /> : <CaretUpOutlined />}
            </span>
          </div>
          <div className={`group-graphs ${collapsed ? 'collapse' : ''}`}>
            {
              !collapsed &&
                <Row gutter={[16, 16]}>
                  {
                    _graphs.length > 0
                      ? _graphs.map((graph, index) => {
                        return (
                          <Col key={graph.id || index} xs={24} sm={24} md={24} lg={12} xl={8}>
                            <Graph graph={graph} timestamp={timestamp} />
                          </Col>
                        )
                      })
                      : <p className='text-center'>Cannot load graphs.</p>
                  }
                </Row>
            }
          </div>
        </div>
      </div>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))
