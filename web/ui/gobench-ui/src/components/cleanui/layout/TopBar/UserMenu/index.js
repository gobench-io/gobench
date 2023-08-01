import React, { useState } from 'react'
import { connect } from 'react-redux'
import { UserOutlined } from '@ant-design/icons'
import { Dropdown, Avatar } from 'antd'
import styles from './style.module.scss'

const mapStateToProps = ({ user }) => ({ user })

const ProfileMenu = ({ dispatch }) => {
  const [showMenu, setShowMenu] = useState(false)

  const toggleShowMenu = () => {
    setShowMenu(!showMenu)
  }
  const items = [
    { label: <strong>Administrator</strong>, key: 'item-1' },
    {
      label: 'log out',
      key: 'item-2',
      icon: <i className='fe fe-log-out me-2' />,
      onClick: () => dispatch({
        type: 'user/LOGOUT'
      })
    },
  ]
  return (
    <Dropdown open={showMenu} menu={{ items }} trigger={['click']} onOpenChange={toggleShowMenu}>
      <div className={styles.dropdown}>
        <Avatar className={styles.avatar} shape='square' size='large' icon={<UserOutlined />} />
      </div>
    </Dropdown>
  )
}

export default connect(mapStateToProps)(ProfileMenu)
