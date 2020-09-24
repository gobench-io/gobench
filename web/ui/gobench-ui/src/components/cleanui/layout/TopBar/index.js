import React from 'react'
// import FavPages from './FavPages'
// import Search from './Search'
// import IssuesHistory from './IssuesHistory'
// import ProjectManagement from './ProjectManagement'
// import LanguageSwitcher from './LanguageSwitcher'
// import Actions from './Actions'
import UserMenu from './UserMenu'
import style from './style.module.scss'
import { Link } from 'react-router-dom'

const TopBar = () => {
  return (
    <div className={style.topbar}>
      {/* <div className="mr-4">
        <FavPages />
      </div> */}
      <div className='logo'>
        <Link to='/'>
          <img src='/resources/images/logo-white.png' width='200px' />
        </Link>
      </div>
      <div className='mr-auto'>
        {/* <Search /> */}
      </div>
      <div className='text-right' style={{ marginRight: 10, opacity: 0.7 }}>
        <img width='24px' src='/resources/images/GitHub-Mark-32px.png' />
        <a href='https://github.com/gobench-io/gobench' target='_blank' rel='noopener noreferrer'>
          &nbsp;Documentation
        </a>
      </div>
      {/* <div className="mr-4 d-none d-md-block">
        <IssuesHistory />
      </div>
      <div className="mb-0 mr-auto d-xl-block d-none">
        <ProjectManagement />
      </div>
      <div className="mr-4 d-none d-sm-block">
        <LanguageSwitcher />
      </div>
      <div className="mr-4 d-none d-sm-block">
        <Actions />
      </div> */}
      <div className='text-right'>
        <UserMenu />
      </div>
    </div>
  )
}

export default TopBar
