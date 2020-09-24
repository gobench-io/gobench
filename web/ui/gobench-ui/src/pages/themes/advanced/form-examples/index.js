import React, { useState } from 'react'
import { Helmet } from 'react-helmet'
import { InboxOutlined } from '@ant-design/icons'
import { Input, Slider, Cascader, Upload, message, Checkbox, Select, Button, Form } from 'antd'

const { Dragger } = Upload
const { Option } = Select

const AdvancedFormExamples = () => {
  const [confirmDirty, setConfirmDirty] = useState(false)

  const handleConfirmBlur = e => {
    const { value } = e.target
    setConfirmDirty(confirmDirty || !!value)
  }

  const opts = {
    name: 'file',
    multiple: true,
    action: 'https://www.mocky.io/v2/5cc8019d300000980a055e76',
    onChange(info) {
      const { status } = info.file
      if (status !== 'uploading') {
        console.log(info.file, info.fileList)
      }
      if (status === 'done') {
        message.success(`${info.file.name} file uploaded successfully.`)
      } else if (status === 'error') {
        message.error(`${info.file.name} file upload failed.`)
      }
    },
  }

  const formItemLayout = {
    labelCol: {
      xs: { span: 24 },
      sm: { span: 4 },
    },
    wrapperCol: {
      xs: { span: 24 },
      sm: { span: 12 },
    },
  }

  const marks = {
    0: '0',
    10: '10',
    20: '20',
    30: '30',
    40: '40',
    50: '50',
    60: '60',
    70: '70',
    80: '80',
    90: '90',
    100: '100',
  }

  const residences = [
    {
      value: 'zhejiang',
      label: 'Zhejiang',
      children: [
        {
          value: 'hangzhou',
          label: 'Hangzhou',
          children: [
            {
              value: 'xihu',
              label: 'West Lake',
            },
          ],
        },
      ],
    },
    {
      value: 'jiangsu',
      label: 'Jiangsu',
      children: [
        {
          value: 'nanjing',
          label: 'Nanjing',
          children: [
            {
              value: 'zhonghuamen',
              label: 'Zhong Hua Men',
            },
          ],
        },
      ],
    },
  ]

  return (
    <div>
      <Helmet title="Advanced / Form Examples" />
      <div className="kit__utils__heading">
        <h5>Form Examples</h5>
      </div>
      <div className="card">
        <div className="card-body">
          <h5 className="mb-4">
            <strong>Basic Form</strong>
          </h5>
          <Form {...formItemLayout} labelAlign="left">
            <Form.Item name="fullname6" label="Fullname">
              <Input placeholder="Your Fullname..." />
            </Form.Item>
            <Form.Item name="email6" label="Your Email...">
              <Input placeholder="Your Email..." />
            </Form.Item>
            <Form.Item name="budget1" label="Budget" className="mb-3">
              <Input placeholder="Your Budget..." addonBefore="$" />
            </Form.Item>
            <Form.Item name="amount" label="Amount">
              <Slider marks={marks} />
            </Form.Item>
            <button type="submit" className="btn btn-success px-5">
              Send Data
            </button>
          </Form>
        </div>
      </div>
      <div className="card">
        <div className="card-body">
          <h5 className="mb-4">
            <strong>Inline Form</strong>
          </h5>
          <Form layout="inline">
            <Form.Item name="budget2" className="mb-1 mt-1">
              <Input placeholder="Coins Amount" addonBefore="$" addonAfter=".00" />
            </Form.Item>
            <Form.Item name="budget3" className="mb-1 mt-1">
              <Input placeholder="8 Digit Pin" addonBefore="$" />
            </Form.Item>
            <button type="button" className="btn btn-success mt-1 mb-1">
              Confirm Transaction
            </button>
          </Form>
        </div>
      </div>
      <div className="card">
        <div className="card-body">
          <h5 className="mb-4">
            <strong>Two Columns</strong>
          </h5>
          <Form layout="vertical">
            <div className="row">
              <div className="col-md-6">
                <Form.Item name="email3" label="E-mail">
                  <Input placeholder="Email" />
                </Form.Item>
              </div>
              <div className="col-md-6">
                <Form.Item name="pass3" label="Password">
                  <Input placeholder="Password" />
                </Form.Item>
              </div>
              <div className="col-12">
                <Form.Item name="address3-1" label="Adress">
                  <Input placeholder="1234 Main St." />
                </Form.Item>
                <Form.Item name="address3-2" label="Adress 2">
                  <Input placeholder="Apartment, studio, or floor" />
                </Form.Item>
              </div>
              <div className="col-md-6">
                <Form.Item name="city3" label="City">
                  <Input />
                </Form.Item>
              </div>
              <div className="col-md-4">
                <Form.Item name="state3" label="State">
                  <Cascader options={residences} />
                </Form.Item>
              </div>
              <div className="col-md-2">
                <Form.Item name="zip" label="Zip">
                  <Input />
                </Form.Item>
              </div>
              <div className="col-12">
                <Form.Item valuePropName="fileList" name="upload3" label="Upload Presentation">
                  <Dragger {...opts}>
                    <p className="ant-upload-drag-icon">
                      <InboxOutlined />
                    </p>
                    <p className="ant-upload-text">Click or drag file to this area to upload</p>
                    <p className="ant-upload-hint">
                      Support for a single or bulk upload. Strictly prohibit from uploading company
                      data or other band files
                    </p>
                  </Dragger>
                </Form.Item>
              </div>
              <div className="col-12">
                <Form.Item valuePropName="checked" name="confirm3">
                  <Checkbox className="text-uppercase">
                    I CONSENT TO HAVING MDTK SOFT COLLECT MY PERSONAL DETAILS.
                  </Checkbox>
                </Form.Item>
                <Form.Item name="confirm4">
                  <button type="button" className="btn btn-light px-5">
                    Sign in
                  </button>
                </Form.Item>
              </div>
            </div>
          </Form>
        </div>
      </div>
      <div className="card">
        <div className="card-body">
          <h5 className="mb-4">
            <strong>Validation & Background</strong>
          </h5>
          <div className="bg-light rounded-lg p-5">
            <div className="row">
              <div className="col-lg-8 mx-auto">
                <Form layout="vertical">
                  <Form.Item name="fullname" label="Username">
                    <Input />
                  </Form.Item>
                  <Form.Item name="gender" label="Gender">
                    <Select placeholder="Select a option and change input text above">
                      <Option value="male">male</Option>
                      <Option value="female">female</Option>
                    </Select>
                  </Form.Item>
                  <Form.Item name="email" label="E-mail">
                    <Input />
                  </Form.Item>
                  <Form.Item name="password" label="Password" hasFeedback>
                    <Input.Password />
                  </Form.Item>
                  <Form.Item name="confirm" label="Confirm Password" hasFeedback>
                    <Input.Password onBlur={handleConfirmBlur} />
                  </Form.Item>
                  <div className="border-top pt-4">
                    <Form.Item>
                      <Button type="primary" htmlType="submit">
                        Submit
                      </Button>
                    </Form.Item>
                  </div>
                </Form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default AdvancedFormExamples
