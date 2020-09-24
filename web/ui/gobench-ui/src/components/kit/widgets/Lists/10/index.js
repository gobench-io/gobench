import React from 'react'
import style from './style.module.scss'

const List10 = () => {
  return (
    <div>
      <ul className="list-unstyled">
        <li className={style.item}>
          <div className={`${style.itemHead} mb-2`}>
            <div className={style.itemPic}>
              <img src="resources/images/avatars/1.jpg" alt="Jamie Rockstar" />
            </div>
            <div className="mr-2">
              <div>Jamie Rockstar</div>
              <div className="text-muted">Backoffice Agent</div>
            </div>
          </div>
          <div className="text-muted mb-2">
            Access Level: <strong className="text-gray-6">Level 7</strong>
          </div>
          <div className="progress">
            <div
              className="progress-bar bg-success"
              style={{
                width: '70%',
              }}
              role="progressbar"
              aria-valuenow={70}
              aria-valuemin={0}
              aria-valuemax={100}
            />
          </div>
        </li>
        <li className={style.item}>
          <div className={`${style.itemHead} mb-2`}>
            <div className={style.itemPic}>
              <img src="resources/images/avatars/2.jpg" alt="Mary Stanform" />
            </div>
            <div className="mr-2">
              <div>Mary Stanform</div>
              <div className="text-muted">Developer</div>
            </div>
          </div>
          <div className="text-muted mb-2">
            Access Level: <strong className="text-gray-6">Level 4</strong>
          </div>
          <div className="progress">
            <div
              className="progress-bar bg-primary"
              style={{
                width: '40%',
              }}
              role="progressbar"
              aria-valuenow={40}
              aria-valuemin={0}
              aria-valuemax={100}
            />
          </div>
        </li>
        <li className={style.item}>
          <div className={`${style.itemHead} mb-2`}>
            <div className={style.itemPic}>
              <img src="resources/images/avatars/5.jpg" alt="Jess Hofner" />
            </div>
            <div className="mr-2">
              <div>Jess Hofner</div>
              <div className="text-muted">CEO</div>
            </div>
          </div>
          <div className="text-muted mb-2">
            Access Level: <strong className="text-gray-6">Level 9</strong>
          </div>
          <div className="progress">
            <div
              className="progress-bar bg-danger"
              style={{
                width: '90%',
              }}
              role="progressbar"
              aria-valuenow={90}
              aria-valuemin={0}
              aria-valuemax={100}
            />
          </div>
        </li>
      </ul>
    </div>
  )
}

export default List10
