import React from 'react'
import { Button, Popover, PopoverHeader, PopoverBody, UncontrolledPopover } from 'reactstrap'

class BootstrapPopoversExample extends React.Component {
  constructor(props) {
    super(props)

    this.toggle = this.toggle.bind(this)
    this.state = {
      popoverOpen: false,
    }
  }

  toggle() {
    const { popoverOpen } = this.state
    this.setState({
      popoverOpen: !popoverOpen,
    })
  }

  render() {
    const { popoverOpen } = this.state

    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Popovers</strong>
        </h5>
        <div className="mb-5">
          <Button id="Popover1" type="button" color="primary">
            Launch Popover
          </Button>
          <Popover placement="bottom" isOpen={popoverOpen} target="Popover1" toggle={this.toggle}>
            <PopoverHeader>Popover Title</PopoverHeader>
            <PopoverBody>
              Sed posuere consectetur est at lobortis. Aenean eu leo quam. Pellentesque ornare sem
              lacinia quam venenatis vestibulum.
            </PopoverBody>
          </Popover>
        </div>

        <h5 className="mb-4">
          <strong>Popovers Trigger</strong>
        </h5>

        <div className="mb-5">
          <Button id="PopoverFocus" type="button" color="light">
            Launch Popover (Focus)
          </Button>{' '}
          <Button id="PopoverClick" type="button" color="light">
            Launch Popover (Click)
          </Button>{' '}
          <Button id="PopoverLegacy" type="button" color="light">
            Launch Popover (Legacy)
          </Button>
          <UncontrolledPopover trigger="focus" placement="bottom" target="PopoverFocus">
            <PopoverHeader>Focus Trigger</PopoverHeader>
            <PopoverBody>
              Focusing on the trigging element makes this popover appear. Blurring (clicking away)
              makes it disappear. You cannot select this text as the popover will disappear when you
              try.
            </PopoverBody>
          </UncontrolledPopover>
          <UncontrolledPopover trigger="click" placement="bottom" target="PopoverClick">
            <PopoverHeader>Click Trigger</PopoverHeader>
            <PopoverBody>
              Clicking on the triggering element makes this popover appear. Clicking on it again
              will make it disappear. You can select this text, but clicking away (somewhere other
              than the triggering element) will not dismiss this popover.
            </PopoverBody>
          </UncontrolledPopover>
          <UncontrolledPopover trigger="legacy" placement="bottom" target="PopoverLegacy">
            <PopoverHeader>Legacy Trigger</PopoverHeader>
            <PopoverBody>
              Legacy is a reactstrap special trigger value (outside of bootstraps spec/standard).
              Before reactstrap correctly supported click and focus, it had a hybrid which was very
              useful and has been brought back as trigger=legacy. One advantage of the legacy
              trigger is that it allows the popover text to be selected while also closing when
              clicking outside the triggering element and popover itself.
            </PopoverBody>
          </UncontrolledPopover>
        </div>
      </div>
    )
  }
}

export default BootstrapPopoversExample
