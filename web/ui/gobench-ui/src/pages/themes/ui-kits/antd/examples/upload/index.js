/* eslint-disable */
import React from 'react'
import { InboxOutlined, UploadOutlined } from '@ant-design/icons'
import { Upload, message, Button } from 'antd'

const { Dragger } = Upload

const props = {
  name: 'file',
  action: 'https://www.mocky.io/v2/5cc8019d300000980a055e76',
  headers: {
    authorization: 'authorization-text',
  },
  onChange(info) {
    if (info.file.status !== 'uploading') {
      console.log(info.file, info.fileList)
    }
    if (info.file.status === 'done') {
      message.success(`${info.file.name} file uploaded successfully`)
    } else if (info.file.status === 'error') {
      message.error(`${info.file.name} file upload failed.`)
    }
  },
}

class AntdUploadExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Upload {...props}>
            <Button>
              <UploadOutlined /> Click to Upload
            </Button>
          </Upload>
        </div>
        <h5 className="mb-3">
          <strong>Drag & Drop</strong>
        </h5>
        <div className="mb-5">
          <Dragger {...props}>
            <p className="ant-upload-drag-icon">
              <InboxOutlined />
            </p>
            <p className="ant-upload-text">Click or drag file to this area to upload</p>
            <p className="ant-upload-hint">
              Support for a single or bulk upload. Strictly prohibit from uploading company data or
              other band files
            </p>
          </Dragger>
        </div>
      </div>
    )
  }
}

export default AntdUploadExample
