import React from 'react'
import style from './style.module.scss'

const General6 = () => {
  return (
    <div className={`${style.container} pt-3`}>
      <div className={`${style.status} bg-danger`} />
      <div className="d-flex flex-nowrap align-items-center pb-3 pl-4 pr-4">
        <div className="mr-auto">
          <div className="text-uppercase font-weight-bold font-size-24 text-dark">-$1,125</div>
          <div className="font-size-18">4512-XXXX-1678-7528</div>
        </div>
        <div className="ml-1 text-danger">
          <i className="fe fe-arrow-right-circle font-size-40" />
        </div>
      </div>
      <div className={`${style.footer} py-3 pl-4`}>To DigitalOcean Cloud Hosting, Winnetka, LA</div>
    </div>
  )
}

export default General6
