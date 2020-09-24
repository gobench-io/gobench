import React from 'react'
import {
  Collapse,
  Navbar,
  NavbarToggler,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink,
  UncontrolledDropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem,
} from 'reactstrap'

class BootstrapNavbarExample extends React.Component {
  constructor(props) {
    super(props)

    this.toggle = this.toggle.bind(this)
    this.toggleToggler = this.toggleToggler.bind(this)
    this.state = {
      isOpen: false,
      isOpenToggler: false,
    }
  }

  toggle() {
    const { isOpen } = this.state

    this.setState({
      isOpen: !isOpen,
    })
  }

  toggleToggler() {
    const { isOpenToggler } = this.state

    this.setState({
      isOpenToggler: !isOpenToggler,
    })
  }

  render() {
    const { isOpen, isOpenToggler } = this.state
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Navbar</strong>
        </h5>
        <div className="mb-5">
          <Navbar color="light" expand="md">
            <NavbarBrand href="/">reactstrap</NavbarBrand>
            <NavbarToggler onClick={this.toggle} />
            <Collapse isOpen={isOpen} navbar>
              <Nav className="ml-auto" navbar>
                <NavItem>
                  <NavLink href="/components/">Components</NavLink>
                </NavItem>
                <NavItem>
                  <NavLink href="https://github.com/reactstrap/reactstrap">GitHub</NavLink>
                </NavItem>
                <UncontrolledDropdown nav inNavbar>
                  <DropdownToggle nav caret>
                    Options
                  </DropdownToggle>
                  <DropdownMenu right>
                    <DropdownItem>Option 1</DropdownItem>
                    <DropdownItem>Option 2</DropdownItem>
                    <DropdownItem divider />
                    <DropdownItem>Reset</DropdownItem>
                  </DropdownMenu>
                </UncontrolledDropdown>
              </Nav>
            </Collapse>
          </Navbar>
        </div>
        <h5 className="mb-4">
          <strong>Navbar Toggler</strong>
        </h5>
        <div className="mb-5">
          <Navbar color="faded">
            <NavbarBrand href="/" className="mr-auto">
              reactstrap
            </NavbarBrand>
            <NavbarToggler onClick={this.toggleToggler} className="mr-2" />
            <Collapse isOpen={!isOpenToggler} navbar>
              <Nav navbar>
                <NavItem>
                  <NavLink href="/components/">Components</NavLink>
                </NavItem>
                <NavItem>
                  <NavLink href="https://github.com/reactstrap/reactstrap">GitHub</NavLink>
                </NavItem>
              </Nav>
            </Collapse>
          </Navbar>
        </div>
      </div>
    )
  }
}

export default BootstrapNavbarExample
