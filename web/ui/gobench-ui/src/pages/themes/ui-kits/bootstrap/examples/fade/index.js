import React from 'react'
import { Button, Fade } from 'reactstrap'

class BootstrapFadeExample extends React.Component {
  constructor(props) {
    super(props)
    this.state = { fadeIn: true }
  }

  toggle = () => {
    const { fadeIn } = this.state

    this.setState({
      fadeIn: !fadeIn,
    })
  }

  render() {
    const { fadeIn } = this.state

    return (
      <div>
        <h5 className="mb-4">
          <strong>Fade Effect</strong>
        </h5>
        <div>
          <Button color="primary" onClick={this.toggle}>
            Toggle Fade
          </Button>
          <Fade in={fadeIn} tag="h5" className="mt-3">
            This content will fade in and out as the button is pressed
          </Fade>
        </div>
      </div>
    )
  }
}

export default BootstrapFadeExample
