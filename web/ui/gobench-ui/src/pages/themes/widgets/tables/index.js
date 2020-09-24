import React from 'react'
import { Helmet } from 'react-helmet'

import Table1 from 'components/kit/widgets/Tables/1'
import Table2 from 'components/kit/widgets/Tables/2'
import Table3 from 'components/kit/widgets/Tables/3'
import Table4 from 'components/kit/widgets/Tables/4'
import Table5 from 'components/kit/widgets/Tables/5'
import Table6 from 'components/kit/widgets/Tables/6'
import Table7 from 'components/kit/widgets/Tables/7'
import Table8 from 'components/kit/widgets/Tables/8'

const Widgets = () => {
  return (
    <div>
      <Helmet title="Widgets / Tables" />
      <div className="row">
        <div className="col-xl-6 col-lg-12">
          <div>
            <h2 className="badge-example">Table / 1</h2>
            <div className="card">
              <div className="card-body">
                <Table1 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">Table / 2</h2>
            <div className="card">
              <div className="card-body">
                <Table2 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">Table / 3</h2>
            <div className="card">
              <Table3 />
            </div>
          </div>
        </div>
        <div className="col-xl-6 col-lg-12">
          <div>
            <h2 className="badge-example">Table / 4</h2>
            <div className="card">
              <div className="card-body">
                <Table4 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">Table / 5</h2>
            <div className="card">
              <div className="card-body">
                <Table5 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">Table / 6</h2>
            <div className="card">
              <div className="card-body">
                <Table6 />
              </div>
            </div>
          </div>
        </div>
        <div className="col-lg-12">
          <div>
            <h2 className="badge-example">Table / 7</h2>
            <div className="card">
              <div className="card-body">
                <Table7 />
              </div>
            </div>
          </div>
        </div>
        <div className="col-lg-12">
          <div>
            <h2 className="badge-example">Table / 8</h2>
            <div className="card">
              <div className="card-body">
                <Table8 />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Widgets
