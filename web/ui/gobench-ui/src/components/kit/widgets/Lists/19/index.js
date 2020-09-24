import React from 'react'
import style from './style.module.scss'

const List19 = () => {
  return (
    <ul className="list-unstyled">
      <li className={style.item}>
        <div className={`${style.itemTime} mr-3`}>16:00</div>
        <div className={style.itemSeparator}>
          <div className={`${style.donut} ${style.danger} mr-3`} />
        </div>
        <div>
          Lorem ipsum dolor sit amit,consectetur eiusmdd tempor incididunt ut labore et dolore magna
          elit enim at minim veniam quis nostrud
        </div>
      </li>
      <li className={style.item}>
        <div className={`${style.itemTime} mr-3`}>15:28</div>
        <div className={style.itemSeparator}>
          <div className={`${style.donut} ${style.danger} mr-3`} />
        </div>
        <div>David Beckham was registered as administrator</div>
      </li>
      <li className={style.item}>
        <div className={`${style.itemTime} mr-3`}>14:26</div>
        <div className={style.itemSeparator}>
          <div className={`${style.donut} ${style.danger} mr-3`} />
        </div>
        <div>
          Lorem ipsum dolor sit amit,consectetur eiusmdd tempor incididunt ut labore et dolore
        </div>
      </li>
      <li className={style.item}>
        <div className={`${style.itemTime} mr-3`}>13:22</div>
        <div className={style.itemSeparator}>
          <div className={`${style.donut} ${style.danger} mr-3`} />
        </div>
        <div>
          Lorem ipsum dolor sit amit,consectetur eiusmdd tempor incididunt ut labore et dolore magna
          elit enim at minim veniam quis nostrud
        </div>
      </li>
    </ul>
  )
}

export default List19
