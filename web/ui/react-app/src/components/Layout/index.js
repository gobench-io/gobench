import React, { useState } from 'react'
import { Layout, Menu, Breadcrumb } from 'antd'
import { Link } from 'react-router-dom'
import MenuLeft from './menu'

import 'antd/dist/antd.css'
import './style.css'

const { Header, Content, Footer, Sider } = Layout
const { SubMenu } = Menu

const MainLayout = (props) => {
  const [collapse, setCollapse] = useState(false)

  return (
    <Layout style={{ minHeight: '100vh' }} className='benchmark-layout'>
      <Sider collapsible collapsed={collapse} onCollapse={() => setCollapse(!collapse)}>
        <h2 className='logo'>
    Gobench

        </h2>
        <MenuLeft theme='dark' mode='inline' defaultSelected={['1']} />

      </Sider>
      <Layout className='site-layout'>
        <Header className='site-layout-background' style={{ padding: 0 }} />
        <Content style={{ margin: '0 16px' }}>
          <Breadcrumb style={{ margin: '16px 0' }}>
            <Breadcrumb.Item>Applications</Breadcrumb.Item>
            <Breadcrumb.Item>Bill</Breadcrumb.Item>
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
    </Layout>
  )
}
export default MainLayout
