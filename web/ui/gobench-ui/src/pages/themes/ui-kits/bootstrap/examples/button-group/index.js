import React from 'react'
import {
  Button,
  ButtonGroup,
  UncontrolledDropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem,
} from 'reactstrap'

class BootstrapButtonGroupExample extends React.Component {
  render() {
    return (
      <div className="row">
        <div className="col-lg-12 mb-5">
          <h5 className="mb-4">
            <strong>Default Group</strong>
          </h5>
          <ButtonGroup className="mr-2 mb-2">
            <Button color="" className="btn btn-success">
              Left
            </Button>
            <Button color="" className="btn btn-success">
              Middle
            </Button>
            <Button color="" className="btn btn-success">
              Right
            </Button>
          </ButtonGroup>
          <ButtonGroup className="mr-2 mb-2">
            <Button color="white" className="btn btn-rounded">
              Left
            </Button>
            <Button color="white">Middle</Button>
            <Button color="danger" className="btn btn-rounded">
              Right
            </Button>
          </ButtonGroup>
        </div>
        <div className="col-lg-12 mb-5">
          <h5 className="mb-4">
            <strong>Nesting Group</strong>
          </h5>
          <ButtonGroup className="mr-2 mb-2">
            <Button color="primary">1</Button>
            <Button color="primary">2</Button>
            <UncontrolledDropdown>
              <DropdownToggle color="light" caret>
                Dropdown
              </DropdownToggle>
              <DropdownMenu>
                <DropdownItem>Action</DropdownItem>
                <DropdownItem>Another action</DropdownItem>
                <DropdownItem>Something else here</DropdownItem>
                <DropdownItem divider />
                <DropdownItem>Separated link</DropdownItem>
              </DropdownMenu>
            </UncontrolledDropdown>
          </ButtonGroup>
          <ButtonGroup className="mr-2 mb-2">
            <Button color="light">1</Button>
            <Button color="light">2</Button>
            <UncontrolledDropdown>
              <DropdownToggle color="light" caret>
                Dropdown
              </DropdownToggle>
              <DropdownMenu>
                <DropdownItem>Action</DropdownItem>
                <DropdownItem>Another action</DropdownItem>
                <DropdownItem>Something else here</DropdownItem>
                <DropdownItem divider />
                <DropdownItem>Separated link</DropdownItem>
              </DropdownMenu>
            </UncontrolledDropdown>
          </ButtonGroup>
        </div>
        <div className="col-lg-12 mb-3">
          <h5 className="mb-4">
            <strong>Vertical Group</strong>
          </h5>
          <div className="btn-group-vertical mr-2 mb-2">
            <Button className="btn btn-danger">Button</Button>
            <div className="btn-group">
              <Button color="danger">Dropdown</Button>
              <div className="dropdown-menu">
                <a className="dropdown-item" href="">
                  Action
                </a>
                <a className="dropdown-item" href="">
                  Another action
                </a>
                <a className="dropdown-item" href="">
                  Something else here
                </a>
                <div className="dropdown-divider" />
                <a className="dropdown-item" href="">
                  Separated link
                </a>
              </div>
            </div>
            <Button className="btn btn-danger">Button</Button>
            <div className="btn-group">
              <Button color="danger">Dropdown</Button>
              <div className="dropdown-menu">
                <a className="dropdown-item" href="">
                  Action
                </a>
                <a className="dropdown-item" href="">
                  Another action
                </a>
                <a className="dropdown-item" href="">
                  Something else here
                </a>
                <div className="dropdown-divider" />
                <a className="dropdown-item" href="">
                  Separated link
                </a>
              </div>
            </div>
          </div>
          <ButtonGroup vertical>
            <UncontrolledDropdown>
              <DropdownToggle color="light" caret>
                Dropdown
              </DropdownToggle>
              <DropdownMenu>
                <DropdownItem>Action</DropdownItem>
                <DropdownItem>Another action</DropdownItem>
                <DropdownItem>Something else here</DropdownItem>
                <DropdownItem divider />
                <DropdownItem>Separated link</DropdownItem>
              </DropdownMenu>
            </UncontrolledDropdown>
            <Button color="light">Button</Button>
            <Button color="primary">Button</Button>
            <Button color="light">Button</Button>
          </ButtonGroup>
        </div>
      </div>
    )
  }
}

export default BootstrapButtonGroupExample
