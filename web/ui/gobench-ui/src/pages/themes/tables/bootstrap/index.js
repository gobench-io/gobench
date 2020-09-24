import React from 'react'
import { Helmet } from 'react-helmet'
import TablesBootstrapBasic from './examples/basic'
import TablesBootstrapBordered from './examples/bordered'
import TablesBootstrapBorderless from './examples/borderless'
import TablesBootstrapDark from './examples/dark'
import TablesBootstrapHover from './examples/hover'
import TablesBootstrapResponsive from './examples/responsive'
import TablesBootstrapSmall from './examples/small'
import TablesBootstrapStriped from './examples/striped'

const TablesBootstrap = () => {
  return (
    <div>
      <Helmet title="Tables / Bootstrap" />
      <div className="kit__utils__heading">
        <h5>
          <span className="mr-3">Bootstrap Tables</span>
          <a
            href="https://ant.design/components/table"
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
          <h5 className="mb-4">
            <strong>Basic</strong>
          </h5>
          <TablesBootstrapBasic />
          <h5 className="mb-4">
            <strong>Dark table</strong>
          </h5>
          <TablesBootstrapDark />
          <h5 className="mb-4">
            <strong>Striped rows</strong>
          </h5>
          <TablesBootstrapStriped />
          <h5 className="mb-4">
            <strong>Bordered table</strong>
          </h5>
          <TablesBootstrapBordered />
          <h5 className="mb-4">
            <strong>Borderless table</strong>
          </h5>
          <TablesBootstrapBorderless />
          <h5 className="mb-4">
            <strong>Hoverable rows</strong>
          </h5>
          <TablesBootstrapHover />
          <h5 className="mb-4">
            <strong>Small table</strong>
          </h5>
          <TablesBootstrapSmall />
          <h5 className="mb-4">
            <strong>Responsive table</strong>
          </h5>
          <TablesBootstrapResponsive />
        </div>
      </div>
    </div>
  )
}

export default TablesBootstrap
