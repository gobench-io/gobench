import React from 'react'
import style from './style.module.scss'

const List16 = () => {
  return (
    <ul className="list-unstyled">
      <li className={style.item}>
        <div className="font-weight-bold mr-3 font-size-18">16:00</div>
        <div className={`${style.separator} bg-success mr-3`} />
        <div>
          <div>Payment Received</div>
          <div className="text-muted">From themeforest</div>
        </div>
      </li>
      <li className={style.item}>
        <div className="font-weight-bold mr-3 font-size-18">15:28</div>
        <div className={`${style.separator} bg-primary mr-3`} />
        <div>
          <div>Payment Received</div>
          <div className="text-muted">From themeforest</div>
        </div>
      </li>
      <li className={style.item}>
        <div className="font-weight-bold mr-3 font-size-18">14:26</div>
        <div className={`${style.separator} mr-3`} />
        <div>
          <div>Payment Received</div>
          <div className="text-muted">From Paypal</div>
        </div>
      </li>
      <li className={style.item}>
        <div className="font-weight-bold mr-3 font-size-18">13:57</div>
        <div className={`${style.separator} mr-3 bg-danger`} />
        <div>
          <div>Payment Received</div>
          <div className="text-muted">From Bitcoin Address</div>
        </div>
      </li>
      <li className={style.item}>
        <div className="font-weight-bold mr-3 font-size-18">13:22</div>
        <div className={`${style.separator} bg-info mr-3`} />
        <div>
          <div>Payment Received</div>
          <div className="text-muted">From themeforest</div>
        </div>
      </li>
      <li className={style.item}>
        <div className="font-weight-bold mr-3 font-size-18">11:08</div>
        <div className={`${style.separator} mr-3`} />
        <div>
          <div>Payment Received</div>
          <div className="text-muted">From themeforest</div>
        </div>
      </li>
      <li className={style.item}>
        <div className="font-weight-bold mr-3 font-size-18">10:01</div>
        <div className={`${style.separator} mr-3`} />
        <div>
          <div>Payment Received</div>
          <div className="text-muted">From themeforest</div>
        </div>
      </li>
    </ul>
  )
}

export default List16
