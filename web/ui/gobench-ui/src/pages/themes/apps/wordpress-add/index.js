import React from 'react'
import { Helmet } from 'react-helmet'
import { Editor } from 'react-draft-wysiwyg'
import { InboxOutlined } from '@ant-design/icons'
import { Input, Checkbox, Select, Upload, Form } from 'antd'

const { Dragger } = Upload
const { Option } = Select

const ExtraAppsWordpressAdd = () => {
  return (
    <div>
      <Helmet title="Wordpress Add" />
      <div className="card">
        <div className="card-body">
          <Form layout="vertical">
            <div className="row">
              <div className="col-md-6">
                <div className="form-group">
                  <Form.Item name="title" label="Title">
                    <Input placeholder="Post title" />
                  </Form.Item>
                </div>
              </div>
            </div>
            <div className="form-group">
              <Form.Item name="type" label="Type">
                <Checkbox.Group>
                  <div className="d-flex flex-wrap">
                    <div className="mr-3 mt-1 mb-1">
                      <Checkbox value="text">Text</Checkbox>
                    </div>
                    <div className="mr-3 mt-1 mb-1">
                      <Checkbox value="video">Video</Checkbox>
                    </div>
                    <div className="mr-3 mt-1 mb-1">
                      <Checkbox value="image">Image</Checkbox>
                    </div>
                    <div className="mr-3 mt-1 mb-1">
                      <Checkbox value="audio">Audio</Checkbox>
                    </div>
                    <div className="mr-3 mt-1 mb-1">
                      <Checkbox value="vimeo">Vimeo</Checkbox>
                    </div>
                  </div>
                </Checkbox.Group>
              </Form.Item>
            </div>
            <div className="form-group">
              <Form.Item name="category" label="Category">
                <Select
                  mode="tags"
                  size="default"
                  placeholder="Select post category"
                  style={{ width: '100%' }}
                >
                  <Option key="travel">Travel</Option>
                  <Option key="lifestyle">Lifestyle</Option>
                  <Option key="nature">Nature</Option>
                  <Option key="Video">Video</Option>
                </Select>
              </Form.Item>
            </div>
            <div className="form-group">
              <Form.Item name="content" label="Content">
                <Editor
                  editorClassName="px-3 border border-gray-1"
                  editorStyle={{
                    height: 250,
                    overflow: 'auto',
                  }}
                />
              </Form.Item>
            </div>
            <div className="form-group">
              <Form.Item valuePropName="fileList" name="files">
                <Dragger>
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
            <Form.Item>
              <button type="submit" className="btn btn-success btn-with-addon text-nowrap">
                <span className="btn-addon">
                  <i className="btn-addon-icon fe fe-plus-circle" />
                </span>
                Add Post
              </button>
            </Form.Item>
          </Form>
        </div>
      </div>
    </div>
  )
}

export default ExtraAppsWordpressAdd
