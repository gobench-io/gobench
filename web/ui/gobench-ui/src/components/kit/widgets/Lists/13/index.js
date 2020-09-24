import React from 'react'
import style from './style.module.scss'

const List13 = () => {
  return (
    <div>
      <div className="text-uppercase mb-3">People you may know</div>
      <ul className="list-unstyled">
        <li className={style.item}>
          <div className={style.itemPic}>
            <img src="resources/images/avatars/1.jpg" alt="Jamie Rockstar" />
          </div>
          <div>
            <div>Jamie Rockstar</div>
            <div className="text-muted">Backoffice Agent</div>
          </div>
        </li>
        <li className={style.item}>
          <div className={style.itemPic}>
            <img src="resources/images/avatars/2.jpg" alt="Katie Banks" />
          </div>
          <div>
            <div>Katie Banks</div>
            <div className="text-muted">Support Agent</div>
          </div>
        </li>
        <li className={style.item}>
          <div className={style.itemPic}>
            <img src="resources/images/avatars/3.jpg" alt="Jessey Kim" />
          </div>
          <div>
            <div>Jessey Kim</div>
            <div className="text-muted">Administrator</div>
          </div>
        </li>
        <li className={style.item}>
          <div className={style.itemPic}>
            <img src="resources/images/avatars/4.jpg" alt="Sam Piterson" />
          </div>
          <div>
            <div>Sam Piterson</div>
            <div className="text-muted">Technical Assistant</div>
          </div>
        </li>
        <li className={style.item}>
          <div className={style.itemPic}>
            <img src="resources/images/avatars/5.jpg" alt="Mary Stanform" />
          </div>
          <div>
            <div>Mary Stanform</div>
            <div className="text-muted">Illustrator</div>
          </div>
        </li>
      </ul>
    </div>
  )
}

export default List13
