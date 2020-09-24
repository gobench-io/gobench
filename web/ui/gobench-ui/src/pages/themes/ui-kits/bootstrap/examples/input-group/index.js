import React from 'react'
import { InputGroup, InputGroupAddon, InputGroupText, Input } from 'reactstrap'

class BootstrapInputGroupExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Input Addons</strong>
        </h5>
        <InputGroup>
          <InputGroupAddon addonType="prepend">@</InputGroupAddon>
          <Input placeholder="username" />
        </InputGroup>
        <br />
        <InputGroup>
          <InputGroupAddon addonType="prepend">
            <InputGroupText>
              <Input addon type="checkbox" aria-label="Checkbox for following text input" />
            </InputGroupText>
          </InputGroupAddon>
          <Input placeholder="Check it out" />
        </InputGroup>
        <br />
        <InputGroup>
          <Input placeholder="username" />
          <InputGroupAddon addonType="append">
            <InputGroupText>@example.com</InputGroupText>
          </InputGroupAddon>
        </InputGroup>
        <br />
        <InputGroup>
          <InputGroupAddon addonType="prepend">
            <InputGroupText>$</InputGroupText>
            <InputGroupText>$</InputGroupText>
          </InputGroupAddon>
          <Input placeholder="Dolla dolla billz yo!" />
          <InputGroupAddon addonType="append">
            <InputGroupText>$</InputGroupText>
            <InputGroupText>$</InputGroupText>
          </InputGroupAddon>
        </InputGroup>
        <br />
        <InputGroup>
          <InputGroupAddon addonType="prepend">$</InputGroupAddon>
          <Input placeholder="Amount" min={0} max={100} type="number" step="1" />
          <InputGroupAddon addonType="append">.00</InputGroupAddon>
        </InputGroup>
      </div>
    )
  }
}

export default BootstrapInputGroupExample
