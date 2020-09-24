import React from 'react'
import style from './style.module.scss'

const List18 = () => {
  return (
    <ul className={`list-unstyled ${style.list}`}>
      <li className={`${style.item} text-muted`}>
        <div className="text-uppercase mb-1">Organic search</div>
        <div>
          <div className="text-nowrap d-inline-block">
            <div className={`${style.donut} ${style.danger}`} />
            <span className="font-weight-bold text-gray-6">1,125,367</span>
          </div>
          +25%
        </div>
      </li>
      <li className={`${style.item} text-muted`}>
        <div className="text-uppercase mb-1">Google</div>
        <div>
          <div className="text-nowrap d-inline-block">
            <div className={`${style.donut} ${style.primary}`} />
            <span className="font-weight-bold text-gray-6">28,235</span>
          </div>
          +85%
        </div>
      </li>
      <li className={`${style.item} text-muted`}>
        <div className="text-uppercase mb-1">Microsoft</div>
        <div>
          <div className="text-nowrap d-inline-block">
            <div className={`${style.donut} ${style.success}`} />
            <span className="font-weight-bold text-gray-6">874,125</span>
          </div>
          +16%
        </div>
      </li>
      <li className={`${style.item} text-muted`}>
        <div className="text-uppercase mb-1">Yandex</div>
        <div>
          <div className="text-nowrap d-inline-block">
            <div className={`${style.donut} ${style.orange}`} />
            <span className="font-weight-bold text-gray-6">28,235</span>
          </div>
          +154%
        </div>
      </li>
      <li className={`${style.item} text-muted`}>
        <div className="text-uppercase mb-1">Bing Search</div>
        <div>
          <div className="text-nowrap d-inline-block">
            <div className={`${style.donut} ${style.default}`} />
            <span className="font-weight-bold text-gray-6">3,267</span>
          </div>
          +87%
        </div>
      </li>
      <li className={`${style.item} text-muted`}>
        <div className="text-uppercase mb-1">Chinese aw inc</div>
        <div>
          <div className="text-nowrap d-inline-block">
            <div className={`${style.donut} ${style.default}`} />
            <span className="font-weight-bold text-gray-6">51,008</span>
          </div>
          +28%
        </div>
      </li>
    </ul>
  )
}

export default List18
