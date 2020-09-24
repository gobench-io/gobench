import React from 'react'
import { Helmet } from 'react-helmet'
import { MailOutlined, UploadOutlined, UserOutlined } from '@ant-design/icons'
import { Input, Button, Upload, Form } from 'antd'
import General15 from 'components/kit/widgets/General/15'
import List15 from 'components/kit/widgets/Lists/15'

const { TextArea } = Input

const ExtraAppsWordpressPost = () => {
  return (
    <div>
      <Helmet title="Wordpress Post" />
      <div className="row">
        <div className="col-xl-9 col-lg-12">
          <div className="card">
            <div className="card-body">
              <div className="mb-2">
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className="text-dark font-size-24 font-weight-bold"
                >
                  [Feature Request] How to enable custom font that comes from svg #2460
                </a>
              </div>
              <div className="mb-3">
                <a className="font-weight-bold" href="#" onClick={e => e.preventDefault()}>
                  zxs2162
                </a>{' '}
                wrote this post 12 days ago · 0 comments
              </div>
              <div className="mb-4">
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className="badge text-blue text-uppercase bg-light font-size-12 mr-2"
                >
                  Umi
                </a>
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className="badge text-blue text-uppercase bg-light font-size-12 mr-2"
                >
                  React-framework
                </a>
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className="badge text-blue text-uppercase bg-light font-size-12 mr-2"
                >
                  Umijs
                </a>
              </div>
              <div>
                <img
                  className="img-fluid mb-4"
                  src="resources/images/content/photos/1.jpeg"
                  alt="Sea"
                />
                <p>
                  Lorem ipsum dolor sit amet, consectetur adipisicing elit. Nihil laborum est
                  perferendis consectetur corporis esse labore minima molestias, exercitationem
                  consequuntur! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Nihil
                  laborum est perferendis consectetur corporis esse labore minima molestias,
                  exercitationem consequuntur! Lorem ipsum dolor sit amet, consectetur adipisicing
                  elit. Nihil laborum est perferendis consectetur corporis esse labore minima
                  molestias, exercitationem consequuntur! Lorem ipsum dolor sit amet, consectetur
                  adipisicing elit. Nihil laborum est perferendis consectetur corporis esse labore
                  minima molestias, exercitationem consequuntur!
                </p>
                <p>
                  Lorem ipsum dolor sit amet, consectetur adipisicing elit. Nihil laborum est
                  perferendis consectetur corporis esse labore minima molestias, exercitationem
                  consequuntur! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Nihil
                  laborum est perferendis consectetur corporis esse labore minima molestias,
                  exercitationem consequuntur!
                </p>
              </div>
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <h6 className="mb-4 text-uppercase">
                <strong>Comments (76)</strong>
              </h6>
              <General15 />
              <a
                href="#"
                onClick={e => e.preventDefault()}
                className="d-block btn btn-light text-primary mt-3"
              >
                Load More
              </a>
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <div className="d-flex align-items-center flex-wrap border-bottom mb-3 pb-3">
                <div className="kit__utils__avatar kit__utils__avatar--size110 mr-3 mb-3 align-items-center flex-shrink-0">
                  <img src="resources/images/avatars/5.jpg" alt="Mary Stanform" />
                </div>
                <div className="mb-3">
                  <div className="font-weight-bold font-size-16 text-dark mb-2">Trinity Parson</div>
                  <p className="font-italic">
                    “I hope you enjoy reading this as much as I enjoyed writing this.”
                  </p>
                  <a href="#" className="btn btn-sm btn-primary">
                    View Profile
                  </a>
                </div>
              </div>
              <h5 className="text-dark mb-4">Leave a comment</h5>
              <Form className="login-form">
                <Form.Item name="userName">
                  <Input
                    prefix={<UserOutlined style={{ color: 'rgba(0,0,0,.25)' }} />}
                    placeholder="Your name"
                  />
                </Form.Item>
                <Form.Item name="mail">
                  <Input
                    prefix={<MailOutlined style={{ color: 'rgba(0,0,0,.25)' }} />}
                    placeholder="Your email"
                  />
                </Form.Item>
                <Form.Item name="message">
                  <TextArea rows={3} placeholder="Your message" />
                </Form.Item>
                <Form.Item>
                  <Button className="mr-2" type="primary" style={{ width: 200 }}>
                    <i className="fa fa-send mr-2" />
                    Send
                  </Button>
                  <Upload>
                    <Button>
                      <UploadOutlined />
                      Attach File
                    </Button>
                  </Upload>
                </Form.Item>
              </Form>
            </div>
          </div>
        </div>
        <div className="col-xl-3 col-lg-12">
          <div className="pb-4 mb-3 border-bottom">
            <label className="font-weight-bold d-block" htmlFor="search-input">
              <span className="mb-2 d-inline-block">Search Post</span>
              <input
                className="form-control width-100p"
                type="text"
                placeholder="Search post..."
                id="search-input"
              />
            </label>
          </div>
          <div className="pb-4 mb-3 border-bottom">
            <label className="font-weight-bold d-block" htmlFor="subscribe-input">
              <span className="mb-2 d-inline-block">Subscribe</span>
              <input
                className="form-control width-100p"
                type="text"
                id="subscribe-input"
                placeholder="Enter your email..."
              />
            </label>
          </div>
          <div className="pb-4 mb-3 border-bottom">
            <div className="font-weight-bold mb-2">Categories</div>
            <div>
              <a
                href="#"
                onClick={e => e.preventDefault()}
                className="badge text-blue text-uppercase bg-light font-size-12 mr-2"
              >
                Umi
              </a>
              <a
                href="#"
                onClick={e => e.preventDefault()}
                className="badge text-blue text-uppercase bg-light font-size-12 mr-2"
              >
                React-framework
              </a>
              <a
                href="#"
                onClick={e => e.preventDefault()}
                className="badge text-blue text-uppercase bg-light font-size-12 mr-2"
              >
                Umijs
              </a>
            </div>
          </div>
          <div className="pb-4 mb-3 border-bottom">
            <div className="font-weight-bold mb-3">Latest Posts</div>
            <List15 />
          </div>
        </div>
      </div>
    </div>
  )
}

export default ExtraAppsWordpressPost
