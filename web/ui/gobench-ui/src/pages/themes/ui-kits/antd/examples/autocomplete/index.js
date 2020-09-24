/* eslint-disable */
import React from 'react'
import { AutoComplete, Input } from 'antd'

const { Option } = AutoComplete

function onSelect(value) {
  console.log('onSelect', value)
}

class AntdAutoCompleteExample extends React.Component {
  state = {
    value: '',
    dataSource: [],
    result: [],
  }

  onSearch = searchText => {
    this.setState({
      dataSource: !searchText ? [] : [searchText, searchText.repeat(2), searchText.repeat(3)],
    })
  }

  onChange = value => {
    this.setState({ value })
  }

  handleSearch = value => {
    let result
    if (!value || value.indexOf('@') >= 0) {
      result = []
    } else {
      result = ['gmail.com', '163.com', 'qq.com'].map(domain => `${value}@${domain}`)
    }
    this.setState({ result })
  }

  render() {
    const { result } = this.state
    const children = result.map(email => <Option key={email}>{email}</Option>)

    return (
      <div>
        <div className="row">
          <div className="col-lg-6 mb-5">
            <h5 className="mb-3">
              <strong>Basic</strong>
            </h5>
            <AutoComplete
              dataSource={this.state.dataSource}
              style={{ width: 200 }}
              onSelect={onSelect}
              onSearch={this.onSearch}
              placeholder="input here"
            />
            <br />
            <br />
            <AutoComplete
              value={this.state.value}
              dataSource={this.state.dataSource}
              style={{ width: 200 }}
              onSelect={onSelect}
              onSearch={this.onSearch}
              onChange={this.onChange}
              placeholder="control mode"
            />
          </div>
          <div className="col-lg-6 mb-5">
            <h5 className="mb-3">
              <strong>Customized</strong>
            </h5>
            <AutoComplete
              style={{ width: 200 }}
              onSearch={this.handleSearch}
              placeholder="input here"
            >
              {children}
            </AutoComplete>
          </div>
        </div>
      </div>
    )
  }
}

export default AntdAutoCompleteExample
