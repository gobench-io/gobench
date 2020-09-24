import React from 'react'
import { Helmet } from 'react-helmet'
import { Tabs, Dropdown, Menu } from 'antd'

const { TabPane } = Tabs

const dropdownMenu = (
  <Menu>
    <Menu.Item>
      <a>Action</a>
    </Menu.Item>
    <Menu.Item>
      <a>Another action</a>
    </Menu.Item>
    <Menu.Item>
      <a>Something else here</a>
    </Menu.Item>
    <div className="dropdown-divider" />
    <Menu.Item>
      <a>Separated link</a>
    </Menu.Item>
  </Menu>
)

const CardsTabbedCards = () => {
  return (
    <div>
      <Helmet title="Cards / Tabbed" />
      <div className="kit__utils__heading">
        <h5>
          <span className="mr-3">Tabbed Cards</span>
          <a
            href="https://getbootstrap.com/docs/4.3/components/card/"
            target="_blank"
            rel="noopener noreferrer"
            className="btn btn-sm btn-light"
          >
            Official Documentation
            <i className="fe fe-corner-right-up" />
          </a>
        </h5>
      </div>
      <div className="row">
        <div className="col-xl-6 col-lg-12">
          <div className="card">
            <div className="card-header card-header-flex">
              <div className="d-flex flex-column justify-content-center mr-auto">
                <h5 className="mb-0">Basic</h5>
              </div>
              <Tabs defaultActiveKey="1" className="kit-tabs">
                <TabPane tab="History" key="1" />
                <TabPane tab="Information" key="2" />
                <TabPane tab="Actions" key="3" />
              </Tabs>
            </div>
            <div className="card-body">
              Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum
              has been the industry&apos;s standard dummy text ever since the 1500s, when an unknown
              printer took a galley of type and scrambled it to make a type specimen book. It has
              survived not only five centuries, but also the leap into electronic typesetting,
              remaining essentially unchanged.
            </div>
          </div>
          <div className="card">
            <div className="card-header card-header-flex">
              <div className="d-flex flex-column justify-content-center mr-auto">
                <h5 className="mb-0">With Dropdown</h5>
              </div>
              <Tabs defaultActiveKey="1" className="kit-tabs">
                <TabPane tab="Messages" key="1" />
              </Tabs>
              <div className="d-inline-flex align-items-center">
                <Dropdown overlay={dropdownMenu} placement="bottomRight">
                  <a className="nav-link dropdown-toggle" role="button">
                    Dropdown
                  </a>
                </Dropdown>
              </div>
            </div>
            <div className="card-body">
              Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum
              has been the industry&apos;s standard dummy text ever since the 1500s, when an unknown
              printer took a galley of type and scrambled it to make a type specimen book. It has
              survived not only five centuries, but also the leap into electronic typesetting,
              remaining essentially unchanged.
            </div>
          </div>
          <div className="card">
            <div className="card-header card-header-flex">
              <div className="d-flex flex-column justify-content-center mr-auto">
                <h5 className="mb-0">Bold Border</h5>
              </div>
              <Tabs defaultActiveKey="1" className="kit-tabs-bold">
                <TabPane tab="History" key="1" />
                <TabPane tab="Actions" key="2" />
              </Tabs>
              <div className="d-inline-flex align-items-center">
                <Dropdown overlay={dropdownMenu} placement="bottomRight">
                  <a className="nav-link dropdown-toggle" role="button">
                    Dropdown
                  </a>
                </Dropdown>
              </div>
            </div>
            <div className="card-body">
              Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum
              has been the industry&apos;s standard dummy text ever since the 1500s, when an unknown
              printer took a galley of type and scrambled it to make a type specimen book. It has
              survived not only five centuries, but also the leap into electronic typesetting,
              remaining essentially unchanged.
            </div>
          </div>
        </div>
        <div className="col-xl-6 col-lg-12">
          <div className="card">
            <div className="card-header card-header-flex">
              <div className="d-flex flex-column justify-content-center mr-auto">
                <h5 className="mb-0">Pills</h5>
              </div>
              <Tabs defaultActiveKey="1" className="kit-tabs-pills">
                <TabPane tab={<a className="nav-link">History</a>} key="1" />
                <TabPane tab={<a className="nav-link">Information</a>} key="2" />
                <TabPane tab={<a className="nav-link">Actions</a>} key="3" />
              </Tabs>
            </div>
            <div className="card-body">
              Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum
              has been the industry&apos;s standard dummy text ever since the 1500s, when an unknown
              printer took a galley of type and scrambled it to make a type specimen book. It has
              survived not only five centuries, but also the leap into electronic typesetting,
              remaining essentially unchanged.
            </div>
          </div>
          <div className="card">
            <div className="card-header card-header-flex">
              <div className="d-flex flex-row align-items-center mr-auto">
                <h5 className="mb-0">
                  <i className="fe fe-activity mr-1 font-size-18 text-muted" />
                  With Icon
                </h5>
              </div>
              <Tabs defaultActiveKey="1" className="kit-tabs-bold">
                <TabPane
                  tab={
                    <span>
                      <i className="fe fe-align-left mr-1" />
                      Left
                    </span>
                  }
                  key="1"
                />
                <TabPane
                  tab={
                    <span>
                      <i className="fe fe-align-center mr-1" />
                      Center
                    </span>
                  }
                  key="2"
                />
                <TabPane
                  tab={
                    <span>
                      <i className="fe fe-align-right mr-1" />
                      Right
                    </span>
                  }
                  key="3"
                />
              </Tabs>
            </div>
            <div className="card-body">
              Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum
              has been the industry&apos;s standard dummy text ever since the 1500s, when an unknown
              printer took a galley of type and scrambled it to make a type specimen book. It has
              survived not only five centuries, but also the leap into electronic typesetting,
              remaining essentially unchanged.
            </div>
          </div>
          <div className="card">
            <div className="card-header card-header-flex">
              <Tabs defaultActiveKey="1" className="kit-tabs-bold">
                <TabPane tab="Notifications" key="1" />
              </Tabs>
              <div className="d-inline-flex align-items-center">
                <Dropdown overlay={dropdownMenu} placement="bottomRight">
                  <a className="nav-link dropdown-toggle" role="button">
                    Dropdown
                  </a>
                </Dropdown>
              </div>
            </div>
            <div className="card-body">
              Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum
              has been the industry&apos;s standard dummy text ever since the 1500s, when an unknown
              printer took a galley of type and scrambled it to make a type specimen book. It has
              survived not only five centuries, but also the leap into electronic typesetting,
              remaining essentially unchanged.
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default CardsTabbedCards
