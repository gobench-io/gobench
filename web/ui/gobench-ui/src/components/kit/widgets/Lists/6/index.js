import React from 'react'
import style from './style.module.scss'

const List6 = () => {
  return (
    <div>
      <ul className="list-unstyled">
        <li className={style.item}>
          <div className={style.itemPic}>
            <i>B</i>
          </div>
          <div className="mr-2">
            <div>
              <strong className="text-dark">Bitcoin</strong>
            </div>
            <div className="text-muted">Deposited</div>
          </div>
          <div className="ml-auto text-right text-nowrap">
            <div>+0.005 BTC</div>
            <div className="text-success">+187.89$</div>
          </div>
        </li>
        <li className={style.item}>
          <div className={style.itemPic}>
            <i>L</i>
          </div>
          <div className="mr-2">
            <div>
              <strong className="text-dark">Litecoin</strong>
            </div>
            <div className="text-muted">Deposited by PayPal</div>
          </div>
          <div className="ml-auto text-right text-nowrap">
            <div>+2.3657 LTC</div>
            <div className="text-success">+187.89$</div>
          </div>
        </li>
        <li className={style.item}>
          <div className={style.itemPic}>
            <i>L</i>
          </div>
          <div className="mr-2">
            <div>
              <strong className="text-dark">Dash</strong>
            </div>
            <div className="text-muted">To Dash adress</div>
          </div>
          <div className="ml-auto text-right text-nowrap">
            <div>-2.3657 LTC</div>
            <div className="text-danger">-2187.89$</div>
          </div>
        </li>
        <li className={style.item}>
          <div className={style.itemPic}>
            <i>B</i>
          </div>
          <div className="mr-2">
            <div>
              <strong className="text-dark">Bitcoin</strong>
            </div>
            <div className="text-muted">Deposited</div>
          </div>
          <div className="ml-auto text-right text-nowrap">
            <div>+0.005 BTC</div>
            <div className="text-success">+187.89$</div>
          </div>
        </li>
        <li className={style.item}>
          <div className={style.itemPic}>
            <i>L</i>
          </div>
          <div className="mr-2">
            <div>
              <strong className="text-dark">Litecoin</strong>
            </div>
            <div className="text-muted">Deposited by PayPal</div>
          </div>
          <div className="ml-auto text-right text-nowrap">
            <div>+2.3657 LTC</div>
            <div className="text-success">+187.89$</div>
          </div>
        </li>
      </ul>
    </div>
  )
}

export default List6
