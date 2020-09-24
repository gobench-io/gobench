import React from 'react'
import style from './style.module.scss'

const List4 = () => {
  return (
    <div>
      <ul className="list-unstyled">
        <li className={style.item}>
          <div className={`${style.itemHead} mb-3`}>
            <div className={style.itemPic}>
              <img src="resources/images/avatars/1.jpg" alt="Mary Stanform" />
            </div>
            <div className="mr-2">
              <div>Jamie Rockstar</div>
              <div className="text-muted">Backoffice Agent</div>
            </div>
            <div className="text-success ml-auto">Active</div>
          </div>
          <p className="mb-4">
            Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum
            has been the industry&apos;s standard dummy text ...
          </p>
        </li>
        <li className={style.item}>
          <div className={`${style.itemHead} mb-3`}>
            <div className={style.itemPic}>
              <img src="resources/images/avatars/2.jpg" alt="Mary Stanform" />
            </div>
            <div className="mr-2">
              <div>Alex Kasie</div>
              <div className="text-muted">Support Agent</div>
            </div>
            <div className="text-danger ml-auto">Suspended</div>
          </div>
          <p className="mb-4">
            Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum
            has been the industry&apos;s standard dummy text ...
          </p>
        </li>
      </ul>
    </div>
  )
}

export default List4
