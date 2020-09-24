/* eslint-disable */
import React from 'react'
import { Layout } from 'antd'

const { Header, Footer, Sider, Content } = Layout

class AntdLayoutExample extends React.Component {
  render() {
    return (
      <div id="components-layout-demo-basic">
        <Layout>
          <Header>Header</Header>
          <Content>Content</Content>
          <Footer>Footer</Footer>
        </Layout>

        <Layout>
          <Header>Header</Header>
          <Layout>
            <Sider>Sider</Sider>
            <Content>Content</Content>
          </Layout>
          <Footer>Footer</Footer>
        </Layout>

        <Layout>
          <Header>Header</Header>
          <Layout>
            <Content>Content</Content>
            <Sider>Sider</Sider>
          </Layout>
          <Footer>Footer</Footer>
        </Layout>

        <Layout>
          <Sider>Sider</Sider>
          <Layout>
            <Header>Header</Header>
            <Content>Content</Content>
            <Footer>Footer</Footer>
          </Layout>
        </Layout>
        <style>
          {`
            #components-layout-demo-basic {
              text-align: center;
            }
            #components-layout-demo-basic .ant-layout-header,
            #components-layout-demo-basic .ant-layout-footer {
              background: #7dbcea;
              color: #fff;
            }
            #components-layout-demo-basic .ant-layout-footer {
              line-height: 1.5;
            }
            #components-layout-demo-basic .ant-layout-sider {
              background: #3ba0e9;
              color: #fff;
              line-height: 120px;
            }
            #components-layout-demo-basic .ant-layout-content {
              background: rgba(16, 142, 233, 1);
              color: #fff;
              min-height: 120px;
              line-height: 120px;
            }
            #components-layout-demo-basic > .ant-layout {
              margin-bottom: 48px;
            }
            #components-layout-demo-basic > .ant-layout:last-child {
              margin: 0;
            }
          `}
        </style>
      </div>
    )
  }
}

export default AntdLayoutExample
