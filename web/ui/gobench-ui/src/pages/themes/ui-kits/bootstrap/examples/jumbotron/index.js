import React from 'react'
import { Jumbotron, Button, Container } from 'reactstrap'

class BootstrapJumbotronExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Jumbotron</strong>
        </h5>
        <div className="mb-5">
          <Jumbotron>
            <h1 className="font-size-70">
              <strong>Hello, world!</strong>
            </h1>
            <p className="lead">
              This is a simple hero unit, a simple Jumbotron-style component for calling extra
              attention to featured content or information.
            </p>
            <hr className="my-2" />
            <p>
              It uses utility classes for typography and spacing to space content out within the
              larger container.
            </p>
            <p className="lead">
              <Button color="primary">Learn More</Button>
            </p>
          </Jumbotron>
        </div>
        <h5 className="mb-4">
          <strong>Fluid Jumbotron</strong>
        </h5>
        <Jumbotron fluid>
          <Container fluid>
            <h1 className="font-size-70">
              <strong>Fluid jumbotron</strong>
            </h1>
            <p className="lead">
              This is a modified jumbotron that occupies the entire horizontal space of its parent.
            </p>
          </Container>
        </Jumbotron>
      </div>
    )
  }
}

export default BootstrapJumbotronExample
