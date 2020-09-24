import React from 'react'
import { Toast, ToastBody, ToastHeader } from 'reactstrap'

class BootstrapToastsExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Toasts</strong>
        </h5>
        <div className="row mb-3">
          <div className="col-lg-6">
            <div className="p-3 my-2 rounded">
              <Toast>
                <ToastHeader>Reactstrap</ToastHeader>
                <ToastBody>This is a toast on a white background — check it out!</ToastBody>
              </Toast>
            </div>
          </div>
          <div className="col-lg-6">
            <div className="p-3 my-2 rounded bg-docs-transparent-grid">
              <Toast>
                <ToastHeader>Reactstrap</ToastHeader>
                <ToastBody>This is a toast on a gridded background — check it out!</ToastBody>
              </Toast>
            </div>
          </div>
        </div>
        <h5 className="mb-4">
          <strong>Colored Toasts</strong>
        </h5>
        <div className="row">
          <div className="col-lg-6">
            <div className="p-3 bg-primary my-2 rounded">
              <Toast>
                <ToastHeader>Reactstrap</ToastHeader>
                <ToastBody>This is a toast on a primary background — check it out!</ToastBody>
              </Toast>
            </div>
          </div>
          <div className="col-lg-6">
            <div className="p-3 bg-secondary my-2 rounded">
              <Toast>
                <ToastHeader>Reactstrap</ToastHeader>
                <ToastBody>This is a toast on a secondary background — check it out!</ToastBody>
              </Toast>
            </div>
          </div>
          <div className="col-lg-6">
            <div className="p-3 bg-success my-2 rounded">
              <Toast>
                <ToastHeader>Reactstrap</ToastHeader>
                <ToastBody>This is a toast on a success background — check it out!</ToastBody>
              </Toast>
            </div>
          </div>
          <div className="col-lg-6">
            <div className="p-3 bg-danger my-2 rounded">
              <Toast>
                <ToastHeader>Reactstrap</ToastHeader>
                <ToastBody>This is a toast on a danger background — check it out!</ToastBody>
              </Toast>
            </div>
          </div>
          <div className="col-lg-6">
            <div className="p-3 bg-warning my-2 rounded">
              <Toast>
                <ToastHeader>Reactstrap</ToastHeader>
                <ToastBody>This is a toast on a warning background — check it out!</ToastBody>
              </Toast>
            </div>
          </div>
          <div className="col-lg-6">
            <div className="p-3 bg-info my-2 rounded">
              <Toast>
                <ToastHeader>Reactstrap</ToastHeader>
                <ToastBody>This is a toast on an info background — check it out!</ToastBody>
              </Toast>
            </div>
          </div>
          <div className="col-lg-6">
            <div className="p-3 bg-dark my-2 rounded">
              <Toast>
                <ToastHeader>Reactstrap</ToastHeader>
                <ToastBody>This is a toast on a dark background — check it out!</ToastBody>
              </Toast>
            </div>
          </div>
          <div className="col-lg-6">
            <div className="p-3 my-2 rounded" style={{ background: 'black' }}>
              <Toast>
                <ToastHeader>Reactstrap</ToastHeader>
                <ToastBody>This is a toast on a black background — check it out!</ToastBody>
              </Toast>
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default BootstrapToastsExample
