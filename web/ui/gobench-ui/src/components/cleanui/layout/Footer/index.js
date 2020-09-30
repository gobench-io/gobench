import React from 'react'
import style from './style.module.scss'

const Footer = () => {
  return (
    <div className={style.footer}>
      <div className={style.footerInner}>
        {/* <a
          href="https://google.com"
          target="_blank"
          rel="noopener noreferrer"
          className={style.logo}
        >
          Gobench
          <span />
        </a>
        <br /> */}
        <p className='mb-0'>
          <img width='24px' src='/resources/images/GitHub-Mark-32px.png' />
          <a href='https://github.com/gobench-io/gobench' target='_blank' rel='noopener noreferrer'>
          &nbsp;Documentation
          </a>
        </p>
      </div>
    </div>
  )
}

export default Footer
