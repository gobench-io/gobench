import React from 'react'
import { Badge, Button } from 'reactstrap'

class BootstrapBadgesExample extends React.Component {
  render() {
    return (
      <div className="row">
        <div className="col-lg-6">
          <h5 className="mb-4">
            <strong>Base</strong>
          </h5>
          <div className="mb-5">
            <h1>
              Example heading <Badge color="light">New</Badge>
            </h1>
            <h2>
              Example heading <Badge color="light">New</Badge>
            </h2>
            <h3>
              Example heading <Badge color="light">New</Badge>
            </h3>
            <h4>
              Example heading <Badge color="light">New</Badge>
            </h4>
            <h5>
              Example heading <Badge color="light">New</Badge>
            </h5>
            <h6>
              Example heading <Badge color="light">New</Badge>
            </h6>
          </div>
        </div>
        <div className="col-lg-6">
          <h5 className="mb-4">
            <strong>In Buttons</strong>
          </h5>
          <div className="mb-5">
            <Button color="primary">
              Notifications <Badge color="light">4</Badge>
            </Button>
          </div>
        </div>
        <div className="col-lg-6">
          <h5 className="mb-4">
            <strong>Colors</strong>
          </h5>
          <div className="mb-5">
            <Badge color="primary" className="mr-2">
              Primary
            </Badge>
            <Badge color="secondary" className="mr-2">
              Secondary
            </Badge>
            <Badge color="success" className="mr-2">
              Success
            </Badge>
            <Badge color="danger" className="mr-2">
              Danger
            </Badge>
            <Badge color="warning" className="mr-2">
              Warning
            </Badge>
            <Badge color="info" className="mr-2">
              Info
            </Badge>
            <Badge color="light" className="mr-2">
              Light
            </Badge>
            <Badge color="dark" className="mr-2">
              Dark
            </Badge>
          </div>
        </div>
        <div className="col-lg-6">
          <h5 className="mb-4">
            <strong>Pills</strong>
          </h5>
          <div className="mb-5">
            <Badge color="primary" pill className="mr-2">
              Primary
            </Badge>
            <Badge color="secondary" pill className="mr-2">
              Secondary
            </Badge>
            <Badge color="success" pill className="mr-2">
              Success
            </Badge>
            <Badge color="danger" pill className="mr-2">
              Danger
            </Badge>
            <Badge color="warning" pill className="mr-2">
              Warning
            </Badge>
            <Badge color="info" pill className="mr-2">
              Info
            </Badge>
            <Badge color="light" pill className="mr-2">
              Light
            </Badge>
            <Badge color="dark" pill className="mr-2">
              Dark
            </Badge>
          </div>
        </div>
        <div className="col-lg-6">
          <h5 className="mb-4">
            <strong>Links</strong>
          </h5>
          <div className="mb-5">
            <Badge href="#" color="primary" className="mr-2">
              Primary
            </Badge>
            <Badge href="#" color="secondary" className="mr-2">
              Secondary
            </Badge>
            <Badge href="#" color="success" className="mr-2">
              Success
            </Badge>
            <Badge href="#" color="danger" className="mr-2">
              Danger
            </Badge>
            <Badge href="#" color="warning" className="mr-2">
              Warning
            </Badge>
            <Badge href="#" color="info" className="mr-2">
              Info
            </Badge>
            <Badge href="#" color="light" className="mr-2">
              Light
            </Badge>
            <Badge href="#" color="dark" className="mr-2">
              Dark
            </Badge>
          </div>
        </div>
      </div>
    )
  }
}

export default BootstrapBadgesExample
