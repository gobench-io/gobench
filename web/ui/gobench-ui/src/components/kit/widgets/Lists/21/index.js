import React from 'react'

const List21 = () => {
  return (
    <div className="card">
      <div className="card-body">
        <div className="pt-5 pb-5 pl-5 pr-5 text-center flex-grow-1">
          <i className="fe fe-inbox font-size-80 mb-3 d-block" />
          <div className="text-dark font-weight-bold font-size-36">Free Plan</div>
          <div className="text-dark font-weight-bold font-size-48 mb-3">
            $0 <span className="align-text-top font-size-28 text-gray-6">/mo</span>
          </div>
          <ul className="list-unstyled font-size-18 mb-5">
            <li>10GB of Bandwidth</li>
            <li>200MB Max File Size</li>
            <li>2GHZ CPU</li>
            <li>256MB Memory</li>
            <li>1 GB Storage</li>
          </ul>
          <a className="btn btn-primary width-100" href="">
            Get Access
          </a>
        </div>
      </div>
    </div>
  )
}

export default List21
