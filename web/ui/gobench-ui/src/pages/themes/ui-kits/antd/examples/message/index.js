/* eslint-disable */
import React from 'react'
import { message, Button } from 'antd'

const success = () => {
  message.success('This is a success message')
}

const error = () => {
  message.error('This is an error message')
}

const info = () => {
  message.info('This is an info message')
}

const warning = () => {
  message.warning('This is a warning message')
}

class AntdMessageExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Button onClick={success} className="mr-3 mb-3">
            success
          </Button>
          <Button onClick={info} className="mr-3 mb-3">
            info
          </Button>
          <Button onClick={warning} className="mr-3 mb-3">
            warning
          </Button>
          <Button onClick={error} className="mr-3 mb-3">
            error
          </Button>
        </div>
      </div>
    )
  }
}

export default AntdMessageExample
