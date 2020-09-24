import React from 'react'
import style from './style.module.scss'

const General19 = () => {
  return (
    <div>
      <div className="rounded overflow-hidden position-relative">
        <img className="img-fluid" src="resources/images/content/photos/6.jpeg" alt="Lion" />
        <div className={`${style.headerText} font-weight-bold text-white`}>
          Clean. Simple. <br />
          Responsive
        </div>
      </div>
      <div className="d-flex flex-column flex-sm-row">
        <div className={`${style.user} text-center pl-4 pr-5 flex-shrink-0`}>
          <div className="kit__utils__avatar kit__utils__avatar--rounded kit__utils__avatar--size84 border border-5 border-white d-inline-block">
            <img src="resources/images/avatars/2.jpg" alt="Mary Stanform" />
          </div>
          <div className="font-size-14 font-weight-bold">Helen Maggie</div>
          <a href="#" className="font-size-14 text-gray-5">
            @hellen_m
          </a>
        </div>
        <p className="pt-3 mb-0">
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ab aspernatur blanditiis debitis
          deleniti distinctio ducimus et ex incidunt maxime minima minus nemo nisi quos repellendus,
          temporibus ullam veniam voluptas voluptate.w
        </p>
      </div>
    </div>
  )
}

export default General19
