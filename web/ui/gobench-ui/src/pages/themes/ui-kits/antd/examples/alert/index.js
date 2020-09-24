/* eslint-disable */
import React from 'react'
import { Alert } from 'antd'

class AntdAlertExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <div className="mb-3">
            <Alert closable message="Success Tips" type="success" showIcon />
          </div>
          <div className="mb-3">
            <Alert closable message="Informational Notes" type="info" showIcon />
          </div>
          <div className="mb-3">
            <Alert closable message="Warning" type="warning" showIcon />
          </div>
          <div className="mb-3">
            <Alert closable message="Error" type="error" showIcon />
          </div>
        </div>
        <h5 className="mb-3">
          <strong>Advanced</strong>
        </h5>
        <div className="mb-3">
          <Alert
            closable
            message="Success Tips"
            description="Detailed description and advice about successful copywriting."
            type="success"
            showIcon
          />
        </div>
        <div className="mb-5">
          <div className="mb-3">
            <Alert
              closable
              message="Informational Notes"
              description="Additional description and information about copywriting."
              type="info"
              showIcon
            />
          </div>
          <div className="mb-3">
            <Alert
              closable
              message="Warning"
              description="This is a warning notice about copywriting."
              type="warning"
              showIcon
            />
          </div>
          <div className="mb-3">
            <Alert
              closable
              message="Error"
              description="This is an error message about copywriting."
              type="error"
              showIcon
            />
          </div>
        </div>
      </div>
    )
  }
}

export default AntdAlertExample
