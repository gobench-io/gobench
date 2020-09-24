/* eslint-disable */
import React from 'react'
import { Modal, Button } from 'antd'

const { confirm } = Modal

function info() {
  Modal.info({
    title: 'This is a notification message',
    content: (
      <div>
        <p>some messages...some messages...</p>
        <p>some messages...some messages...</p>
      </div>
    ),
    onOk() {},
  })
}

function success() {
  Modal.success({
    title: 'This is a success message',
    content: 'some messages...some messages...',
  })
}

function error() {
  Modal.error({
    title: 'This is an error message',
    content: 'some messages...some messages...',
  })
}

function warning() {
  Modal.warning({
    title: 'This is a warning message',
    content: 'some messages...some messages...',
  })
}

class AntdModalExample extends React.Component {
  state = { visible: false }

  showModal = () => {
    this.setState({
      visible: true,
    })
  }

  handleOk = () => {
    this.setState({
      visible: false,
    })
  }

  handleCancel = () => {
    this.setState({
      visible: false,
    })
  }

  showConfirm = () => {
    confirm({
      title: 'Do you Want to delete these items?',
      content: 'Some descriptions',
      onOk() {
        console.log('OK')
      },
      onCancel() {
        console.log('Cancel')
      },
    })
  }

  showDeleteConfirm = () => {
    confirm({
      title: 'Are you sure delete this task?',
      content: 'Some descriptions',
      okText: 'Yes',
      okType: 'danger',
      cancelText: 'No',
      onOk() {
        console.log('OK')
      },
      onCancel() {
        console.log('Cancel')
      },
    })
  }

  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic & Confirm</strong>
        </h5>
        <div className="mb-5">
          <Button type="primary" onClick={this.showModal} className="mb-3 mr-3">
            Open Modal
          </Button>
          <Button onClick={this.showConfirm} className="mb-3 mr-3">
            Confirm
          </Button>
          <Button onClick={this.showDeleteConfirm} type="dashed" className="mb-3 mr-3">
            Delete
          </Button>
        </div>
        <h5 className="mb-3">
          <strong>Notification Modals</strong>
        </h5>
        <div className="mb-5">
          <Button onClick={info} className="mb-3 mr-3">
            Info
          </Button>
          <Button onClick={success} className="mb-3 mr-3">
            Success
          </Button>
          <Button onClick={error} className="mb-3 mr-3">
            Error
          </Button>
          <Button onClick={warning} className="mb-3 mr-3">
            Warning
          </Button>
        </div>
        <Modal
          title="Basic Modal"
          visible={this.state.visible}
          onOk={this.handleOk}
          onCancel={this.handleCancel}
        >
          <p>Some contents...</p>
          <p>Some contents...</p>
          <p>Some contents...</p>
        </Modal>
      </div>
    )
  }
}

export default AntdModalExample
