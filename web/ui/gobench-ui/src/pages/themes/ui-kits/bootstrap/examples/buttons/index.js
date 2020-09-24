import React from 'react'
import { Button } from 'reactstrap'

class BootstrapButtonsExample extends React.Component {
  render() {
    return (
      <div className="row">
        <div className="col-lg-6 mb-5">
          <h5 className="mb-4">
            <strong>Default Buttons</strong>
          </h5>
          <Button className="mr-2 mb-2">Clear</Button>
          <Button color="default" className="mr-2 mb-2">
            Default
          </Button>
          <Button color="default" className="mr-2 mb-2" disabled>
            Default Disabled
          </Button>
          <Button color="link" className="mr-2 mb-2">
            Link
          </Button>
          <br />
          <Button color="primary" className="mr-2 mb-2">
            Primary
          </Button>
          <Button color="secondary" className="mr-2 mb-2">
            Secondary
          </Button>
          <Button color="success" className="mr-2 mb-2">
            Success
          </Button>
          <Button color="danger" className="mr-2 mb-2">
            Danger
          </Button>
          <Button color="warning" className="mr-2 mb-2">
            Warning
          </Button>
          <Button color="info" className="mr-2 mb-2">
            Info
          </Button>
          <Button color="light" className="mr-2 mb-2">
            Light
          </Button>
          <Button color="dark" className="mr-2 mb-2">
            Dark
          </Button>
          <br />
          <Button color="info" outline className="mr-2 mb-2">
            Info
          </Button>
          <Button color="light" outline className="mr-2 mb-2">
            Light
          </Button>
          <Button color="dark" outline className="mr-2 mb-2">
            Dark
          </Button>
          <Button color="primary" outline className="mr-2 mb-2">
            Primary
          </Button>
          <Button color="secondary" outline className="mr-2 mb-2">
            Secondary
          </Button>
          <Button color="success" outline className="mr-2 mb-2">
            Success
          </Button>
          <Button color="danger" outline className="mr-2 mb-2">
            Danger
          </Button>
          <Button color="warning" outline className="mr-2 mb-2">
            Warning
          </Button>
        </div>
        <div className="col-lg-6 mb-5">
          <h5 className="mb-4">
            <strong>Sizing</strong>
          </h5>
          <Button className="mr-2 mb-2">Normal</Button>
          <Button color="default" className="btn-default mr-2 mb-2">
            Normal
          </Button>
          <Button color="default" className="btn-default mr-2 mb-2" disabled="">
            Normal
          </Button>
          <Button href="#" color="link" className="btn-link mr-2 mb-2">
            Normal
          </Button>
          <br />
          <Button color="primary" className="btn-lg mr-2 mb-2">
            Large
          </Button>
          <Button color="danger" className="mr-2 mb-2">
            Normal
          </Button>
          <Button color="warning" className="btn-lg mr-2 mb-2">
            Large
          </Button>
          <Button color="info" className="mr-2 mb-2">
            Normal
          </Button>
          <Button color="success" className="btn-sm mr-2 mb-2">
            Small
          </Button>
          <br />
          <Button outline color="danger" className="mr-2 mb-2">
            Noraml
          </Button>
          <Button outline color="primary" className="btn-sm mr-2 mb-2">
            Small
          </Button>
          <Button outline color="warning" className="btn-lg mr-2 mb-2">
            Large
          </Button>
          <Button outline color="success" className="mr-2 mb-2">
            Normal
          </Button>
        </div>
        <div className="col-lg-6 mb-5">
          <h5 className="mb-4">
            <strong>Squared</strong>
          </h5>
          <Button className="btn-squared mr-2 mb-2">Clear</Button>
          <Button color="default" className="btn-squared btn-default mr-2 mb-2">
            Default
          </Button>
          <Button color="default" className="btn-squared btn-default mr-2 mb-2" disabled>
            Default Disabled
          </Button>
          <Button color="link" className="btn-squared btn-link mr-2 mb-2">
            Link
          </Button>
          <br />
          <Button color="warning" className="btn-squared mr-2 mb-2">
            Warning
          </Button>
          <Button color="info" className="btn-squared mr-2 mb-2">
            Info
          </Button>
          <Button color="light" className="btn-squared mr-2 mb-2">
            Light
          </Button>
          <Button color="success" className="btn-squared mr-2 mb-2">
            Success
          </Button>
          <Button color="danger" className="btn-squared mr-2 mb-2">
            Danger
          </Button>
          <Button color="dark" className="btn-squared mr-2 mb-2">
            Dark
          </Button>
          <Button color="primary" className="btn-squared mr-2 mb-2">
            Primary
          </Button>
          <Button color="secondary" className="btn-squared mr-2 mb-2">
            Secondary
          </Button>
          <br />
          <Button color="dark" outline className="mr-2 mb-2">
            Dark
          </Button>
          <Button color="success" outline className="mr-2 mb-2">
            Success
          </Button>
          <Button color="danger" outline className="mr-2 mb-2">
            Danger
          </Button>
          <Button color="warning" outline className="mr-2 mb-2">
            Warning
          </Button>
          <Button color="primary" outline className="mr-2 mb-2">
            Primary
          </Button>
          <Button color="secondary" outline className="mr-2 mb-2">
            Secondary
          </Button>
          <Button color="info" outline className="mr-2 mb-2">
            Info
          </Button>
          <Button color="light" outline className="mr-2 mb-2">
            Light
          </Button>
        </div>
        <div className="col-lg-6 mb-5">
          <h5 className="mb-4">
            <strong>Rounded</strong>
          </h5>
          <Button className="btn-rounded mr-2 mb-2">Clear</Button>
          <Button color="default" className="btn-rounded btn-default mr-2 mb-2">
            Default
          </Button>
          <Button color="default" className="btn-rounded btn-default mr-2 mb-2" disabled>
            Default Disabled
          </Button>
          <Button color="link" className="btn-rounded btn-link mr-2 mb-2">
            Link
          </Button>
          <br />
          <Button color="primary" className="btn-rounded mr-2 mb-2">
            Primary
          </Button>
          <Button color="secondary" className="btn-rounded mr-2 mb-2">
            Secondary
          </Button>
          <Button color="success" className="btn-rounded mr-2 mb-2">
            Success
          </Button>
          <Button color="danger" className="btn-rounded mr-2 mb-2">
            Danger
          </Button>
          <Button color="dark" className="btn-rounded mr-2 mb-2">
            Dark
          </Button>
          <Button color="warning" className="btn-rounded mr-2 mb-2">
            Warning
          </Button>
          <Button color="info" className="btn-rounded mr-2 mb-2">
            Info
          </Button>
          <Button color="light" className="btn-rounded mr-2 mb-2">
            Light
          </Button>
          <br />
          <Button color="warning" outline className="mr-2 mb-2">
            Warning
          </Button>
          <Button color="primary" outline className="mr-2 mb-2">
            Primary
          </Button>
          <Button color="seconaary" outline className="mr-2 mb-2">
            Secondary
          </Button>
          <Button color="info" outline className="mr-2 mb-2">
            Info
          </Button>
          <Button color="light" outline className="mr-2 mb-2">
            Light
          </Button>
          <Button color="dark" outline className="mr-2 mb-2">
            Dark
          </Button>
          <Button color="success" outline className="mr-2 mb-2">
            Success
          </Button>
          <Button color="danger" outline className="mr-2 mb-2">
            Danger
          </Button>
        </div>
        <div className="col-lg-6 mb-5">
          <h5 className="mb-4">
            <strong>Icon Buttons</strong>
          </h5>
          <div className="btn-group mr-2 mb-2" aria-label="" role="group">
            <Button color="success">
              <i className="fe fe-edit mr-1" aria-hidden="true" />
              Edit
            </Button>
            <Button color="success">
              <i className="fe fe-send mr-1" aria-hidden="true" />
              Reply
            </Button>
            <Button color="success">
              <i className="fe fe-share mr-1" aria-hidden="true" />
              Share
            </Button>
          </div>
          <div className="btn-group mr-2 mb-2" aria-label="" role="group">
            <Button color="light">
              <i className="fe fe-edit mr-1" aria-hidden="true" />
              Edit
            </Button>
            <Button color="light">
              <i className="fe fe-send mr-1" aria-hidden="true" />
              Reply
            </Button>
            <Button color="light">
              <i className="fe fe-share mr-1" aria-hidden="true" />
              Share
            </Button>
          </div>
          <br />
          <Button color="primary" className="btn-rounded mr-2 mb-2">
            <i className="fe fe-activity" />
          </Button>
          <Button color="secondary" className="btn-rounded mr-2 mb-2">
            <i className="fe fe-alert-triangle" />
          </Button>
          <Button color="success" className="btn-rounded mr-2 mb-2">
            <i className="fe fe-anchor" />
          </Button>
          <Button color="danger" className="btn-rounded mr-2 mb-2">
            <i className="fe fe-award" />
          </Button>
          <Button color="warning" className="btn-rounded mr-2 mb-2">
            <i className="fe fe-battery" />
          </Button>
          <Button color="info" className="btn-rounded mr-2 mb-2">
            <i className="fe fe-clock" />
          </Button>
          <Button color="light" className="btn-rounded mr-2 mb-2">
            <i className="fe fe-code" />
          </Button>
          <Button color="dark" className="btn-rounded mr-2 mb-2">
            <i className="fe fe-credit-card" />
          </Button>
        </div>
        <div className="col-lg-4 mb-3">
          <h5 className="mb-4">
            <strong>Block Buttons</strong>
          </h5>
          <Button block color="primary" className="btn-rounded mr-2 mb-2">
            <i className="fe fe-save mr-1" />
            Primary
          </Button>
          <Button block color="success" className="btn-rounded mr-2 mb-2">
            Success
          </Button>
          <Button block className="btn-rounded mr-2 mb-2">
            Clear
          </Button>
        </div>
      </div>
    )
  }
}

export default BootstrapButtonsExample
