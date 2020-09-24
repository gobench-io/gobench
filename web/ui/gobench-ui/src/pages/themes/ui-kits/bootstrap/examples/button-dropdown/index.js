import React from 'react'
import {
  UncontrolledButtonDropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem,
  Button,
} from 'reactstrap'

class BootstrapButtonDropdownExample extends React.Component {
  render() {
    return (
      <div className="row">
        <div className="col-lg-12 mb-5">
          <h5 className="mb-4">
            <strong>Default Dropdowns</strong>
          </h5>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle caret>Clear</DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="default" caret>
              Default
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="primary" caret>
              Primary
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="secondary" caret>
              Secondary
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="success" caret>
              Success
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="info" caret>
              Info
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="warning" caret>
              Warning
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="danger" caret>
              Danger
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="light" caret>
              Light
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="dark" caret>
              Dark
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <br />
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-default" caret>
              Default
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-primary" caret>
              Primary
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-secondary" caret>
              Secondary
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-success" caret>
              Success
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-info" caret>
              Info
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-warning" caret>
              Warning
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-danger" caret>
              Danger
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-light" caret>
              Light
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-dark" caret>
              Dark
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </div>
        <div className="col-lg-4 mb-5">
          <h5 className="mb-4">
            <strong>Alignment</strong>
          </h5>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-default" caret>
              Right Aligned Dropdown
            </DropdownToggle>
            <DropdownMenu right>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </div>
        <div className="col-lg-4 mb-5">
          <h5 className="mb-4">
            <strong>Dividers</strong>
          </h5>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-default" caret>
              With dividers
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem divider />
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </div>
        <div className="col-lg-4 mb-5">
          <h5 className="mb-4">
            <strong>Headers</strong>
          </h5>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-default" caret>
              With headers
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem header>Header</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </div>
        <div className="col-lg-4 mb-5">
          <h5 className="mb-4">
            <strong>Icons</strong>
          </h5>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-default" caret>
              With icons
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem>
                <i className="fe fe-activity mr-2" />
                Reply
              </DropdownItem>
              <DropdownItem>
                <i className="fe fe-clock mr-2" />
                Share
              </DropdownItem>
              <DropdownItem>
                <i className="fe fe-credit-card mr-2" />
                Delete
              </DropdownItem>
              <DropdownItem>
                <i className="fe fe-plus-circle mr-2" />
                Add
              </DropdownItem>
              <DropdownItem divider />
              <DropdownItem>
                <i className="fe fe-settings mr-2" />
                Settings
              </DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </div>
        <div className="col-lg-4 mb-5">
          <h5 className="mb-4">
            <strong>Active Item</strong>
          </h5>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-default" caret>
              With active item
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem divider />
              <DropdownItem active>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </div>
        <div className="col-lg-4 mb-5">
          <h5 className="mb-4">
            <strong>Disabled Item</strong>
          </h5>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-default" caret>
              Disabled items
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem divider />
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem disabled>Another Action</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Enabled</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </div>
        <div className="col-lg-4 mb-5">
          <h5 className="mb-4">
            <strong>Sizing</strong>
          </h5>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-default" size="lg" caret>
              Large
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem divider />
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
          <UncontrolledButtonDropdown className="mb-2 mr-2">
            <DropdownToggle color="outline-default" size="sm" caret>
              Small
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem divider />
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </div>
        <div className="col-lg-4 mb-5">
          <h5 className="mb-4">
            <strong>Split Buttons</strong>
          </h5>
          <UncontrolledButtonDropdown>
            <Button color="outline-default">Split buttons</Button>
            <DropdownToggle caret color="outline-default" />
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </div>
        <div className="col-lg-4 mb-5">
          <h5 className="mb-4">
            <strong>Dropup</strong>
          </h5>
          <UncontrolledButtonDropdown className="mb-2 mr-2" direction="up">
            <DropdownToggle color="outline-default" caret>
              Default
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem header>Header</DropdownItem>
              <DropdownItem divider />
              <DropdownItem disabled>Action</DropdownItem>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Another Action</DropdownItem>
            </DropdownMenu>
          </UncontrolledButtonDropdown>
        </div>
      </div>
    )
  }
}

export default BootstrapButtonDropdownExample
