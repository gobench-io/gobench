/* eslint-disable */
import React from 'react'
import { Affix, Button } from 'antd'

class AntdAffixExample extends React.Component {
  state = {
    top: 10,
    bottom: 10,
  }

  render() {
    return (
      <div className="height-700 d-flex align-items-center">
        <Affix offsetTop={this.state.top}>
          <Button
            type="primary"
            onClick={() => {
              this.setState({
                top: this.state.top + 10,
              })
            }}
          >
            Affix top
          </Button>
        </Affix>
        <br />
        <Affix offsetBottom={this.state.bottom}>
          <Button
            type="primary"
            onClick={() => {
              this.setState({
                bottom: this.state.bottom + 10,
              })
            }}
          >
            Affix bottom
          </Button>
        </Affix>
      </div>
    )
  }
}

export default AntdAffixExample
