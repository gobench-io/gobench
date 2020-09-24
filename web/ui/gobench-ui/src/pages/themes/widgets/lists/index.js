import React from 'react'
import { Helmet } from 'react-helmet'

import List1 from 'components/kit/widgets/Lists/1'
import List2 from 'components/kit/widgets/Lists/2'
import List3 from 'components/kit/widgets/Lists/3'
import List4 from 'components/kit/widgets/Lists/4'
import List5 from 'components/kit/widgets/Lists/5'
import List6 from 'components/kit/widgets/Lists/6'
import List7 from 'components/kit/widgets/Lists/7'
import List8 from 'components/kit/widgets/Lists/8'
import List9 from 'components/kit/widgets/Lists/9'
import List10 from 'components/kit/widgets/Lists/10'
import List11 from 'components/kit/widgets/Lists/11'
import List12 from 'components/kit/widgets/Lists/12'
import List13 from 'components/kit/widgets/Lists/13'
import List14 from 'components/kit/widgets/Lists/14'
import List15 from 'components/kit/widgets/Lists/15'
import List16 from 'components/kit/widgets/Lists/16'
import List17 from 'components/kit/widgets/Lists/17'
import List18 from 'components/kit/widgets/Lists/18'
import List19 from 'components/kit/widgets/Lists/19'
import List20 from 'components/kit/widgets/Lists/20'
import List21 from 'components/kit/widgets/Lists/21'
import List21v1 from 'components/kit/widgets/Lists/21v1'
import List21v2 from 'components/kit/widgets/Lists/21v2'

const Widgets = () => {
  return (
    <div>
      <Helmet title="Widgets / Lists" />
      <div className="row">
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">List / 1</h2>
            <div className="card">
              <div className="card-body">
                <List1 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 5</h2>
            <div className="card">
              <div className="card-body">
                <List5 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 8</h2>
            <div className="card">
              <div className="card-body">
                <List8 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 11</h2>
            <div className="card">
              <div className="card-body">
                <List11 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 14</h2>
            <div className="card">
              <div className="card-body">
                <List14 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 17</h2>
            <div className="card">
              <div className="card-body">
                <List17 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 20</h2>
            <div className="card">
              <div className="card-body">
                <List20 />
              </div>
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">List / 2</h2>
            <div className="card">
              <List2 />
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 6</h2>
            <div className="card">
              <div className="card-body">
                <List6 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 9</h2>
            <div className="card">
              <div className="card-body">
                <List9 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 12</h2>
            <div className="card">
              <div className="card-body">
                <List12 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 15</h2>
            <div className="card">
              <div className="card-body">
                <List15 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 19</h2>
            <div className="card">
              <div className="card-body">
                <List19 />
              </div>
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div>
            <h2 className="badge-example">List / 3</h2>
            <div className="card">
              <div className="card-body">
                <List3 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 4</h2>
            <div className="card">
              <div className="card-body">
                <List4 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 7</h2>
            <div className="card">
              <div className="card-body">
                <List7 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 10</h2>
            <div className="card">
              <div className="card-body">
                <List10 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 13</h2>
            <div className="card">
              <div className="card-body">
                <List13 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 16</h2>
            <div className="card">
              <div className="card-body">
                <List16 />
              </div>
            </div>
          </div>
          <div>
            <h2 className="badge-example">List / 18</h2>
            <div className="card">
              <div className="card-body">
                <List18 />
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="row">
        <div className="col-lg-4">
          <h2 className="badge-example">List / 21</h2>
          <List21 />
        </div>
        <div className="col-lg-4">
          <h2 className="badge-example">List / 21-1</h2>
          <List21v1 />
        </div>
        <div className="col-lg-4">
          <h2 className="badge-example">List / 22-2</h2>
          <List21v2 />
        </div>
      </div>
    </div>
  )
}

export default Widgets
