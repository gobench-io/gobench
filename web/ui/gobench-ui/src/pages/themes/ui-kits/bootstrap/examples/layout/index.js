import React from 'react'
import { Container, Row, Col } from 'reactstrap'

class BootstrapLayoutExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Layout</strong>
        </h5>
        <div className="mb-5">
          <div className="kit__utils__docs">
            <Container fluid>
              <Row>
                <Col>.col</Col>
              </Row>
              <Row>
                <Col>.col</Col>
                <Col>.col</Col>
                <Col>.col</Col>
                <Col>.col</Col>
              </Row>
              <Row>
                <Col xs="3">.col-3</Col>
                <Col xs="auto">.col-auto - variable width content</Col>
                <Col xs="3">.col-3</Col>
              </Row>
              <Row>
                <Col xs="6">.col-6</Col>
                <Col xs="6">.col-6</Col>
              </Row>
              <Row>
                <Col xs="6" sm="4">
                  .col-6 .col-sm-4
                </Col>
                <Col xs="6" sm="4">
                  .col-6 .col-sm-4
                </Col>
                <Col sm="4">.col-sm-4</Col>
              </Row>
              <Row>
                <Col sm={{ size: 6, order: 2, offset: 1 }}>.col-sm-6 .order-sm-2 .offset-sm-1</Col>
              </Row>
              <Row>
                <Col sm="12" md={{ size: 6, offset: 3 }}>
                  .col-sm-12 .col-md-6 .offset-md-3
                </Col>
              </Row>
              <Row>
                <Col sm={{ size: 'auto', offset: 1 }}>.col-sm-auto .offset-sm-1</Col>
                <Col sm={{ size: 'auto', offset: 1 }}>.col-sm-auto .offset-sm-1</Col>
              </Row>
            </Container>
          </div>
        </div>
        <h5 className="mb-4">
          <strong>Containers</strong>
        </h5>
        <div className="mb-5">
          <div className="kit__utils__docs">
            <Container className="themed-container">.container</Container>
            <Container className="themed-container" fluid="sm">
              .container-sm
            </Container>
            <Container className="themed-container" fluid="md">
              .container-md
            </Container>
            <Container className="themed-container" fluid="lg">
              .container-lg
            </Container>
            <Container className="themed-container" fluid="xl">
              .container-xl
            </Container>
            <Container className="themed-container" fluid>
              .container-fluid
            </Container>
          </div>
        </div>
        <h5 className="mb-4">
          <strong>Row Columns</strong>
        </h5>
        <div className="mb-5">
          <div className="kit__utils__docs">
            <Container fluid>
              <Row xs="2">
                <Col>Column</Col>
                <Col>Column</Col>
                <Col>Column</Col>
                <Col>Column</Col>
              </Row>
              <Row xs="3">
                <Col>Column</Col>
                <Col>Column</Col>
                <Col>Column</Col>
                <Col>Column</Col>
              </Row>
              <Row xs="4">
                <Col>Column</Col>
                <Col>Column</Col>
                <Col>Column</Col>
                <Col>Column</Col>
              </Row>
              <Row xs="4">
                <Col>Column</Col>
                <Col>Column</Col>
                <Col xs="6">Column</Col>
                <Col>Column</Col>
              </Row>
              <Row xs="1" sm="2" md="4">
                <Col>Column</Col>
                <Col>Column</Col>
                <Col>Column</Col>
                <Col>Column</Col>
              </Row>
            </Container>
          </div>
        </div>
      </div>
    )
  }
}

export default BootstrapLayoutExample
