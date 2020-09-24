import React from 'react'
import { Tooltip } from 'antd'
import { Helmet } from 'react-helmet'
import { iconsData } from './data.json'

const IconsFontawesome = () => {
  return (
    <div>
      <Helmet title="Icons / Fontawesome" />
      <div className="kit__utils__heading">
        <h5>
          <span className="mr-3">Fontawesome</span>
          <a
            href="http://fontawesome.io/"
            rel="noopener noreferrer"
            target="_blank"
            className="btn btn-sm btn-light"
          >
            Official Documentation
            <i className="fe fe-corner-right-up" />
          </a>
        </h5>
      </div>
      <div className="card">
        <div className="card-body">
          <div className="row">
            <div className="col-lg-12">
              <h5 className="text-black">
                <strong>The iconic font</strong>
              </h5>
              <p className="text-muted">
                The complete set of 634 icons in Font Awesome 4.6.3
                <br />
                License: MIT License. You can use it for commercial projects, open source projects,
                or really just about whatever you want
              </p>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 offset-xl-3">
              {iconsData.map(iconsSet => (
                <div className="text-center" key={iconsSet.setName}>
                  <h3 className="text-block mt-5 mb-4">
                    <strong>{iconsSet.setName}</strong>
                  </h3>
                  <ul className="kit__utils__iconPresent list-unstyled">
                    {iconsSet.icons.map(icon => (
                      <Tooltip title={`fa ${icon}`} key={icon}>
                        <li>
                          <i className={`fa ${icon}`} />
                        </li>
                      </Tooltip>
                    ))}
                  </ul>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default IconsFontawesome
