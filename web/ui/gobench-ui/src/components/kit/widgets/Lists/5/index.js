import React from 'react'
import style from './style.module.scss'

const List5 = () => {
  return (
    <div>
      <ul className="list-unstyled">
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={style.itemPic}>
              <i className="fe fe-file-text" />
            </div>
            <div className="mr-2">
              <div>Payment Received</div>
              <div className="text-muted">3 minutes ago</div>
            </div>
            <div className={style.itemAction}>
              <span />
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={style.itemPic}>
              <i className="fe fe-mail" />
            </div>
            <div className="mr-2">
              <div>Message Removed</div>
              <div className="text-muted">2 hours ago</div>
            </div>
            <div className={style.itemAction}>
              <span />
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={style.itemPic}>
              <i className="fe fe-grid" />
            </div>
            <div className="mr-2">
              <div>Parcel Received</div>
              <div className="text-muted">6 hours ago</div>
            </div>
            <div className={style.itemAction}>
              <span />
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={style.itemPic}>
              <i className="fe fe-database" />
            </div>
            <div className="mr-2">
              <div>Parcel Recived</div>
              <div className="text-muted">15 hours ago</div>
            </div>
            <div className={style.itemAction}>
              <span />
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={style.itemPic}>
              <i className="fe fe-flag" />
            </div>
            <div className="mr-2">
              <div>User Activated</div>
              <div className="text-muted">2 days ago</div>
            </div>
            <div className={style.itemAction}>
              <span />
            </div>
          </a>
        </li>
      </ul>
    </div>
  )
}

export default List5
