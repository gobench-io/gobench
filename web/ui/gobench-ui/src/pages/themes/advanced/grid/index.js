import React from 'react'
import { Helmet } from 'react-helmet'

const AdvancedGrid = () => {
  return (
    <div>
      <Helmet title="Advanced / Grid" />
      <div className="kit__utils__heading">
        <h5>
          <span className="mr-3">Grid</span>
          <a
            href="https://getbootstrap.com/docs/4.3/layout/grid/"
            target="_blank"
            rel="noopener noreferrer"
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
            <div className="col-lg-12 mb-5">
              <h5 className="mb-4">
                <strong>How it works</strong>
              </h5>
              <div className="row text-center">
                <div className="col-sm">
                  <div className="mb-4 px-3 py-2 bg-light">One of three columns</div>
                </div>
                <div className="col-sm">
                  <div className="mb-4 px-3 py-2 bg-light">One of three columns</div>
                </div>
                <div className="col-sm">
                  <div className="mb-4 px-3 py-2 bg-light">One of three columns</div>
                </div>
              </div>
            </div>
            <div className="col-lg-12 mb-5">
              <h5 className="mb-4">
                <strong>Grid options</strong>
              </h5>
              <table className="table table-bordered table-striped">
                <thead>
                  <tr>
                    <th />
                    <th className="text-center">
                      Extra small
                      <br />
                      <small>&lt;576px</small>
                    </th>
                    <th className="text-center">
                      Small
                      <br />
                      <small>≥576px</small>
                    </th>
                    <th className="text-center">
                      Medium
                      <br />
                      <small>≥768px</small>
                    </th>
                    <th className="text-center">
                      Large
                      <br />
                      <small>≥992px</small>
                    </th>
                    <th className="text-center">
                      Extra large
                      <br />
                      <small>≥1200px</small>
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <th className="text-nowrap" scope="row">
                      Max container width
                    </th>
                    <td>None (auto)</td>
                    <td>540px</td>
                    <td>720px</td>
                    <td>960px</td>
                    <td>1140px</td>
                  </tr>
                  <tr>
                    <th className="text-nowrap" scope="row">
                      Class prefix
                    </th>
                    <td>
                      <code>.col-</code>
                    </td>
                    <td>
                      <code>.col-sm-</code>
                    </td>
                    <td>
                      <code>.col-md-</code>
                    </td>
                    <td>
                      <code>.col-lg-</code>
                    </td>
                    <td>
                      <code>.col-xl-</code>
                    </td>
                  </tr>
                  <tr>
                    <th className="text-nowrap" scope="row">
                      # of columns
                    </th>
                    <td colSpan="5">12</td>
                  </tr>
                  <tr>
                    <th className="text-nowrap" scope="row">
                      Gutter width
                    </th>
                    <td colSpan="5">30px (15px on each side of a column)</td>
                  </tr>
                  <tr>
                    <th className="text-nowrap" scope="row">
                      Nestable
                    </th>
                    <td colSpan="5">Yes</td>
                  </tr>
                  <tr>
                    <th className="text-nowrap" scope="row">
                      Column ordering
                    </th>
                    <td colSpan="5">Yes</td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div className="col-lg-12 mb-5">
              <h5 className="mb-4">
                <strong>Stacked-to-Horizontal</strong>
              </h5>
              <div className="row">
                <div className="col-md-8">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-8</div>
                </div>
                <div className="col-md-4">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-4</div>
                </div>
              </div>
              <div className="row">
                <div className="col-md-4">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-4</div>
                </div>
                <div className="col-md-4">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-4</div>
                </div>
                <div className="col-md-4">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-4</div>
                </div>
              </div>
              <div className="row">
                <div className="col-md-6">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-6</div>
                </div>
                <div className="col-md-6">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-6</div>
                </div>
              </div>
              <div className="row">
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
                <div className="col-md-1">
                  <div className="mb-4 px-3 py-2 bg-light">.col-md-1</div>
                </div>
              </div>
            </div>
            <div className="col-lg-12 mb-5">
              <h5 className="mb-4">
                <strong>Desktop & Mobile</strong>
              </h5>
              <div className="row">
                <div className="col-xs-12 col-md-8">
                  <div className="mb-4 px-3 py-2 bg-light">.col-xs-12 .col-md-8</div>
                </div>
                <div className="col-xs-6 col-md-4">
                  <div className="mb-4 px-3 py-2 bg-light">col-xs-6 .col-md-4</div>
                </div>
              </div>
              <div className="row">
                <div className="col-xs-6 col-md-4">
                  <div className="mb-4 px-3 py-2 bg-light">.col-xs-6 .col-md-4</div>
                </div>
                <div className="col-xs-6 col-md-4">
                  <div className="mb-4 px-3 py-2 bg-light">.col-xs-6 .col-md-4</div>
                </div>
                <div className="col-xs-6 col-md-4">
                  <div className="mb-4 px-3 py-2 bg-light">.col-xs-6 .col-md-4</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default AdvancedGrid
