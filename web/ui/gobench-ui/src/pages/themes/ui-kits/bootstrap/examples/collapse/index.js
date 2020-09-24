import React from 'react'
import { Collapse, Button, CardBody, Card, UncontrolledCollapse } from 'reactstrap'

class BootstrapCollapseExample extends React.Component {
  constructor(props) {
    super(props)
    this.toggle = this.toggle.bind(this)
    this.state = { collapse: false }
  }

  toggle() {
    this.setState(state => ({ collapse: !state.collapse }))
  }

  render() {
    const { collapse } = this.state
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Collapse</strong>
        </h5>
        <div className="mb-5">
          <div className="mb-5">
            <div className="accordion">
              <div className="card">
                <div className="card-header" id="headingOne">
                  <div className="card-title">
                    <span className="accordion-indicator pull-right">
                      <i className="plus fe fe-plus" />
                      <i className="minus fe fe-minus" />
                    </span>
                    <a>Collapsible Group Item #1</a>
                  </div>
                </div>
                <UncontrolledCollapse toggler="#headingOne">
                  <div className="card-body">
                    Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry
                    richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor
                    brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor,
                    sunt aliqua put a bird on it squid single-origin coffee nulla assumenda
                    shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson
                    cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo.
                    Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt
                    you probably havent heard of them accusamus labore sustainable VHS.
                  </div>
                </UncontrolledCollapse>
              </div>
              <div className="card">
                <div className="card-header" id="headingTwo">
                  <div className="card-title">
                    <span className="accordion-indicator pull-right">
                      <i className="plus fe fe-plus" />
                      <i className="minus fe fe-minus" />
                    </span>
                    <a>Collapsible Group Item #2</a>
                  </div>
                </div>
                <UncontrolledCollapse toggler="#headingTwo">
                  <div className="card-body">
                    Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry
                    richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor
                    brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor,
                    sunt aliqua put a bird on it squid single-origin coffee nulla assumenda
                    shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson
                    cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo.
                    Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt
                    you probably havent heard of them accusamus labore sustainable VHS.
                  </div>
                </UncontrolledCollapse>
              </div>
              <div className="card">
                <div className="card-header" id="headingThree">
                  <div className="card-title">
                    <span className="accordion-indicator pull-right">
                      <i className="plus fe fe-plus" />
                      <i className="minus fe fe-minus" />
                    </span>
                    <a>Collapsible Group Item #3</a>
                  </div>
                </div>
                <UncontrolledCollapse toggler="#headingThree">
                  <div className="card-body">
                    Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry
                    richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor
                    brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor,
                    sunt aliqua put a bird on it squid single-origin coffee nulla assumenda
                    shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson
                    cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo.
                    Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt
                    you probably havent heard of them accusamus labore sustainable VHS.
                  </div>
                </UncontrolledCollapse>
              </div>
            </div>
          </div>
        </div>
        <h5 className="mb-4">
          <strong>YTarget Collapse</strong>
        </h5>
        <div>
          <Button color="primary" onClick={this.toggle} style={{ marginBottom: '1rem' }}>
            Toggle
          </Button>
          <Collapse isOpen={collapse}>
            <Card>
              <CardBody>
                Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry
                richardson ad squid. Nihil anim keffiyeh helvetica, craft beer labore wes anderson
                cred nesciunt sapiente ea proident.
              </CardBody>
            </Card>
          </Collapse>
        </div>
      </div>
    )
  }
}

export default BootstrapCollapseExample
