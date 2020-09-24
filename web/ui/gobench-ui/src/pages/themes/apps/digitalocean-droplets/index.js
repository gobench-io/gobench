import React from 'react'
import { Helmet } from 'react-helmet'
import { Tabs, Dropdown, Menu } from 'antd'
import Chart4 from 'components/kit/widgets/Charts/4'
import Chart4v1 from 'components/kit/widgets/Charts/4v1'
import Chart4v2 from 'components/kit/widgets/Charts/4v2'
import Chart4v3 from 'components/kit/widgets/Charts/4v3'
import List5 from 'components/kit/widgets/Lists/5'

const { TabPane } = Tabs

const dropdownMenu = (
  <Menu>
    <Menu.Item>
      <a href="#" onClick={e => e.preventDefault()}>
        Action
      </a>
    </Menu.Item>
    <Menu.Item>
      <a href="#" onClick={e => e.preventDefault()}>
        Another action
      </a>
    </Menu.Item>
    <Menu.Item>
      <a href="#" onClick={e => e.preventDefault()}>
        Something else here
      </a>
    </Menu.Item>
    <div className="dropdown-divider" />
    <Menu.Item>
      <a href="#" onClick={e => e.preventDefault()}>
        Separated link
      </a>
    </Menu.Item>
  </Menu>
)

const ExtraAppsDigitaloceanDroplets = () => {
  return (
    <div>
      <Helmet title="DigitalOcean Droplets" />
      <div className="d-flex flex-wrap align-items-center">
        <div className="kit__utils__avatar kit__utils__avatar--size64 flex-shrink-0 mr-5 mb-3">
          <img src="resources/images/avatars/2.jpg" alt="Mary Stanform" />
        </div>
        <div className="mr-auto mb-3">
          <div className="text-dark font-weight-bold font-size-24">
            <span className="mr-3">Mediatec Software</span>
            <span className="align-middle text-primary text-uppercase font-size-12 badge badge-light">
              Default
            </span>
          </div>
          <div>
            Operational / Developer tooling / Update your project information under Settings
          </div>
        </div>
        <a
          className="btn btn-light btn-lg text-blue font-size-14"
          href="#"
          onClick={e => e.preventDefault()}
        >
          Move Resources →
        </a>
      </div>
      <Tabs className="kit-tabs-bordered mb-3" defaultActiveKey="1">
        <TabPane tab="Resources" key="1" />
        <TabPane tab="Activity" key="2" />
        <TabPane tab="Settings" key="3" />
      </Tabs>
      <div className="row mb-4">
        <div className="col-xl-3 col-lg-6">
          <div className="card">
            <div className="card-body">
              <Chart4 />
            </div>
          </div>
        </div>
        <div className="col-xl-3 col-lg-6">
          <div className="card">
            <div className="card-body">
              <Chart4v1 />
            </div>
          </div>
        </div>
        <div className="col-xl-3 col-lg-6">
          <div className="card">
            <div className="card-body">
              <Chart4v2 />
            </div>
          </div>
        </div>
        <div className="col-xl-3 col-lg-6">
          <div className="card">
            <div className="card-body">
              <Chart4v3 />
            </div>
          </div>
        </div>
      </div>
      <h6 className="mb-4 text-uppercase">
        <strong>Droplets (2)</strong>
      </h6>
      <div className="mb-5">
        <div className="card mb-3">
          <div className="card-body py-2">
            <div className="d-flex align-items-center flex-wrap">
              <div className="d-flex flex-nowrap align-items-center width-200 flex-shrink-1 mr-2">
                <div className="kit__utils__donut kit__utils__donut--danger mr-2 flex-shrink-0" />
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className="font-weight-bold text-blue text-nowrap"
                >
                  mediatec-main-server
                </a>
              </div>
              <div className="flex-grow-1 mr-2">FRA1 / 1GB / 25GB Disk</div>
              <div className="flex-grow-1 mr-auto">46.101.103.230</div>
              <div className="dropdown d-inline-block">
                <Dropdown overlay={dropdownMenu} placement="bottomRight">
                  <button
                    type="button"
                    className="btn btn-light dropdown-toggle dropdown-toggle-noarrow"
                  >
                    <i className="fe fe-more-vertical" />
                  </button>
                </Dropdown>
              </div>
            </div>
          </div>
        </div>
        <div className="card mb-3">
          <div className="card-body py-2">
            <div className="d-flex align-items-center flex-wrap">
              <div className="d-flex flex-nowrap align-items-center width-200 flex-shrink-1 mr-2">
                <div className="kit__utils__donut kit__utils__donut--success mr-2 flex-shrink-0" />
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className="font-weight-bold text-blue text-nowrap"
                >
                  mediatec-cdn
                </a>
              </div>
              <div className="flex-grow-1 mr-2">FRA1 / 1GB / 25GB Disk</div>
              <div className="flex-grow-1 mr-auto">46.101.103.230</div>
              <div className="dropdown d-inline-block">
                <Dropdown overlay={dropdownMenu} placement="bottomRight">
                  <button
                    type="button"
                    className="btn btn-light dropdown-toggle dropdown-toggle-noarrow"
                  >
                    <i className="fe fe-more-vertical" />
                  </button>
                </Dropdown>
              </div>
            </div>
          </div>
        </div>
      </div>
      <h6 className="mb-4 text-uppercase">
        <strong>Create something new</strong>
      </h6>
      <div className="row">
        <div className="col-lg-4">
          <List5 />
        </div>
        <div className="col-lg-4">
          <List5 />
        </div>
        <div className="col-lg-4">
          <List5 />
        </div>
      </div>
    </div>
  )
}

export default ExtraAppsDigitaloceanDroplets
