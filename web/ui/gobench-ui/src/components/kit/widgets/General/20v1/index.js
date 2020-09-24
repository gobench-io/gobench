import React from 'react'
import { Carousel } from 'antd'
import style from './style.module.scss'

const General20v1 = () => {
  return (
    <div className={style.carousel}>
      <Carousel>
        <div className="height-150 pt-4 pl-4 overflow-hidden position-relative bg-light">
          <h2 className="font-size-24 font-weight-bold">Sales Growth</h2>
          <a className="font-size-18" href="#">
            View Report
          </a>
          <i className={`${style.icon} fe fe-server text-gray-4`} />
        </div>
        <div className="height-150 pt-4 pl-4 overflow-hidden position-relative bg-light">
          <h2 className="font-size-24 font-weight-bold">To Do</h2>
          <a className="font-size-18" href="#">
            Upgrade / Finish
          </a>
          <i className={`${style.icon} fe fe-settings text-gray-4`} />
        </div>
      </Carousel>
    </div>
  )
}

export default General20v1
