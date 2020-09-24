import React from 'react'
import style from './style.module.scss'

const List20 = () => {
  return (
    <ul className="list-unstyled">
      <li className={style.item}>
        <div className={style.itemPic}>
          <img src="resources/images/avatars/1.jpg" alt="Jamie Rockstar" />
        </div>
        <div className="flex-fill">
          <div className="font-weight-bold text-dark">Jamie Rockstar</div>
          <div className="text-muted">Backoffice Agent</div>
        </div>
        <a href="" className="btn btn-outline-primary align-self-end">
          Follow
        </a>
      </li>
      <li className={style.item}>
        <div className={style.itemPic}>
          <img src="resources/images/avatars/2.jpg" alt="Katie Banks" />
        </div>
        <div className="flex-fill">
          <div className="font-weight-bold text-dark">Katie Banks</div>
          <div className="text-muted">Support Agent</div>
        </div>
        <a href="" className="btn btn-outline-primary align-self-end">
          Follow
        </a>
      </li>
      <li className={style.item}>
        <div className={style.itemPic}>
          <img src="resources/images/avatars/3.jpg" alt="Jessey Kim" />
        </div>
        <div className="flex-fill">
          <div className="font-weight-bold text-dark">Jessey Kim</div>
          <div className="text-muted">Administrator</div>
        </div>
        <a href="" className="btn btn-outline-primary align-self-end">
          Follow
        </a>
      </li>
      <li className={style.item}>
        <div className={style.itemPic}>
          <img src="resources/images/avatars/4.jpg" alt="Sam Piterson" />
        </div>
        <div className="flex-fill">
          <div className="font-weight-bold text-dark">Sam Piterson</div>
          <div className="text-muted">Technical Assistant</div>
        </div>
        <a href="" className="btn btn-outline-primary align-self-end">
          Follow
        </a>
      </li>
      <li className={style.item}>
        <div className={style.itemPic}>
          <img src="resources/images/avatars/5.jpg" alt="Mary Stanform" />
        </div>
        <div className="flex-fill">
          <div className="font-weight-bold text-dark">Mary Stanform</div>
          <div className="text-muted">Illustrator</div>
        </div>
        <a href="" className="btn btn-outline-primary align-self-end">
          Follow
        </a>
      </li>
    </ul>
  )
}

export default List20
