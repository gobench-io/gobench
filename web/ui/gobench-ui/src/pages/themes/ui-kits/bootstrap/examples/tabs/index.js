import React from 'react'
import {
  TabContent,
  TabPane,
  Nav,
  NavItem,
  NavLink,
  UncontrolledDropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem,
} from 'reactstrap'
import classnames from 'classnames'

class BootstrapTabsExample extends React.Component {
  constructor(props) {
    super(props)

    this.toggle = this.toggle.bind(this)
    this.state = {
      activeTab: '1',
    }
  }

  toggle(tab) {
    const { activeTab } = this.state
    if (activeTab !== tab) {
      this.setState({
        activeTab: tab,
      })
    }
  }

  render() {
    const { activeTab } = this.state

    return (
      <div>
        <div className="row">
          <div className="col-lg-6">
            <h5 className="mb-4">
              <strong>Custom Tabs</strong>
            </h5>
            <div className="mb-5">
              <Nav tabs className="nav-tabs-line">
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '1' })}
                    onClick={() => {
                      this.toggle('1')
                    }}
                  >
                    Active
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '2' })}
                    onClick={() => {
                      this.toggle('2')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '3' })}
                    onClick={() => {
                      this.toggle('3')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink disabled>Disabled</NavLink>
                </NavItem>
              </Nav>
              <TabContent activeTab={activeTab}>
                <TabPane tabId="1">
                  <div className="p-4 mt-4 border rounded">Tab 1 Contents</div>
                </TabPane>
                <TabPane tabId="2">
                  <div className="p-4 mt-4 border rounded">Tab 2 Contents</div>
                </TabPane>
                <TabPane tabId="3">
                  <div className="p-4 mt-4 border rounded">Tab 3 Contents</div>
                </TabPane>
              </TabContent>
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-4">
              <strong>No Bottom Border</strong>
            </h5>
            <div className="mb-5">
              <Nav tabs className="nav-tabs-line nav-tabs-noborder">
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '1' })}
                    onClick={() => {
                      this.toggle('1')
                    }}
                  >
                    Active
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '2' })}
                    onClick={() => {
                      this.toggle('2')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '3' })}
                    onClick={() => {
                      this.toggle('3')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink disabled>Disabled</NavLink>
                </NavItem>
              </Nav>
              <TabContent activeTab={activeTab}>
                <TabPane tabId="1">
                  <div className="p-4 mt-4 border rounded">Tab 1 Contents</div>
                </TabPane>
                <TabPane tabId="2">
                  <div className="p-4 mt-4 border rounded">Tab 2 Contents</div>
                </TabPane>
                <TabPane tabId="3">
                  <div className="p-4 mt-4 border rounded">Tab 3 Contents</div>
                </TabPane>
              </TabContent>
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-4">
              <strong>Custom Tabs Bold</strong>
            </h5>
            <div className="mb-5">
              <Nav tabs className="nav-tabs-line nav-tabs-line-bold">
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '1' })}
                    onClick={() => {
                      this.toggle('1')
                    }}
                  >
                    Active
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '2' })}
                    onClick={() => {
                      this.toggle('2')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '3' })}
                    onClick={() => {
                      this.toggle('3')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink disabled>Disabled</NavLink>
                </NavItem>
              </Nav>
              <TabContent activeTab={activeTab}>
                <TabPane tabId="1">
                  <div className="p-4 mt-4 border rounded">Tab 1 Contents</div>
                </TabPane>
                <TabPane tabId="2">
                  <div className="p-4 mt-4 border rounded">Tab 2 Contents</div>
                </TabPane>
                <TabPane tabId="3">
                  <div className="p-4 mt-4 border rounded">Tab 3 Contents</div>
                </TabPane>
              </TabContent>
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-4">
              <strong>Custom Tabs Bold with Icons</strong>
            </h5>
            <div className="mb-5">
              <Nav tabs className="nav-tabs-line nav-tabs-line-bold">
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '1' })}
                    onClick={() => {
                      this.toggle('1')
                    }}
                  >
                    <i className="fe fe-activity mr-1" />
                    Active
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '2' })}
                    onClick={() => {
                      this.toggle('2')
                    }}
                  >
                    <i className="fe fe-clock mr-1" />
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '3' })}
                    onClick={() => {
                      this.toggle('3')
                    }}
                  >
                    <i className="fe fe-trash mr-1" />
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink disabled>
                    <i className="fe fe-plus-circle" />
                    Disabled
                  </NavLink>
                </NavItem>
              </Nav>
              <TabContent activeTab={activeTab}>
                <TabPane tabId="1">
                  <div className="p-4 mt-4 border rounded">Tab 1 Contents</div>
                </TabPane>
                <TabPane tabId="2">
                  <div className="p-4 mt-4 border rounded">Tab 2 Contents</div>
                </TabPane>
                <TabPane tabId="3">
                  <div className="p-4 mt-4 border rounded">Tab 3 Contents</div>
                </TabPane>
              </TabContent>
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-4">
              <strong>Custom Tabs Bold with Dropdown</strong>
            </h5>
            <div className="mb-5">
              <Nav tabs className="nav-tabs-line nav-tabs-line-bold">
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '1' })}
                    onClick={() => {
                      this.toggle('1')
                    }}
                  >
                    Active
                  </NavLink>
                </NavItem>
                <UncontrolledDropdown nav>
                  <DropdownToggle nav caret>
                    Dropdown
                  </DropdownToggle>
                  <DropdownMenu>
                    <DropdownItem header>Header</DropdownItem>
                    <DropdownItem disabled>Action</DropdownItem>
                    <DropdownItem>Another Action</DropdownItem>
                    <DropdownItem divider />
                    <DropdownItem>Another Action</DropdownItem>
                  </DropdownMenu>
                </UncontrolledDropdown>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '3' })}
                    onClick={() => {
                      this.toggle('3')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink disabled>Disabled</NavLink>
                </NavItem>
              </Nav>
              <TabContent activeTab={activeTab}>
                <TabPane tabId="1">
                  <div className="p-4 mt-4 border rounded">Tab 1 Contents</div>
                </TabPane>
                <TabPane tabId="2">
                  <div className="p-4 mt-4 border rounded">Tab 2 Contents</div>
                </TabPane>
                <TabPane tabId="3">
                  <div className="p-4 mt-4 border rounded">Tab 3 Contents</div>
                </TabPane>
              </TabContent>
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-4">
              <strong>Default Tabs</strong>
            </h5>
            <div className="mb-5">
              <Nav tabs>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '1' })}
                    onClick={() => {
                      this.toggle('1')
                    }}
                  >
                    Active
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '2' })}
                    onClick={() => {
                      this.toggle('2')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '3' })}
                    onClick={() => {
                      this.toggle('3')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink disabled>Disabled</NavLink>
                </NavItem>
              </Nav>
              <TabContent activeTab={activeTab}>
                <TabPane tabId="1">
                  <div className="p-4 mt-4 border rounded">Tab 1 Contents</div>
                </TabPane>
                <TabPane tabId="2">
                  <div className="p-4 mt-4 border rounded">Tab 2 Contents</div>
                </TabPane>
                <TabPane tabId="3">
                  <div className="p-4 mt-4 border rounded">Tab 3 Contents</div>
                </TabPane>
              </TabContent>
            </div>
          </div>

          <div className="col-lg-6">
            <h5 className="mb-4">
              <strong>Default Pills</strong>
            </h5>
            <div className="mb-5">
              <Nav pills>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '1' })}
                    onClick={() => {
                      this.toggle('1')
                    }}
                  >
                    Active
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '2' })}
                    onClick={() => {
                      this.toggle('2')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '3' })}
                    onClick={() => {
                      this.toggle('3')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink disabled>Disabled</NavLink>
                </NavItem>
              </Nav>
              <TabContent activeTab={activeTab}>
                <TabPane tabId="1">
                  <div className="p-4 mt-4 border rounded">Tab 1 Contents</div>
                </TabPane>
                <TabPane tabId="2">
                  <div className="p-4 mt-4 border rounded">Tab 2 Contents</div>
                </TabPane>
                <TabPane tabId="3">
                  <div className="p-4 mt-4 border rounded">Tab 3 Contents</div>
                </TabPane>
              </TabContent>
            </div>
          </div>
          <div className="col-lg-6">
            <h5 className="mb-4">
              <strong>Justified Pills</strong>
            </h5>
            <div className="mb-5">
              <Nav pills justified>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '1' })}
                    onClick={() => {
                      this.toggle('1')
                    }}
                  >
                    Active
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '2' })}
                    onClick={() => {
                      this.toggle('2')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '3' })}
                    onClick={() => {
                      this.toggle('3')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink disabled>Disabled</NavLink>
                </NavItem>
              </Nav>
              <TabContent activeTab={activeTab}>
                <TabPane tabId="1">
                  <div className="p-4 mt-4 border rounded">Tab 1 Contents</div>
                </TabPane>
                <TabPane tabId="2">
                  <div className="p-4 mt-4 border rounded">Tab 2 Contents</div>
                </TabPane>
                <TabPane tabId="3">
                  <div className="p-4 mt-4 border rounded">Tab 3 Contents</div>
                </TabPane>
              </TabContent>
            </div>
          </div>

          <div className="col-lg-6">
            <h5 className="mb-4">
              <strong>Default Pills with Icons</strong>
            </h5>
            <div className="mb-5">
              <Nav pills>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '1' })}
                    onClick={() => {
                      this.toggle('1')
                    }}
                  >
                    <i className="fe fe-activity mr-1" />
                    Active
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '2' })}
                    onClick={() => {
                      this.toggle('2')
                    }}
                  >
                    <i className="fe fe-clock mr-1" />
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '3' })}
                    onClick={() => {
                      this.toggle('3')
                    }}
                  >
                    <i className="fe fe-trash mr-1" />
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink disabled>
                    <i className="fe fe-plus-circle" />
                    Disabled
                  </NavLink>
                </NavItem>
              </Nav>
              <TabContent activeTab={activeTab}>
                <TabPane tabId="1">
                  <div className="p-4 mt-4 border rounded">Tab 1 Contents</div>
                </TabPane>
                <TabPane tabId="2">
                  <div className="p-4 mt-4 border rounded">Tab 2 Contents</div>
                </TabPane>
                <TabPane tabId="3">
                  <div className="p-4 mt-4 border rounded">Tab 3 Contents</div>
                </TabPane>
              </TabContent>
            </div>
          </div>

          <div className="col-lg-6">
            <h5 className="mb-4">
              <strong>Default Pills with Dropdown</strong>
            </h5>
            <div className="mb-5">
              <Nav pills>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '1' })}
                    onClick={() => {
                      this.toggle('1')
                    }}
                  >
                    Active
                  </NavLink>
                </NavItem>
                <UncontrolledDropdown nav>
                  <DropdownToggle nav caret>
                    Dropdown
                  </DropdownToggle>
                  <DropdownMenu>
                    <DropdownItem header>Header</DropdownItem>
                    <DropdownItem disabled>Action</DropdownItem>
                    <DropdownItem>Another Action</DropdownItem>
                    <DropdownItem divider />
                    <DropdownItem>Another Action</DropdownItem>
                  </DropdownMenu>
                </UncontrolledDropdown>
                <NavItem>
                  <NavLink
                    className={classnames({ active: activeTab === '3' })}
                    onClick={() => {
                      this.toggle('3')
                    }}
                  >
                    Link
                  </NavLink>
                </NavItem>
                <NavItem>
                  <NavLink disabled>Disabled</NavLink>
                </NavItem>
              </Nav>
              <TabContent activeTab={activeTab}>
                <TabPane tabId="1">
                  <div className="p-4 mt-4 border rounded">Tab 1 Contents</div>
                </TabPane>
                <TabPane tabId="2">
                  <div className="p-4 mt-4 border rounded">Tab 2 Contents</div>
                </TabPane>
                <TabPane tabId="3">
                  <div className="p-4 mt-4 border rounded">Tab 3 Contents</div>
                </TabPane>
              </TabContent>
            </div>
          </div>
          <div className="col-lg-12">
            <h5 className="mb-4">
              <strong>Vertical Pills</strong>
            </h5>
            <div className="mb-5">
              <div className="row">
                <div className="col-3">
                  <Nav pills vertical>
                    <NavItem>
                      <NavLink
                        className={classnames({ active: activeTab === '1' })}
                        onClick={() => {
                          this.toggle('1')
                        }}
                      >
                        Home
                      </NavLink>
                    </NavItem>
                    <NavItem>
                      <NavLink
                        className={classnames({ active: activeTab === '2' })}
                        onClick={() => {
                          this.toggle('2')
                        }}
                      >
                        Profile
                      </NavLink>
                    </NavItem>
                    <NavItem>
                      <NavLink
                        className={classnames({ active: activeTab === '3' })}
                        onClick={() => {
                          this.toggle('3')
                        }}
                      >
                        Messages
                      </NavLink>
                    </NavItem>
                    <NavItem>
                      <NavLink disabled>Disabled</NavLink>
                    </NavItem>
                  </Nav>
                </div>
                <div className="col-9">
                  <TabContent activeTab={activeTab}>
                    <TabPane tabId="1">
                      Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                      Lorem Ipsum has been the industrys standard dummy text ever since the 1500s,
                      when an unknown printer took a galley of type and scrambled it to make a type
                      specimen book. It has survived not only five centuries, but also the leap into
                      electronic typesetting, remaining essentially unchanged. It was popularised in
                      the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,
                      and more recently with desktop publishing software like Aldus PageMaker
                      including versions of Lorem Ipsum.
                    </TabPane>
                    <TabPane tabId="2">
                      Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                      When an unknown printer took a galley of type and scrambled it to make a type
                      specimen book. It has survived not only five centuries, but also the leap into
                      electronic typesetting, remaining essentially unchanged. It was popularised in
                      the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,
                      and more recently with desktop publishing software like Aldus PageMaker
                      including versions of Lorem Ipsum.
                    </TabPane>
                    <TabPane tabId="3">
                      Lorem Ipsum has been the industrys standard dummy text ever since the 1500s,
                      when an unknown printer took a galley of type and scrambled it to make a type
                      specimen book. It has survived not only five centuries, but also the leap into
                      electronic typesetting, remaining essentially unchanged. It was popularised in
                      the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,
                      and more recently with desktop publishing software like Aldus PageMaker
                      including versions of Lorem Ipsum.
                    </TabPane>
                  </TabContent>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default BootstrapTabsExample
