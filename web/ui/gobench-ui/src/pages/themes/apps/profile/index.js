import React, { useState } from 'react'
import { Helmet } from 'react-helmet'
import { Tabs, Input, Button, Upload, Form } from 'antd'
import General1 from 'components/kit/widgets/General/1'
import General10v1 from 'components/kit/widgets/General/10v1'
import General12v1 from 'components/kit/widgets/General/12v1'
import General14 from 'components/kit/widgets/General/14'
import General15 from 'components/kit/widgets/General/15'
import List19 from 'components/kit/widgets/Lists/19'

const { TabPane } = Tabs

const AppsProfile = () => {
  const [tabKey, setTabKey] = useState('1')

  const changeTab = key => {
    setTabKey(key)
  }

  return (
    <div>
      <Helmet title="Profile" />
      <div className="row">
        <div className="col-xl-4 col-lg-12">
          <div className="card">
            <div className="card-body">
              <General10v1 />
            </div>
          </div>
          <div className="card text-white bg-primary">
            <General12v1 />
          </div>
          <div className="card">
            <div className="card-body">
              <General1 />
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <List19 />
            </div>
          </div>
        </div>
        <div className="col-xl-8 col-lg-12">
          <div className="card">
            <div className="card-header card-header-flex flex-column">
              <div className="d-flex flex-wrap border-bottom pt-3 pb-4 mb-3">
                <div className="mr-5">
                  <div className="text-dark font-size-18 font-weight-bold">David Beckham</div>
                  <div className="text-gray-6">@david100</div>
                </div>
                <div className="mr-5 text-center">
                  <div className="text-dark font-size-18 font-weight-bold">100</div>
                  <div className="text-gray-6">Posts</div>
                </div>
                <div className="mr-5 text-center">
                  <div className="text-dark font-size-18 font-weight-bold">17,256</div>
                  <div className="text-gray-6">Followers</div>
                </div>
              </div>
              <Tabs activeKey={tabKey} className="mr-auto kit-tabs-bold" onChange={changeTab}>
                <TabPane tab="Agent Wall" key="1" />
                <TabPane tab="Messages" key="2" />
                <TabPane tab="Settings" key="3" />
              </Tabs>
            </div>
            <div className="card-body">
              {tabKey === '1' && (
                <div>
                  <General15 />
                  <General15 />
                </div>
              )}
              {tabKey === '2' && <General14 />}
              {tabKey === '3' && (
                <Form layout="vertical" className="login-form">
                  <h5 className="text-black mt-4">
                    <strong>Personal Information</strong>
                  </h5>
                  <div className="row">
                    <div className="col-lg-6">
                      <Form.Item name="userName" label="Username">
                        <Input placeholder="Username" />
                      </Form.Item>
                    </div>
                    <div className="col-lg-6">
                      <Form.Item name="email" label="Email">
                        <Input placeholder="Email" />
                      </Form.Item>
                    </div>
                  </div>
                  <h5 className="text-black mt-4">
                    <strong>New Password</strong>
                  </h5>
                  <div className="row">
                    <div className="col-lg-6">
                      <Form.Item label="Password">
                        <Input placeholder="New password" />
                      </Form.Item>
                    </div>
                    <div className="col-lg-6">
                      <Form.Item name="confirmpassword" label="Confirm Password">
                        <Input placeholder="Confirm password" />
                      </Form.Item>
                    </div>
                  </div>
                  <div className="row">
                    <div className="col-lg-6">
                      <h5 className="text-black mt-4 mb-3">
                        <strong>Profile Avatar</strong>
                      </h5>
                      <Upload>
                        <Button>
                          <i className="fe fe-upload mr-2" /> Click to Upload
                        </Button>
                      </Upload>
                    </div>
                    <div className="col-lg-6">
                      <h5 className="text-black mt-4 mb-3">
                        <strong>Profile Background</strong>
                      </h5>
                      <Upload>
                        <Button>
                          <i className="fe fe-upload mr-2" /> Click to Upload
                        </Button>
                      </Upload>
                    </div>
                  </div>
                  <div className="form-actions">
                    <Button
                      style={{ width: 200 }}
                      type="primary"
                      htmlType="submit"
                      className="mr-3"
                    >
                      Submit
                    </Button>
                    <Button htmlType="submit">Cancel</Button>
                  </div>
                </Form>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default AppsProfile
