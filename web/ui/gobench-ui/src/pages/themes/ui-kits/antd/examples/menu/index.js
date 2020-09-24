/* eslint-disable */
import React from 'react'
import { AppstoreOutlined, MailOutlined, SettingOutlined } from '@ant-design/icons'
import { Menu } from 'antd'

const { SubMenu } = Menu

class AntdMenuExample extends React.Component {
  state = {
    current: 'mail',
  }

  handleClick = e => {
    console.log('click ', e)
    this.setState({
      current: e.key,
    })
  }

  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Horizontal</strong>
        </h5>
        <div className="mb-5">
          <Menu onClick={this.handleClick} selectedKeys={[this.state.current]} mode="horizontal">
            <Menu.Item key="mail">
              <MailOutlined />
              Navigation One
            </Menu.Item>
            <Menu.Item key="app" disabled>
              <AppstoreOutlined />
              Navigation Two
            </Menu.Item>
            <SubMenu
              title={
                <span className="submenu-title-wrapper">
                  <SettingOutlined />
                  Navigation Three - Submenu
                </span>
              }
            >
              <Menu.ItemGroup title="Item 1">
                <Menu.Item key="setting:1">Option 1</Menu.Item>
                <Menu.Item key="setting:2">Option 2</Menu.Item>
              </Menu.ItemGroup>
              <Menu.ItemGroup title="Item 2">
                <Menu.Item key="setting:3">Option 3</Menu.Item>
                <Menu.Item key="setting:4">Option 4</Menu.Item>
              </Menu.ItemGroup>
            </SubMenu>
            <Menu.Item key="alipay">
              <a href="https://ant.design" target="_blank" rel="noopener noreferrer">
                Navigation Four - Link
              </a>
            </Menu.Item>
          </Menu>
        </div>
        <h5 className="mb-3">
          <strong>Vertical</strong>
        </h5>
        <div className="mb-5">
          <Menu onClick={this.handleClick} style={{ width: 256 }} mode="vertical">
            <SubMenu
              key="sub1"
              title={
                <span>
                  <MailOutlined />
                  <span>Navigation One</span>
                </span>
              }
            >
              <Menu.ItemGroup title="Item 1">
                <Menu.Item key="1">Option 1</Menu.Item>
                <Menu.Item key="2">Option 2</Menu.Item>
              </Menu.ItemGroup>
              <Menu.ItemGroup title="Iteom 2">
                <Menu.Item key="3">Option 3</Menu.Item>
                <Menu.Item key="4">Option 4</Menu.Item>
              </Menu.ItemGroup>
            </SubMenu>
            <SubMenu
              key="sub2"
              title={
                <span>
                  <AppstoreOutlined />
                  <span>Navigation Two</span>
                </span>
              }
            >
              <Menu.Item key="5">Option 5</Menu.Item>
              <Menu.Item key="6">Option 6</Menu.Item>
              <SubMenu key="sub3" title="Submenu">
                <Menu.Item key="7">Option 7</Menu.Item>
                <Menu.Item key="8">Option 8</Menu.Item>
              </SubMenu>
            </SubMenu>
            <SubMenu
              key="sub4"
              title={
                <span>
                  <SettingOutlined />
                  <span>Navigation Three</span>
                </span>
              }
            >
              <Menu.Item key="9">Option 9</Menu.Item>
              <Menu.Item key="10">Option 10</Menu.Item>
              <Menu.Item key="11">Option 11</Menu.Item>
              <Menu.Item key="12">Option 12</Menu.Item>
            </SubMenu>
          </Menu>
        </div>
      </div>
    )
  }
}

export default AntdMenuExample
