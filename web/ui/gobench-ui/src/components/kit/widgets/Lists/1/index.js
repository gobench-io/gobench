import React from 'react'
import style from './style.module.scss'

const List1 = () => {
  return (
    <div>
      <div className="text-uppercase font-size-12 mb-2 text-gray-6">Documents (3)</div>
      <ul className="list-unstyled">
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={`${style.itemPic} mr-3`}>
              <i className={`${style.itemIcon} fe fe-file`} />
            </div>
            <div>
              <div className="text-blue">System Requirements.pdf</div>
              <div className="text-muted">568kb</div>
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={`${style.itemPic} mr-3`}>
              <i className={`${style.itemIcon} fe fe-file`} />
            </div>
            <div>
              <div className="text-blue">Queue Info.pdf</div>
              <div className="text-muted">1.2mb</div>
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className={`${style.itemPic} mr-3`}>
              <i className={`${style.itemIcon} fe fe-file`} />
            </div>
            <div>
              <div className="text-blue">Affected_app.mov</div>
              <div className="text-muted">67mb</div>
            </div>
          </a>
        </li>
      </ul>
      <div className="text-uppercase font-size-12 mb-2 text-gray-6">Agents (4)</div>
      <ul className="list-unstyled">
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className="kit__utils__avatar mr-3">
              <img src="resources/images/avatars/5.jpg" alt="Mary Stanform" />
            </div>
            <div>
              <div className="text-blue">Mary Stanform</div>
              <div className="text-muted">Sales Manager</div>
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className="kit__utils__avatar mr-3">
              <img src="resources/images/avatars/1.jpg" alt="Jamie Rockstar" />
            </div>
            <div>
              <div className="text-blue">Jamie Rockstar</div>
              <div className="text-muted">Blackoffice Agent</div>
            </div>
          </a>
        </li>
        <li className={style.item}>
          <a className={style.itemLink}>
            <div className="kit__utils__avatar mr-3">
              <img src="resources/images/avatars/4.jpg" alt="Jamie Rockstar" />
            </div>
            <div>
              <div className="text-blue">David Bowie</div>
              <div className="text-muted">Blackoffice Agent</div>
            </div>
          </a>
        </li>
      </ul>
    </div>
  )
}

export default List1
