import React from 'react'
import style from './style.module.scss'

const List3 = () => {
  return (
    <div>
      <ul className="list-unstyled">
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={`${style.itemMeta} font-weight-bold text-danger`}>
              2:28
              <br />
              PM
            </div>
            <div className="mr-3">
              <div>Payment Received</div>
              <div className="text-muted">Mary has approved your quote.</div>
            </div>
            <div className={style.itemAction}>
              <span />
              <span />
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={`${style.itemMeta} font-weight-bold text-blue`}>
              1:02
              <br />
              AM
            </div>
            <div className="mr-3">
              <div>Account Activated</div>
              <div className="text-muted">Mary has approved your quote.</div>
            </div>
            <div className={style.itemAction}>
              <span />
              <span />
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={`${style.itemMeta} font-weight-bold`}>
              2:28
              <br />
              PM
            </div>
            <div className="mr-3">
              <div>Payment Received</div>
              <div className="text-muted">Mary has approved your quote.</div>
            </div>
            <div className={style.itemAction}>
              <span />
              <span />
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={`${style.itemMeta} font-weight-bold`}>
              2:28
              <br />
              PM
            </div>
            <div className="mr-3">
              <div>Payment Received</div>
              <div className="text-muted">Mary has approved your quote.</div>
            </div>
            <div className={style.itemAction}>
              <span />
              <span />
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={`${style.itemMeta} font-weight-bold`}>
              2:28
              <br />
              PM
            </div>
            <div className="mr-3">
              <div>Payment Received</div>
              <div className="text-muted">Mary has approved your quote.</div>
            </div>
            <div className={style.itemAction}>
              <span />
              <span />
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={`${style.itemMeta} font-weight-bold`}>
              2:28
              <br />
              PM
            </div>
            <div className="mr-3">
              <div>Payment Received</div>
              <div className="text-muted">Mary has approved your quote.</div>
            </div>
            <div className={style.itemAction}>
              <span />
              <span />
            </div>
          </a>
        </li>
      </ul>
    </div>
  )
}

export default List3
