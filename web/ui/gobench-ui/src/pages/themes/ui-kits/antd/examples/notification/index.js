/* eslint-disable */
import React from 'react'
import { Button, notification } from 'antd'

const openNotificationWithIcon = type => {
  notification[type]({
    message: 'Notification Title',
    description:
      'This is the content of the notification. This is the content of the notification. This is the content of the notification.',
  })
}

class AntdNotificationExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Button onClick={() => openNotificationWithIcon('success')} className="mb-3 mr-3">
            success
          </Button>
          <Button onClick={() => openNotificationWithIcon('info')} className="mb-3 mr-3">
            info
          </Button>
          <Button onClick={() => openNotificationWithIcon('warning')} className="mb-3 mr-3">
            warning
          </Button>
          <Button onClick={() => openNotificationWithIcon('error')} className="mb-3 mr-3">
            error
          </Button>
        </div>
      </div>
    )
  }
}

export default AntdNotificationExample
