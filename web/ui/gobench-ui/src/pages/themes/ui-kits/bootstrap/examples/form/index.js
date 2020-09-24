import React from 'react'
import { Button, Form, FormGroup, Label, Input, FormText, Col } from 'reactstrap'

class BootstrapFormExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Form</strong>
        </h5>
        <div className="mb-5">
          <Form>
            <FormGroup>
              <Label for="exampleEmail">Email</Label>
              <Input type="email" name="email" id="exampleEmail" placeholder="with a placeholder" />
            </FormGroup>
            <FormGroup>
              <Label for="examplePassword">Password</Label>
              <Input
                type="password"
                name="password"
                id="examplePassword"
                placeholder="password placeholder"
              />
            </FormGroup>
            <FormGroup>
              <Label for="exampleSelect">Select</Label>
              <Input type="select" name="select" id="exampleSelect">
                <option>1</option>
                <option>2</option>
                <option>3</option>
                <option>4</option>
                <option>5</option>
              </Input>
            </FormGroup>
            <FormGroup>
              <Label for="exampleSelectMulti">Select Multiple</Label>
              <Input type="select" name="selectMulti" id="exampleSelectMulti" multiple>
                <option>1</option>
                <option>2</option>
                <option>3</option>
                <option>4</option>
                <option>5</option>
              </Input>
            </FormGroup>
            <FormGroup>
              <Label for="exampleText">Text Area</Label>
              <Input type="textarea" name="text" id="exampleText" />
            </FormGroup>
            <FormGroup>
              <Label for="exampleFile">File</Label>
              <Input type="file" name="file" id="exampleFile" />
              <FormText color="muted">
                This is some placeholder block-level help text for the above input. Its a bit
                lighter and easily wraps to a new line.
              </FormText>
            </FormGroup>
            <FormGroup tag="fieldset">
              <FormGroup check>
                <Label check className="kit__utils__control kit__utils__control__radio">
                  <Input type="radio" name="radio2" checked />
                  <span className="kit__utils__control__indicator" />
                  Option one is this and that—be sure to include why its great
                </Label>
              </FormGroup>
              <FormGroup check>
                <Label check className="kit__utils__control kit__utils__control__radio">
                  <Input type="radio" name="radio2" />
                  <span className="kit__utils__control__indicator" />
                  Option two can be something else and selecting it will deselect option one
                </Label>
              </FormGroup>
              <FormGroup check disabled>
                <Label check className="kit__utils__control kit__utils__control__radio">
                  <Input type="radio" name="radio2" disabled />
                  <span className="kit__utils__control__indicator" />
                  Option three is disabled
                </Label>
              </FormGroup>
            </FormGroup>
            <FormGroup check>
              <Label check className="kit__utils__control kit__utils__control__checkbox">
                <Input type="checkbox" id="checkbox2" />
                <span className="kit__utils__control__indicator" />
                Check me out
              </Label>
            </FormGroup>
            <div className="border-top mt-4 pt-4">
              <Button color="primary" className="px-5">
                Submit
              </Button>
            </div>
          </Form>
        </div>
        <h5 className="mb-4">
          <strong>Horizontal Form</strong>
        </h5>
        <div className="mb-5">
          <Form>
            <FormGroup row>
              <Label for="exampleEmail" sm={2}>
                Email
              </Label>
              <Col sm={10}>
                <Input
                  type="email"
                  name="email"
                  id="exampleEmail"
                  placeholder="with a placeholder"
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="examplePassword" sm={2}>
                Password
              </Label>
              <Col sm={10}>
                <Input
                  type="password"
                  name="password"
                  id="examplePassword"
                  placeholder="password placeholder"
                />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="exampleSelect" sm={2}>
                Select
              </Label>
              <Col sm={10}>
                <Input type="select" name="select" id="exampleSelect" />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="exampleSelectMulti" sm={2}>
                Select Multiple
              </Label>
              <Col sm={10}>
                <Input type="select" name="selectMulti" id="exampleSelectMulti" multiple />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="exampleText" sm={2}>
                Text Area
              </Label>
              <Col sm={10}>
                <Input type="textarea" name="text" id="exampleText" />
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="exampleFile" sm={2}>
                File
              </Label>
              <Col sm={10}>
                <Input type="file" name="file" id="exampleFile" />
                <FormText color="muted">
                  This is some placeholder block-level help text for the above input. Its a bit
                  lighter and easily wraps to a new line.
                </FormText>
              </Col>
            </FormGroup>
            <FormGroup row>
              <div className="col-sm-2">
                <legend className="col-form-label">Radio Buttons</legend>
              </div>
              <Col sm={10}>
                <FormGroup check>
                  <Label check className="kit__utils__control kit__utils__control__radio">
                    <Input type="radio" name="radio2" checked />
                    <span className="kit__utils__control__indicator" />
                    Option one is this and that—be sure to include why its great
                  </Label>
                </FormGroup>
                <FormGroup check>
                  <Label check className="kit__utils__control kit__utils__control__radio">
                    <Input type="radio" name="radio2" />
                    <span className="kit__utils__control__indicator" />
                    Option two can be something else and selecting it will deselect option one
                  </Label>
                </FormGroup>
                <FormGroup check disabled>
                  <Label check className="kit__utils__control kit__utils__control__radio">
                    <Input type="radio" name="radio2" disabled />
                    <span className="kit__utils__control__indicator" />
                    Option three is disabled
                  </Label>
                </FormGroup>
              </Col>
            </FormGroup>
            <FormGroup row>
              <Label for="checkbox2" sm={2}>
                Checkbox
              </Label>
              <Col sm={{ size: 10 }}>
                <FormGroup check>
                  <Label check className="kit__utils__control kit__utils__control__checkbox">
                    <Input type="checkbox" id="checkbox2" />
                    <span className="kit__utils__control__indicator" />
                    Check me out
                  </Label>
                </FormGroup>
              </Col>
            </FormGroup>
            <div className="border-top mt-4 pt-4">
              <FormGroup check row>
                <Col sm={{ size: 10, offset: 2 }}>
                  <Button color="primary" className="px-5">
                    Submit
                  </Button>
                </Col>
              </FormGroup>
            </div>
          </Form>
        </div>
      </div>
    )
  }
}

export default BootstrapFormExample
