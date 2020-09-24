import React, { useState } from 'react'
import { Helmet } from 'react-helmet'
import { Select, Tabs } from 'antd'
import General16 from 'components/kit/widgets/General/16'
import style from './style.module.scss'
import data from './data.json'

const { Option } = Select
const { TabPane } = Tabs

const EcommerceProductDetails = () => {
  const [activeImgIndex, setActiveImgIndex] = useState(0)
  const [isFavourite, setIsFavourite] = useState(false)

  const setFavourite = e => {
    e.preventDefault()
    setIsFavourite(!isFavourite)
  }

  const setActiveImg = (e, index) => {
    e.preventDefault()
    setActiveImgIndex(index)
  }

  return (
    <div>
      <Helmet title="Ecommerce: Product Details" />
      <div className="cui__utils__heading">
        <strong>Ecommerce: Product Details</strong>
      </div>
      <div className="card overflow-hidden">
        <div className={style.new}>New</div>
        <div className="card-body">
          <div className="row">
            <div className="col-lg-4">
              <a
                className={`${style.favourite} ${isFavourite ? 'text-dark' : 'text-gray-3'}`}
                href="#"
                onClick={setFavourite}
              >
                <i className="fe fe-heart font-size-21" />
              </a>
              <div className={`${style.image} height-250 mb-3`}>
                <img className="img-fluid" src={data.descr.images[activeImgIndex]} alt="Product" />
              </div>
              <div className="d-flex flex-wrap mb-3">
                {data.descr.images.map((image, index) => (
                  <a
                    href="#"
                    className={`${index === activeImgIndex ? 'border-primary' : ''} ${
                      style.thumb
                    } width-100 height-100 border m-2`}
                    onClick={e => {
                      setActiveImg(e, index)
                    }}
                    key={image}
                  >
                    <img className="img-fluid" src={image} alt="Product" />
                  </a>
                ))}
              </div>
            </div>
            <div className="col-lg-8">
              <div className="font-size-24 font-weight-bold text-dark mb-2">
                $199.28 <del className="align-text-top font-size-14">$299.28</del>
              </div>
              <div className="pb-3 mb-4 border-bottom">
                <a href="" className="text-blue font-size-18">
                  TP-Link AC1750 Smart WiFi Router - Dual
                  <br />
                  Band Gigabit Wireless Internet Routers for
                  <br />
                  Home
                </a>
              </div>
              <div className="mb-4 width-300">
                <Select defaultValue="Red" style={{ width: 300 }}>
                  <Option value="red">Red</Option>
                  <Option value="black">Black</Option>
                  <Option value="cyan">Cyan</Option>
                  <Option value="blue">Blue</Option>
                </Select>
              </div>
              <a
                href="#"
                onClick={e => e.preventDefault()}
                className="width-200 btn btn-success btn-with-addon mr-auto mb-3 text-nowrap d-none d-md-block"
              >
                <span className="btn-addon">
                  <i className="btn-addon-icon fe fe-plus-circle" />
                </span>
                Add To Card
              </a>
              <Tabs defaultActiveKey="1" className="kit-tabs-bordered">
                <TabPane tab="Information" key="1" />
                <TabPane tab="Description" key="2" />
              </Tabs>
              <div className="card-body px-0">
                <p>
                  Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem
                  Ipsum has been the industry&apos;s standard dummy text ever since the 1500s, when
                  an unknown printer took a galley of type and scrambled it to make a type specimen
                  book. It has survived not only five centuries, but also the leap into electronic
                  typesetting, remaining essentially unchanged.
                </p>
                <p>
                  Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem
                  Ipsum has been the industry&apos;s standard dummy text ever since the 1500s, when
                  an unknown printer took a galley of type.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="cui__utils__heading">Related Products</div>
      <div className="row">
        {data.products.map(product => {
          const { isNew, isFav, image, name, price, oldPrice } = product
          return (
            <div className="col-lg-4" key={Math.random()}>
              <General16
                isNew={isNew}
                isFavourite={isFav}
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
  )
}

export default EcommerceProductDetails
