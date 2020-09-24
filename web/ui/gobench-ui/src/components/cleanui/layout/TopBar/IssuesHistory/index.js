import React from 'react'
import { FormattedMessage } from 'react-intl'
import { Menu, Dropdown } from 'antd'
import { Link } from 'react-router-dom'
import styles from './style.module.scss'

const IssuesHistory = () => {
  const menu = (
    <Menu selectable={false}>
      <Menu.Item>
        <Link to="/">Current Search</Link>
      </Menu.Item>
      <Menu.Item>
        <Link to="/">Search for issues</Link>
      </Menu.Item>
      <Menu.Divider />
      <Menu.ItemGroup title="Opened">
        <Menu.Item>
          <Link to="/">
            <i className="fe fe-check-circle mr-2" /> CUI-125 Project Implemen...
          </Link>
        </Menu.Item>
        <Menu.Item>
          <Link to="/">
            <i className="fe fe-check-circle mr-2" /> CUI-147 Active History Is...
          </Link>
        </Menu.Item>
        <Menu.Item>
          <Link to="/">
            <i className="fe fe-check-circle mr-2" /> CUI-424 Ionicons Integrat...
          </Link>
        </Menu.Item>
        <Menu.Item>
          <Link to="/">More...</Link>
        </Menu.Item>
      </Menu.ItemGroup>
      <Menu.ItemGroup title="Filters">
        <Menu.Item>
          <Link to="/">My Open Issues</Link>
        </Menu.Item>
        <Menu.Item>
          <Link to="/">Reported by Me</Link>
        </Menu.Item>
      </Menu.ItemGroup>
      <Menu.Divider />
      <Menu.Item>
        <Link to="/">
          <i className="fe fe-settings mr-2" /> Settings
        </Link>
      </Menu.Item>
    </Menu>
  )
  return (
    <Dropdown overlay={menu} trigger={['click']} placement="bottomLeft">
      <div className={styles.dropdown}>
        <i className={`${styles.icon} fe fe-folder`} />
        <span className="d-none d-xl-inline">
          <FormattedMessage id="topBar.issuesHistory" />
        </span>
      </div>
    </Dropdown>
  )
}

export default IssuesHistory
