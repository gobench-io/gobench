import React, { useState, useEffect } from 'react'
import { Layout, Breadcrumb } from 'antd'
import { Link } from 'react-router-dom'
import MenuLeft from './menu'
import { RiseOutlined } from '@ant-design/icons'
import { useInterval, INTERVAL } from '../../realtimeHelpers'
import GoBenchAPI from '../../api/gobench'
import { statusColors, iconStatus } from '../Status'
import { ApplicationsListContext, SpinnerContext } from '../../context'
import 'antd/dist/antd.css'
import './style.css'

const { Header, Content, Footer, Sider } = Layout

const MainLayout = (props) => {
  const [collapse, setCollapse] = useState(false)
  const [app, setApp] = useState({})
  const [fetching, setIsFetching] = useState(true)

  const cancelRunApplication = (id) => {
    GoBenchAPI.cancelApplication(id).then(() => {
      GoBenchAPI.getApplications().then(apps => {
        setApp({ ...app, apps })
      })
    })
  }

  useEffect(() => {
    if (!app.cancelRunApplication) {
      setApp({ ...app, cancelRunApplication })
    }
    if (!app.apps) {
      GoBenchAPI.getApplications().then(apps => {
        setApp({ ...app, apps })
        setIsFetching(false)
      })
    }
  }, [app, cancelRunApplication])

  useInterval(() => {
    if (app.apps && app.apps.length > 0) {
      GoBenchAPI.getApplications().then(apps => {
        setApp({ ...app, apps })
      })
    }
  }, INTERVAL)

  return (
    <Layout style={{ minHeight: '100vh' }} className='benchmark-layout'>
      <ApplicationsListContext.Provider value={app}>
        <SpinnerContext.Provider value={fetching}>
          <Sider collapsible collapsed={collapse} onCollapse={() => setCollapse(!collapse)}>
            <h2 className='logo'>
              {collapse
                ? <RiseOutlined style={{ color: '#1890ff' }} />
                : <Link to='/'> <img width='125' src='/resources/gobench-logo.png' /></Link>}
            </h2>
            <MenuLeft
              className='benchmark-menu'
              theme='light'
              mode='inline'
              defaultSelected={['1']}
            />

          </Sider>
          <Layout className='site-layout'>
            <Header className='site-layout-background' style={{ padding: 0, textAlign: 'center' }}>
              <Link to='/'><img width='150' src='resources/gobench-logo-full.png' /></Link>
            </Header>
            <Content style={{ margin: '0 16px' }}>
              <Breadcrumb style={{ margin: '16px 0' }}>
                <Breadcrumb.Item>Applications</Breadcrumb.Item>
              </Breadcrumb>
              <div className='site-layout-background' style={{ padding: 24, minHeight: 360 }}>
                {props.children}
              </div>
            </Content>
            <Footer style={{ textAlign: 'center' }}>Gobench ©2020 Created by
              <Link className='gobench-sponsor' to='https://github.com/gobench-io/gobench'>
                <img width='120' src='https://camo.githubusercontent.com/974d4b314bb0c8293c13a778dc0d72bc3ad7abf4/68747470733a2f2f766572696b2d7374617469632e73332d75732d776573742d322e616d617a6f6e6177732e636f6d2f6c6f676f2f766572696b5f6c6f676f2e737667' />
              </Link>
            </Footer>
          </Layout>
        </SpinnerContext.Provider>
      </ApplicationsListContext.Provider>
    </Layout>
  )
}
export default MainLayout
