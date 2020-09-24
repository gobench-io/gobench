/* eslint-disable */
import React from 'react'
import { DownOutlined } from '@ant-design/icons'
import { Menu, Dropdown, Button } from 'antd'
const menu = (
  <Menu>
    <Menu.Item>
      <a target="_blank" rel="noopener noreferrer" href="http://www.alipay.com/">
        1st menu item
      </a>
    </Menu.Item>
    <Menu.Item>
      <a target="_blank" rel="noopener noreferrer" href="http://www.taobao.com/">
        2nd menu item
      </a>
    </Menu.Item>
    <Menu.Item>
      <a target="_blank" rel="noopener noreferrer" href="http://www.tmall.com/">
        3rd menu item
      </a>
    </Menu.Item>
  </Menu>
)

class AntdDropdownExample extends React.Component {
  render() {
    return (
      <div>
        <div className="row">
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Basic</strong>
            </h5>
            <div className="mb-5">
              <Dropdown overlay={menu}>
                <a className="ant-dropdown-link">
                  Hover me <DownOutlined />
                </a>
              </Dropdown>
            </div>
            <h5 className="mb-3">
              <strong>Placement</strong>
            </h5>
            <div className="mb-5">
              <Dropdown overlay={menu} placement="bottomLeft" className="mr-3 mb-3">
                <Button>bottomLeft</Button>
              </Dropdown>
              <Dropdown overlay={menu} placement="bottomCenter" className="mr-3 mb-3">
                <Button>bottomCenter</Button>
              </Dropdown>
              <Dropdown overlay={menu} placement="bottomRight" className="mr-3 mb-3">
                <Button>bottomRight</Button>
              </Dropdown>
              <br />
              <Dropdown overlay={menu} placement="topLeft" className="mr-3 mb-3">
                <Button>topLeft</Button>
              </Dropdown>
              <Dropdown overlay={menu} placement="topCenter" className="mr-3 mb-3">
                <Button>topCenter</Button>
              </Dropdown>
              <Dropdown overlay={menu} placement="topRight" className="mr-3 mb-3">
                <Button>topRight</Button>
              </Dropdown>
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Context Menu</strong>
            </h5>
            <div className="mb-5">
              <Dropdown overlay={menu} trigger={['contextMenu']}>
                <div
                  className="bg-light"
                  style={{
                    textAlign: 'center',
                    height: 200,
                    lineHeight: '200px',
                  }}
                >
                  Right Click on Me
                </div>
              </Dropdown>
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default AntdDropdownExample
