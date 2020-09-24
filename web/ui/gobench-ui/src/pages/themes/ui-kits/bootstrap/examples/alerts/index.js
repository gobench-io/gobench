import React from 'react'
import { Alert } from 'reactstrap'

class BootstrapAlertsExample extends React.Component {
  render() {
    return (
      <div className="row">
        <div className="col-lg-6">
          <h5 className="mb-3">
            <strong>Simple</strong>
          </h5>
          <Alert color="primary">This is a primary alert — check it out!</Alert>
          <Alert color="secondary">This is a secondary alert — check it out!</Alert>
          <Alert color="success">This is a success alert — check it out!</Alert>
          <Alert color="danger">This is a danger alert — check it out!</Alert>
          <Alert color="warning">This is a warning alert — check it out!</Alert>
          <Alert color="info">This is a info alert — check it out!</Alert>
          <Alert color="light">This is a light alert — check it out!</Alert>
          <Alert color="dark">This is a dark alert — check it out!</Alert>
        </div>
        <div className="col-lg-6">
          <h5 className="mb-3">
            <strong>Alert with List</strong>
          </h5>
          <Alert color="light">
            <p>
              <strong>Read documentation and check devices:</strong>
            </p>
            <ul>
              <li>Connections</li>
              <li>Cables &amp; Accessories</li>
              <li>Display &amp; Touch</li>
              <li>Drivers</li>
            </ul>
          </Alert>
        </div>
      </div>
    )
  }
}

export default BootstrapAlertsExample
