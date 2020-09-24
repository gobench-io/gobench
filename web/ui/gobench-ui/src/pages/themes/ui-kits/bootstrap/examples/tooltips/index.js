import React from 'react'
import { Tooltip, UncontrolledTooltip, Button } from 'reactstrap'

class BootstrapTooltipsExample extends React.Component {
  constructor(props) {
    super(props)

    this.toggle = this.toggle.bind(this)
    this.state = {
      tooltipOpen: false,
    }
  }

  toggle() {
    const { tooltipOpen } = this.state

    this.setState({
      tooltipOpen: !tooltipOpen,
    })
  }

  render() {
    const { tooltipOpen } = this.state

    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Tooltip</strong>
        </h5>
        <div className="mb-5">
          <p>
            Somewhere in here is a{' '}
            <span className="kit__utils__link" id="TooltipExample">
              tooltip
            </span>
            .
          </p>
          <Tooltip
            placement="right"
            isOpen={tooltipOpen}
            target="TooltipExample"
            toggle={this.toggle}
          >
            Hello world!
          </Tooltip>
        </div>
        <h5 className="mb-4">
          <strong>Directions</strong>
        </h5>
        <div className="mb-5">
          <Button color="light" id="u1" className="mr-3">
            Tooltip on top
          </Button>
          <Button color="light" id="u2" className="mr-3">
            Tooltip on top
          </Button>
          <Button color="light" id="u3" className="mr-3">
            Tooltip on top
          </Button>
          <Button color="light" id="u4" className="mr-3">
            Tooltip on top
          </Button>
          <UncontrolledTooltip placement="top" target="u1">
            Hello world!
          </UncontrolledTooltip>
          <UncontrolledTooltip placement="right" target="u2">
            Hello world!
          </UncontrolledTooltip>
          <UncontrolledTooltip placement="right" target="u3">
            Hello world!
          </UncontrolledTooltip>
          <UncontrolledTooltip placement="right" target="u4">
            Hello world!
          </UncontrolledTooltip>
        </div>
      </div>
    )
  }
}

export default BootstrapTooltipsExample
