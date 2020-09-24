import React from 'react'
import { Tooltip } from 'antd'
import { Helmet } from 'react-helmet'
import { iconsData } from './data.json'

const IconsFeatherIcons = () => {
  return (
    <div>
      <Helmet title="Icons / Feather Icons" />
      <div className="kit__utils__heading">
        <h5>
          <span className="mr-3">Feather Icons</span>
          <a
            href="https://linearicons.com/free"
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
                <strong>Simply beautiful open source icons</strong>
              </h5>
              <p className="text-muted">
                The complete set of 279 icons in Feathericons v4.21.0
                <br />
                License: MIT License. You can use it for commercial projects, open source projects,
                or really just about whatever you want
              </p>
            </div>
          </div>
          <div className="row">
            <div className="col-xl-6 offset-xl-3">
              <h3 className="text-block mt-5 mb-4 text-center">
                <strong>General Icons</strong>
              </h3>
              <ul className="kit__utils__iconPresent list-unstyled">
                {iconsData.map(icon => (
                  <Tooltip title={`fe ${icon}`} key={icon}>
                    <li>
                      <i className={`fe ${icon}`} />
                    </li>
                  </Tooltip>
                ))}
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default IconsFeatherIcons
