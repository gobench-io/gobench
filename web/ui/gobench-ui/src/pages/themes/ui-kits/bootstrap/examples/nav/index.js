import React from 'react'
import { Nav, NavItem, NavLink } from 'reactstrap'

class BootstrapNavExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Based Navs</strong>
        </h5>
        <div className="mb-5">
          <Nav>
            <NavItem>
              <NavLink>Link</NavLink>
            </NavItem>
            <NavItem>
              <NavLink>Link</NavLink>
            </NavItem>
            <NavItem>
              <NavLink>Another Link</NavLink>
            </NavItem>
            <NavItem>
              <NavLink disabled>Disabled Link</NavLink>
            </NavItem>
          </Nav>
        </div>
        <h5 className="mb-4">
          <strong>Centered Navs</strong>
        </h5>
        <div className="mb-5">
          <Nav className="justify-content-center">
            <NavLink>Link</NavLink> <NavLink>Link</NavLink> <NavLink>Another Link</NavLink>{' '}
            <NavLink disabled>Disabled Link</NavLink>
          </Nav>
        </div>
        <h5 className="mb-4">
          <strong>Right Aligned Navs</strong>
        </h5>
        <div className="mb-5">
          <Nav className="justify-content-end">
            <NavLink>Link</NavLink> <NavLink>Link</NavLink> <NavLink>Another Link</NavLink>{' '}
            <NavLink disabled>Disabled Link</NavLink>
          </Nav>
        </div>
        <h5 className="mb-4">
          <strong>Vertical Navs</strong>
        </h5>
        <div className="mb-5">
          <Nav vertical>
            <NavLink>Link</NavLink> <NavLink>Link</NavLink> <NavLink>Another Link</NavLink>{' '}
            <NavLink disabled>Disabled Link</NavLink>
          </Nav>
        </div>
      </div>
    )
  }
}

export default BootstrapNavExample
