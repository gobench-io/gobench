import React from 'react'
import {
  Card,
  CardImg,
  CardText,
  CardBody,
  CardTitle,
  CardSubtitle,
  Button,
  CardLink,
} from 'reactstrap'

class BootstrapCardExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Cards</strong>
        </h5>
        <div>
          <Card>
            <CardImg
              top
              width="100%"
              src="https://via.placeholder.com/1300x300/f0f2f4/e4e9f0"
              alt="Card cap"
            />
            <CardBody>
              <CardTitle>Card title</CardTitle>
              <CardSubtitle>Card subtitle</CardSubtitle>
              <CardText>
                Some quick example text to build on the card title and make up the bulk of the cards
                content.
              </CardText>
              <Button color="primary">Button</Button>
            </CardBody>
          </Card>
        </div>
        <div className="mb-5">
          <Card>
            <CardBody>
              <CardTitle>Card title</CardTitle>
              <CardSubtitle>Card subtitle</CardSubtitle>
            </CardBody>
            <CardImg
              top
              width="100%"
              src="https://via.placeholder.com/1300x300/f0f2f4/e4e9f0"
              alt="Card cap"
            />
            <CardBody>
              <CardText>
                Some quick example text to build on the card title and make up the bulk of the cards
                content.
              </CardText>
              <CardLink>Card Link</CardLink>
              <CardLink>Another Link</CardLink>
            </CardBody>
          </Card>
        </div>
        <h5 className="mb-4">
          <strong>Colorful Cards</strong>
        </h5>
        <div className="row">
          <div className="col-lg-6">
            <Card body inverse style={{ backgroundColor: '#333', borderColor: '#333' }}>
              <CardTitle>Special Title Treatment</CardTitle>
              <CardText>
                With supporting text below as a natural lead-in to additional content.
              </CardText>
              <Button>Button</Button>
            </Card>
          </div>
          <div className="col-lg-6">
            <Card body inverse color="primary">
              <CardTitle>Special Title Treatment</CardTitle>
              <CardText>
                With supporting text below as a natural lead-in to additional content.
              </CardText>
              <Button color="light">Button</Button>
            </Card>
          </div>
          <div className="col-lg-6">
            <Card body inverse color="success">
              <CardTitle>Special Title Treatment</CardTitle>
              <CardText>
                With supporting text below as a natural lead-in to additional content.
              </CardText>
              <Button color="light">Button</Button>
            </Card>
          </div>
          <div className="col-lg-6">
            <Card body inverse color="info">
              <CardTitle>Special Title Treatment</CardTitle>
              <CardText>
                With supporting text below as a natural lead-in to additional content.
              </CardText>
              <Button color="light">Button</Button>
            </Card>
          </div>
          <div className="col-lg-6">
            <Card body inverse color="warning">
              <CardTitle>Special Title Treatment</CardTitle>
              <CardText>
                With supporting text below as a natural lead-in to additional content.
              </CardText>
              <Button color="light">Button</Button>
            </Card>
          </div>
          <div className="col-lg-6">
            <Card body inverse color="danger">
              <CardTitle>Special Title Treatment</CardTitle>
              <CardText>
                With supporting text below as a natural lead-in to additional content.
              </CardText>
              <Button color="light">Button</Button>
            </Card>
          </div>
        </div>
      </div>
    )
  }
}

export default BootstrapCardExample
