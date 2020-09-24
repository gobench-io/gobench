import React from 'react'
import { Spinner } from 'reactstrap'

class BootstrapSpinnersExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Progress</strong>
        </h5>
        <div className="mb-5">
          <Spinner color="primary" />
          <Spinner color="secondary" />
          <Spinner color="success" />
          <Spinner color="danger" />
          <Spinner color="warning" />
          <Spinner color="info" />
          <Spinner color="light" />
          <Spinner color="dark" />
        </div>
        <h5 className="mb-4">
          <strong>Growing Spinner</strong>
        </h5>
        <div className="mb-5">
          <Spinner type="grow" color="primary" />
          <Spinner type="grow" color="secondary" />
          <Spinner type="grow" color="success" />
          <Spinner type="grow" color="danger" />
          <Spinner type="grow" color="warning" />
          <Spinner type="grow" color="info" />
          <Spinner type="grow" color="light" />
          <Spinner type="grow" color="dark" />
        </div>
      </div>
    )
  }
}

export default BootstrapSpinnersExample
