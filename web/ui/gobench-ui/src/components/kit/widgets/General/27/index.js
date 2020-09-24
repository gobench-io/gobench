import React from 'react'
import style from './style.module.scss'

const General27 = () => {
  return (
    <div className={`${style.container} pl-5 pr-5 pt-5 pb-5 mb-auto text-dark font-size-32`}>
      <div className="font-weight-bold mb-3">Page not found</div>
      <div className="text-gray-6 font-size-24">
        This page is deprecated, deleted, or does not exist at all
      </div>
      <div className="font-weight-bold font-size-70 mb-1">404 —</div>
      <a href="" className="btn btn-outline-primary width-100">
        Go Back
      </a>
    </div>
  )
}

export default General27
