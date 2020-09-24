import React from 'react'
import style from './style.module.scss'

const List9 = () => {
  return (
    <div>
      <ul className="list-unstyled">
        <li className={style.item}>
          <div
            className={style.itemDots}
            style={{ backgroundImage: 'url(resources/images/3-rounds.png)' }}
          />
          <div className={style.itemPicContainer}>
            <div className={`${style.itemPic} bg-success`} />
            <i className={`${style.itemIcon} text-success fe fe-file-text`} />
          </div>
          <div>
            <div>
              <strong className="text-primary">Bitcoin</strong> lorem Ipsum is simply dummy text of
              the printing and...
            </div>
            <div className="text-muted">Deposited</div>
          </div>
        </li>
        <li className={style.item}>
          <div
            className={style.itemDots}
            style={{ backgroundImage: 'url(resources/images/3-rounds.png)' }}
          />
          <div className={style.itemPicContainer}>
            <div className={`${style.itemPic} bg-info`} />
            <i className={`${style.itemIcon} text-info fe fe-mail`} />
          </div>
          <div>
            <div>
              <strong className="text-primary">Litecoint</strong> lorem Ipsum is simply dummy text
              of the printing and...
            </div>
            <div className="text-muted">Deposited by PayPal</div>
          </div>
        </li>
        <li className={style.item}>
          <div
            className={style.itemDots}
            style={{ backgroundImage: 'url(resources/images/3-rounds.png)' }}
          />
          <div className={style.itemPicContainer}>
            <div className={`${style.itemPic} bg-danger`} />
            <i className={`${style.itemIcon} text-danger fe fe-grid`} />
          </div>
          <div>
            <div>
              <strong>Dash</strong> lorem Ipsum is simply dummy text of the printing and...
            </div>
            <div className="text-muted">To Dash adress</div>
          </div>
        </li>
        <li className={style.item}>
          <div
            className={style.itemDots}
            style={{ backgroundImage: 'url(resources/images/3-rounds.png)' }}
          />
          <div className={style.itemPicContainer}>
            <div className={`${style.itemPic} bg-primary`} />
            <i className={`${style.itemIcon} text-primary fe fe-database`} />
          </div>
          <div>
            <div>
              <strong>Bitcoin</strong> lorem Ipsum is simply dummy text of the printing and...
            </div>
            <div className="text-muted">Deposited</div>
          </div>
        </li>
        <li className={style.item}>
          <div
            className={style.itemDots}
            style={{ backgroundImage: 'url(resources/images/3-rounds.png)' }}
          />
          <div className={style.itemPicContainer}>
            <div className={`${style.itemPic} bg-success`} />
            <i className={`${style.itemIcon} text-success fe fe-flag`} />
          </div>
          <div>
            <div>
              <strong>Litecoin</strong> lorem Ipsum is simply dummy text of the printing and...
            </div>
            <div className="text-muted">Deposited by PayPal</div>
          </div>
        </li>
      </ul>
    </div>
  )
}

export default List9
