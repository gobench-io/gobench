import React, { useState } from 'react'
import style from './style.module.scss'

const General16 = ({ isFavourite, isNew, image, name, price, oldPrice }) => {
  const [favourite, setFavourite] = useState(isFavourite)

  const setIsFavourite = e => {
    e.preventDefault()
    setFavourite(!favourite)
  }

  return (
    <div className="card overflow-hidden">
      <div hidden={!isNew} className={style.new}>
        New
      </div>
      <div className="card-body">
        <a
          role="menuitem"
          className={`${style.favourite} ${favourite ? 'text-dark' : 'text-gray-3'}`}
          onClick={setIsFavourite}
          onKeyPress={setIsFavourite}
          tabIndex="0"
        >
          <i className="fe fe-heart font-size-21" />
        </a>
        <div className={`${style.image} border-bottom height-250 mb-3`}>
          <img className="img-fluid" src={image} alt={name} />
        </div>
        <div className="font-size-24 font-weight-bold text-dark mb-2">
          {price}{' '}
          <del hidden={!oldPrice} className="align-text-top font-size-14">
            {oldPrice}
          </del>
        </div>
        <div>
          <a className="text-blue font-size-18">{name}</a>
        </div>
      </div>
    </div>
  )
}

export default General16
