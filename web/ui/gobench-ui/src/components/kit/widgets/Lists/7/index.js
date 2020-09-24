import React from 'react'
import style from './style.module.scss'

const List7 = () => {
  return (
    <div>
      <ul className={`${style.list} list-unstyled`}>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={style.itemMeta}>
              <div className={`${style.donut} ${style.gray2} ${style.md}`} />
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
            <div className={style.itemMeta}>
              <div className={`${style.donut} ${style.success} ${style.md}`} />
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
            <div className={style.itemMeta}>
              <div className={`${style.donut} ${style.danger} ${style.md}`} />
            </div>
            <div className="mr-3">
              <div>User Deleted</div>
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
            <div className={style.itemMeta}>
              <div className={`${style.donut} ${style.gray2} ${style.md}`} />
            </div>
            <div className="mr-3">
              <div>Message Received</div>
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
            <div className={style.itemMeta}>
              <div className={`${style.donut} ${style.info} ${style.md}`} />
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
            <div className={style.itemMeta}>
              <div className={`${style.donut} ${style.gray2} ${style.md}`} />
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
      </ul>
    </div>
  )
}

export default List7
