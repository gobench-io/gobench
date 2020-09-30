import React, { useState } from 'react'
import { FormattedMessage } from 'react-intl'
import { connect } from 'react-redux'
import { UserOutlined } from '@ant-design/icons'
import { Menu, Dropdown, Avatar } from 'antd'
import styles from './style.module.scss'

const mapStateToProps = ({ user }) => ({ user })

const ProfileMenu = ({ dispatch }) => {
  const [count, setCount] = useState(7)

  const logout = e => {
    e.preventDefault()
    dispatch({
      type: 'user/LOGOUT'
    })
  }

  const addCount = () => {
    setCount(count + 1)
  }

  const menu = (
    <Menu selectable={false}>
      <Menu.Item>
        <strong>
          <FormattedMessage id='administrator' />
        </strong>
        {/* <div>
          <strong className="mr-1">
            <FormattedMessage id="topBar.profileMenu.billingPlan" />:{' '}
          </strong>
          Professional
        </div> */}
        {/* <div>
          <strong>
            <FormattedMessage id="topBar.profileMenu.role" />:{' '}
          </strong>
          {user.role || '—'}
        </div> */}
      </Menu.Item>
      <Menu.Divider />
      {/* <Menu.Item>
        <div>
          <strong>
            <FormattedMessage id="topBar.profileMenu.email" />:{' '}
          </strong>
          {user.email || '—'}
          <br />
          <strong>
            <FormattedMessage id="topBar.profileMenu.phone" />:{' '}
          </strong>
          {user.phone || '—'}
        </div>
      </Menu.Item> */}
      {/* <Menu.Divider /> */}
      {/* <Menu.Item>
        <a href="#" onClick={e => e.preventDefault()}>
          <i className="fe fe-user mr-2" />
          <FormattedMessage id="topBar.profileMenu.editProfile" />
        </a>
      </Menu.Item>
      <Menu.Divider /> */}
      <Menu.Item>
        <a onClick={logout}>
          <i className='fe fe-log-out mr-2' />
          <FormattedMessage id='topBar.profileMenu.logout' />
        </a>
      </Menu.Item>
    </Menu>
  )
  return (
    <Dropdown overlay={menu} trigger={['click']} onVisibleChange={addCount}>
      <div className={styles.dropdown}>
        {/* <Badge count={count}> */}
        <Avatar className={styles.avatar} shape='square' size='large' icon={<UserOutlined />} />
        {/* </Badge> */}
      </div>
    </Dropdown>
  )
}

export default connect(mapStateToProps)(ProfileMenu)
