import React from 'react'
import { Breadcrumb, BreadcrumbItem } from 'reactstrap'

class BootstrapBreadcrumbsExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Breadcrumbs</strong>
        </h5>
        <Breadcrumb>
          <BreadcrumbItem active>Home</BreadcrumbItem>
        </Breadcrumb>
        <Breadcrumb>
          <BreadcrumbItem>
            <a>Home</a>
          </BreadcrumbItem>
          <BreadcrumbItem active>Library</BreadcrumbItem>
        </Breadcrumb>
        <Breadcrumb>
          <BreadcrumbItem>
            <a>Home</a>
          </BreadcrumbItem>
          <BreadcrumbItem>
            <a>Library</a>
          </BreadcrumbItem>
          <BreadcrumbItem active>Data</BreadcrumbItem>
        </Breadcrumb>
      </div>
    )
  }
}

export default BootstrapBreadcrumbsExample
