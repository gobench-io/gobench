import React from 'react'
import { Helmet } from 'react-helmet'
import TablesAntBasic from './examples/basic'
// import TablesAntdRowSelection from './examples/row-selection'
// import TablesAntdFilterSorter from './examples/filter-sorter'
// import TablesAntdCustomFilter from './examples/custom-filter'
// import TablesAntdExpandableRow from './examples/expandable-row'
// import TablesAntdFixed from './examples/fixed-header-columns'
// import TablesAntdResizable from './examples/resizable'

const TablesAntd = () => {
  return (
    <div>
      <Helmet title="Tables / Antd" />
      <div className="kit__utils__heading">
        <h5>
          <span className="mr-3">Antd Tables</span>
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
            <strong>Basic Usage</strong>
          </h5>
          <TablesAntBasic />
          {/* <h5 className="mb-4">
            <strong>Row Selection</strong>
          </h5>
          <TablesAntdRowSelection />
          <h5 className="mb-4">
            <strong>Filter and Sorter</strong>
          </h5>
          <TablesAntdFilterSorter />
          <h5 className="mb-4">
            <strong>Custom Filter Panel</strong>
          </h5>
          <TablesAntdCustomFilter />
          <h5 className="mb-4">
            <strong>Expandable Row</strong>
          </h5>
          <TablesAntdExpandableRow />
          <h5 className="mb-4">
            <strong>Fixed Header and Columns</strong>
          </h5>
          <TablesAntdFixed />
          <h5 className="mb-4">
            <strong>Resizable column</strong>
          </h5>
          <TablesAntdResizable /> */}
        </div>
      </div>
    </div>
  )
}

export default TablesAntd
