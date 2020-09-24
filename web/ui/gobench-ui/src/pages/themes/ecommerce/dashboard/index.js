import React from 'react'
import { Tabs } from 'antd'
import { Helmet } from 'react-helmet'
import Chart3 from 'components/kit/widgets/Charts/3'
import General2 from 'components/kit/widgets/General/2'
import General2v1 from 'components/kit/widgets/General/2v1'
import General2v2 from 'components/kit/widgets/General/2v2'
import General2v3 from 'components/kit/widgets/General/2v3'
import List11 from 'components/kit/widgets/Lists/11'
import List12 from 'components/kit/widgets/Lists/12'
import General16 from 'components/kit/widgets/General/16'
import productsData from './data.json'

const { TabPane } = Tabs

const EcommerceDashboard = () => {
  return (
    <div>
      <Helmet title="Ecommerce: Dashboard" />
      <div className="cui__utils__heading">
        <strong>Ecommerce: Dashboard</strong>
      </div>
      <div className="row">
        <div className="col-xl-8 col-lg-12">
          <div className="card">
            <Tabs className="kit-tabs-bordered pt-2" defaultActiveKey="2">
              <TabPane tab="Orders" key="1" />
              <TabPane tab="Revenue" key="2" />
            </Tabs>
            <div className="card-body">
              <Chart3 />
            </div>
          </div>
          <div className="card">
            <Tabs className="kit-tabs-bordered pt-2" defaultActiveKey="2">
              <TabPane tab="Bestsellers" key="1" />
              <TabPane tab="Most Viewed" key="2" />
              <TabPane tab="Highest Rated" key="3" />
            </Tabs>
            <div className="card-body">
              <div className="row">
                {productsData.map(product => {
                  const { isNew, isFavorite, image, name, price, oldPrice } = product
                  return (
                    <div className="col-lg-6" key={Math.random()}>
                      <General16
                        isNew={isNew}
                        isFavorite={isFavorite}
                        image={image}
                        name={name}
                        price={price}
                        oldPrice={oldPrice}
                      />
                    </div>
                  )
                })}
              </div>
            </div>
          </div>
        </div>
        <div className="col-xl-4 col-lg-12">
          <div className="card text-white bg-success">
            <div className="card-body">
              <General2v3 />
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <General2 />
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <General2v1 />
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <General2v2 />
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <List11 />
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <List12 />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default EcommerceDashboard
