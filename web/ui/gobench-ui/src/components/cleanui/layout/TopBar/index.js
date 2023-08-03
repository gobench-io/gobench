import React from 'react'
import UserMenu from './UserMenu'
import style from './style.module.scss'
import { Link } from 'react-router-dom'

const TopBar = () => {
  return (
    <div className={style.topbar}>
      <div className='logo'>
        <Link to='/'>
          <img src='/resources/images/logo-white.png' width='200px' alt='' />
        </Link>
      </div>
      <div className='me-auto'>
      </div>
      <div className='text-end' style={{ marginRight: 10, opacity: 0.7 }}>
        <img width='24px' src='/resources/images/GitHub-Mark-32px.png' alt='' />
        <a href='https://github.com/gobench-io/gobench' target='_blank' rel='noopener noreferrer'>
          &nbsp;Documentation
        </a>
      </div>
      <div className='text-end'>
        <UserMenu />
      </div>
    </div>
  )
}

export default TopBar
