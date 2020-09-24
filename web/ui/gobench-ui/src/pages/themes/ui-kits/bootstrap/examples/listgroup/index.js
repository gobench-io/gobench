import React from 'react'
import { ListGroup, ListGroupItem } from 'reactstrap'

class BootstrapListGroupExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Layout</strong>
        </h5>
        <div className="mb-5">
          <ListGroup>
            <ListGroupItem>Cras justo odio</ListGroupItem>
            <ListGroupItem>Dapibus ac facilisis in</ListGroupItem>
            <ListGroupItem>Morbi leo risus</ListGroupItem>
            <ListGroupItem>Porta ac consectetur ac</ListGroupItem>
            <ListGroupItem>Vestibulum at eros</ListGroupItem>
          </ListGroup>
        </div>
        <h5 className="mb-4">
          <strong>Anchors and buttons</strong>
        </h5>
        <div className="mb-5">
          <div className="mb-5">
            <ListGroup>
              <ListGroupItem active tag="a" href="#" action>
                Cras justo odio
              </ListGroupItem>
              <ListGroupItem tag="a" href="#" action>
                Dapibus ac facilisis in
              </ListGroupItem>
              <ListGroupItem tag="a" href="#" action>
                Morbi leo risus
              </ListGroupItem>
              <ListGroupItem tag="a" href="#" action>
                Porta ac consectetur ac
              </ListGroupItem>
              <ListGroupItem disabled tag="a" href="#" action>
                Vestibulum at eros
              </ListGroupItem>
            </ListGroup>
          </div>
          <ListGroup>
            <ListGroupItem active tag="button" action>
              Cras justo odio
            </ListGroupItem>
            <ListGroupItem tag="button" action>
              Dapibus ac facilisis in
            </ListGroupItem>
            <ListGroupItem tag="button" action>
              Morbi leo risus
            </ListGroupItem>
            <ListGroupItem tag="button" action>
              Porta ac consectetur ac
            </ListGroupItem>
            <ListGroupItem disabled tag="button" action>
              Vestibulum at eros
            </ListGroupItem>
          </ListGroup>
        </div>
        <h5 className="mb-4">
          <strong>Horizontal</strong>
        </h5>
        <div className="mb-5">
          <ListGroup horizontal="lg">
            <ListGroupItem tag="a" href="#">
              Cras justo odio
            </ListGroupItem>
            <ListGroupItem tag="a" href="#">
              Dapibus ac facilisis in
            </ListGroupItem>
            <ListGroupItem tag="a" href="#">
              Morbi leo risus
            </ListGroupItem>
            <ListGroupItem tag="a" href="#">
              Porta ac consectetur ac
            </ListGroupItem>
            <ListGroupItem tag="a" href="#">
              Vestibulum at eros
            </ListGroupItem>
          </ListGroup>
        </div>
      </div>
    )
  }
}

export default BootstrapListGroupExample
