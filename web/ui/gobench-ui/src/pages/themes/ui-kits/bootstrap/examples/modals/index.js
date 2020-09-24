import React from 'react'
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap'

class BootstrapModalsExample extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      modal: false,
      modalCentered: false,
    }

    this.toggle = this.toggle.bind(this)
    this.toggleCentered = this.toggleCentered.bind(this)
  }

  toggle() {
    this.setState(prevState => ({
      modal: !prevState.modal,
    }))
  }

  toggleCentered() {
    this.setState(prevState => ({
      modalCentered: !prevState.modalCentered,
    }))
  }

  render() {
    const { modal, modalCentered } = this.state
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Modals</strong>
        </h5>
        <div className="mb-5">
          <Button color="primary" onClick={this.toggle} className="mr-3">
            Launch demo modal
          </Button>
          <Modal isOpen={modal} toggle={this.toggle}>
            <ModalHeader toggle={this.toggle}>Modal title</ModalHeader>
            <ModalBody>
              Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor
              incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud
              exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure
              dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
              Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt
              mollit anim id est laborum.
            </ModalBody>
            <ModalFooter>
              <Button color="light" onClick={this.toggle}>
                Cancel
              </Button>{' '}
              <Button color="primary" onClick={this.toggle}>
                Do Something
              </Button>
            </ModalFooter>
          </Modal>
          <Button color="primary" onClick={this.toggleCentered}>
            Vertically centered modal
          </Button>
          <Modal
            isOpen={modalCentered}
            toggle={this.toggleCentered}
            className="modal-dialog-centered"
          >
            <ModalHeader toggle={this.toggleCentered}>Vertically centered modal</ModalHeader>
            <ModalBody>
              Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor
              incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud
              exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure
              dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
              Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt
              mollit anim id est laborum.
            </ModalBody>
            <ModalFooter>
              <Button color="light" onClick={this.toggleCentered}>
                Cancel
              </Button>{' '}
              <Button color="primary" onClick={this.toggleCentered}>
                Do Something
              </Button>
            </ModalFooter>
          </Modal>
        </div>
      </div>
    )
  }
}

export default BootstrapModalsExample
