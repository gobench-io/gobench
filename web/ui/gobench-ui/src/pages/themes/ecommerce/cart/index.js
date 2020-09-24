import React from 'react'
import { Helmet } from 'react-helmet'
import { CreditCardOutlined } from '@ant-design/icons'
import { Table, InputNumber, Input, Form } from 'antd'
import { ordersTableData } from './data.json'

const EcommerceCart = () => {
  const columns = [
    {
      title: '#',
      dataIndex: 'number',
      width: 50,
    },
    {
      title: 'Description',
      dataIndex: 'description',
      render: text => (
        <a className="btn btn-sm btn-light" href="#" onClick={e => e.preventDefault()}>
          {text}
        </a>
      ),
    },
    {
      title: 'Quantity',
      dataIndex: 'quantity',
      render: text => <InputNumber defaultValue={text} size="small" />,
      className: 'text-right',
    },
    {
      title: 'Unit Cost',
      dataIndex: 'unitcost',
      className: 'text-right',
    },
    {
      title: 'Total',
      dataIndex: 'total',
      className: 'text-right',
    },
    {
      title: '',
      dataIndex: '',
      render: () => (
        <a href="#" onClick={e => e.preventDefault()} className="btn btn-sm btn-light">
          <i className="fe fe-trash mr-1" /> Remove
        </a>
      ),
      className: 'text-right',
    },
  ]

  return (
    <div>
      <Helmet title="Ecommerce: Cart" />
      <div className="cui__utils__heading">
        <strong>Ecommerce: Cart</strong>
      </div>
      <div className="card">
        <div className="card-body">
          <h6 className="mb-4 text-uppercase">
            <strong>Order items</strong>
          </h6>
          <div className="text-nowrap mb-4">
            <Table columns={columns} dataSource={ordersTableData} pagination={false} />
          </div>
          <h6 className="mb-4 text-uppercase">
            <strong>Shipment details</strong>
          </h6>
          <div className="row mb-4">
            <div className="col-md-8">
              <Form layout="vertical">
                <div className="row">
                  <div className="col-md-6">
                    <div className="form-group">
                      <Form.Item
                        label="Email"
                        name="email"
                        rules={[{ required: true, message: 'Please input your Email!' }]}
                      >
                        <Input id="checkout-email" placeholder="Email" />
                      </Form.Item>
                    </div>
                  </div>
                  <div className="col-md-6">
                    <div className="form-group">
                      <Form.Item
                        label="Phone Number"
                        name="phoneNumber"
                        rules={[{ required: true, message: 'Please input your Phone Number!' }]}
                      >
                        <Input id="checkout-phnum" placeholder="Phone Number" />
                      </Form.Item>
                    </div>
                  </div>
                </div>
                <div className="row">
                  <div className="col-md-6">
                    <div className="form-group">
                      <Form.Item
                        label="Name"
                        name="name"
                        rules={[{ required: true, message: 'Please input your Name!' }]}
                      >
                        <Input id="checkout-name" placeholder="Name" />
                      </Form.Item>
                    </div>
                  </div>
                  <div className="col-md-6">
                    <div className="form-group">
                      <Form.Item
                        label="Surname"
                        name="surname"
                        rules={[{ required: true, message: 'Please input your Surname!' }]}
                      >
                        <Input id="checkout-surname" placeholder="Surname" />
                      </Form.Item>
                    </div>
                  </div>
                </div>
                <div className="form-group">
                  <Form.Item
                    label="City"
                    name="city"
                    rules={[{ required: true, message: 'Please input your City!' }]}
                  >
                    <Input id="checkout-city" placeholder="City" />
                  </Form.Item>
                </div>
                <div className="form-group">
                  <Form.Item
                    label="Address"
                    name="address"
                    rules={[{ required: true, message: 'Please input your Address!' }]}
                  >
                    <Input id="checkout-adress" placeholder="Adress" className="mb-3" />
                  </Form.Item>
                </div>
                <h6 className="mb-4 text-uppercase">
                  <strong>Shipment details</strong>
                </h6>
                <div className="form-group">
                  <Form.Item
                    label="Card Number"
                    name="cardnum"
                    rules={[{ required: true, message: 'Please input Card Number!' }]}
                  >
                    <Input addonBefore={<CreditCardOutlined />} placeholder="Card Number" />
                  </Form.Item>
                </div>
                <div className="row">
                  <div className="col-md-7">
                    <div className="form-group">
                      <Form.Item
                        label="Expiration Date"
                        name="expirationdate"
                        rules={[{ required: true, message: 'Please input Card Expiration Date!' }]}
                      >
                        <Input id="checkout-cardexpdate" placeholder="MM / YY" />
                      </Form.Item>
                    </div>
                  </div>
                  <div className="col-md-5 pull-right">
                    <div className="form-group">
                      <Form.Item
                        label="Card CVC"
                        name="cardcvc"
                        rules={[{ required: true, message: 'Please input Card CVC!' }]}
                      >
                        <Input id="checkout-cardholder" placeholder="CVC" />
                      </Form.Item>
                    </div>
                  </div>
                </div>
                <div className="form-group">
                  <Form.Item
                    label="Card Holder Name"
                    name="cardholdername"
                    rules={[{ required: true, message: 'Please input Card Holder Name!' }]}
                  >
                    <Input id="checkout-cardholder" placeholder="Name and Surname" />
                  </Form.Item>
                </div>
              </Form>
            </div>
            <div className="col-md-4">
              <h4 className="text-black mb-3">
                <strong>General Info</strong>
              </h4>
              <h2>
                <i className="fa fa-cc-visa text-primary mr-1" />
                <i className="fa fa-cc-mastercard text-default mr-1" />
                <i className="fa fa-cc-amex text-default" />
              </h2>
              <p>
                Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor
                incididunt ut labore et dolore magna aliqua.
              </p>
              <p>
                Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip
                ex ea commodo consequat.
              </p>
            </div>
          </div>
          <div className="border-top text-dark font-size-18 pt-4 text-right">
            <p className="mb-1">
              Sub - Total amount: <strong className="font-size-24">$5,700.00</strong>
            </p>
            <p className="mb-1">
              VAT: <strong className="font-size-24">$57.00</strong>
            </p>
            <p className="mb-4">
              Grand Total: <strong className="font-size-36">$5,757.00</strong>
            </p>
            <a
              href="#"
              onClick={e => e.preventDefault()}
              className="btn btn-lg btn-success width-200 mb-2"
            >
              Order Now
            </a>
          </div>
        </div>
      </div>
    </div>
  )
}

export default EcommerceCart
