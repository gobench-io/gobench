import React from 'react'
import style from './style.module.scss'

const List17 = () => {
  return (
    <ul className="list-unstyled">
      <li className={style.item}>
        <div className={`${style.separator} bg-success mr-3`} />
        <label htmlFor="checkbox-1" className={`${style.control} ${style.checkbox} mb-0`}>
          <input type="checkbox" id="checkbox-1" />
          <span className={`${style.controlIndicator}`} />
          <div className="d-inline-block">
            <div>Payment Received</div>
            <div className="text-muted">From themeforest</div>
          </div>
        </label>
      </li>
      <li className={style.item}>
        <div className={`${style.separator} bg-primary mr-3`} />
        <label htmlFor="checkbox-2" className={`${style.control} ${style.checkbox} mb-0`}>
          <input type="checkbox" id="checkbox-2" />
          <span className={`${style.controlIndicator}`} />
          <div className="d-inline-block">
            <div>Payment Approved</div>
            <div className="text-muted">From themeforest</div>
          </div>
        </label>
      </li>
      <li className={style.item}>
        <div className={`${style.separator} mr-3`} />
        <label htmlFor="checkbox-3" className={`${style.control} ${style.checkbox} mb-0`}>
          <input type="checkbox" id="checkbox-3" />
          <span className={`${style.controlIndicator}`} />
          <div className="d-inline-block">
            <div>Payment Received</div>
            <div className="text-muted">From Paypal</div>
          </div>
        </label>
      </li>
      <li className={style.item}>
        <div className={`${style.separator} bg-danger mr-3`} />
        <label htmlFor="checkbox-4" className={`${style.control} ${style.checkbox} mb-0`}>
          <input type="checkbox" id="checkbox-4" />
          <span className={`${style.controlIndicator}`} />
          <div className="d-inline-block">
            <div>Withdrawal Failed</div>
            <div className="text-muted">From Bitcoin Address</div>
          </div>
        </label>
      </li>
      <li className={style.item}>
        <div className={`${style.separator} bg-info mr-3`} />
        <label htmlFor="checkbox-5" className={`${style.control} ${style.checkbox} mb-0`}>
          <input type="checkbox" id="checkbox-5" />
          <span className={`${style.controlIndicator}`} />
          <div className="d-inline-block">
            <div>Payment Received</div>
            <div className="text-muted">From themeforest</div>
          </div>
        </label>
      </li>
      <li className={style.item}>
        <div className={`${style.separator} mr-3`} />
        <label htmlFor="checkbox-6" className={`${style.control} ${style.checkbox} mb-0`}>
          <input type="checkbox" id="checkbox-6" />
          <span className={`${style.controlIndicator}`} />
          <div className="d-inline-block">
            <div>Payment Received</div>
            <div className="text-muted">From themeforest</div>
          </div>
        </label>
      </li>
      <li className={style.item}>
        <div className={`${style.separator} mr-3`} />
        <label htmlFor="checkbox-7" className={`${style.control} ${style.checkbox} mb-0`}>
          <input type="checkbox" id="checkbox-7" />
          <span className={`${style.controlIndicator}`} />
          <div className="d-inline-block">
            <div>Payment Received</div>
            <div className="text-muted">From themeforest</div>
          </div>
        </label>
      </li>
    </ul>
  )
}

export default List17
