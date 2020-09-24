import React from 'react'
import { Helmet } from 'react-helmet'

import Chart1 from 'components/kit/widgets/Charts/1'
import Chart2 from 'components/kit/widgets/Charts/2'
import Chart3 from 'components/kit/widgets/Charts/3'
import Chart4 from 'components/kit/widgets/Charts/4'
import Chart4v1 from 'components/kit/widgets/Charts/4v1'
import Chart4v2 from 'components/kit/widgets/Charts/4v2'
import Chart4v3 from 'components/kit/widgets/Charts/4v3'
import Chart5 from 'components/kit/widgets/Charts/5'
import Chart6 from 'components/kit/widgets/Charts/6'
import Chart7 from 'components/kit/widgets/Charts/7'
import Chart8 from 'components/kit/widgets/Charts/8'
import Chart9 from 'components/kit/widgets/Charts/9'
import Chart10 from 'components/kit/widgets/Charts/10'
import Chart11 from 'components/kit/widgets/Charts/11'
import Chart11v1 from 'components/kit/widgets/Charts/11v1'
import Chart11v2 from 'components/kit/widgets/Charts/11v2'
import Chart12 from 'components/kit/widgets/Charts/12'
import Chart12v1 from 'components/kit/widgets/Charts/12v1'
import Chart13 from 'components/kit/widgets/Charts/13'
import Chart13v1 from 'components/kit/widgets/Charts/13v1'
import Chart13v2 from 'components/kit/widgets/Charts/13v2'

const Widgets = () => {
  return (
    <div>
      <Helmet title="Widgets / Charts" />
      <div className="row">
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 1</h2>
            <div className="card">
              <div className="card-body">
                <Chart1 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">Chart / 2</h2>
            <div className="card">
              <Chart2 />
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 9</h2>
            <div className="card">
              <div className="card-body">
                <Chart9 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">Chart / 5</h2>
            <div className="card">
              <div className="card-body">
                <Chart5 />
              </div>
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 10</h2>
            <div className="card">
              <div className="card-body">
                <Chart10 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">Chart / 6</h2>
            <div className="card">
              <Chart6 />
            </div>
          </div>
        </div>
        <div className="col-xl-6 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 3</h2>
            <div className="card">
              <div className="card-body">
                <Chart3 />
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-lg-6">
              <div>
                <h2 className="badge-example">Chart / 4</h2>
                <div className="card">
                  <div className="card-body">
                    <Chart4 />
                  </div>
                </div>
              </div>
            </div>
            <div className="col-lg-6">
              <div>
                <h2 className="badge-example">Chart / 4-1</h2>
                <div className="card">
                  <div className="card-body">
                    <Chart4v1 />
                  </div>
                </div>
              </div>
            </div>
            <div className="col-lg-6">
              <div>
                <h2 className="badge-example">Chart / 4-2</h2>
                <div className="card">
                  <div className="card-body">
                    <Chart4v2 />
                  </div>
                </div>
              </div>
            </div>
            <div className="col-lg-6">
              <div>
                <h2 className="badge-example">Chart / 4-3</h2>
                <div className="card">
                  <div className="card-body">
                    <Chart4v3 />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="col-xl-6 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 7</h2>
            <div className="card">
              <div className="card-body">
                <Chart7 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">Chart / 8</h2>
            <div className="card">
              <div className="card-body">
                <Chart8 />
              </div>
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 11</h2>
            <div className="card">
              <Chart11 />
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 11-1</h2>
            <div className="card">
              <Chart11v1 />
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 11-2</h2>
            <div className="card">
              <Chart11v2 />
            </div>
          </div>
        </div>
        <div className="col-xl-6 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 12</h2>
            <div className="card">
              <div className="card-body">
                <Chart12 />
              </div>
            </div>
          </div>
        </div>
        <div className="col-xl-6 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 12-1</h2>
            <div className="card">
              <div className="card-body">
                <Chart12v1 />
              </div>
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 13</h2>
            <div className="card">
              <Chart13 />
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 13-1</h2>
            <div className="card">
              <Chart13v1 />
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">Chart / 13-2</h2>
            <div className="card">
              <Chart13v2 />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Widgets
