/* eslint-disable */
import React from 'react'
import { Transfer, Switch } from 'antd'

const mockData = []
for (let i = 0; i < 20; i++) {
  mockData.push({
    key: i.toString(),
    title: `content${i + 1}`,
    description: `description of content${i + 1}`,
    disabled: i % 3 < 1,
  })
}

const oriTargetKeys = mockData.filter(item => +item.key % 3 > 1).map(item => item.key)

class AntdTransferExample extends React.Component {
  state = {
    targetKeys: oriTargetKeys,
    selectedKeys: [],
    disabled: false,
  }

  handleChange = (nextTargetKeys, direction, moveKeys) => {
    this.setState({ targetKeys: nextTargetKeys })

    console.log('targetKeys: ', nextTargetKeys)
    console.log('direction: ', direction)
    console.log('moveKeys: ', moveKeys)
  }

  handleSelectChange = (sourceSelectedKeys, targetSelectedKeys) => {
    this.setState({ selectedKeys: [...sourceSelectedKeys, ...targetSelectedKeys] })

    console.log('sourceSelectedKeys: ', sourceSelectedKeys)
    console.log('targetSelectedKeys: ', targetSelectedKeys)
  }

  handleScroll = (direction, e) => {
    console.log('direction:', direction)
    console.log('target:', e.target)
  }

  handleDisable = disabled => {
    this.setState({ disabled })
  }

  render() {
    const { targetKeys, selectedKeys, disabled } = this.state
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Transfer
            dataSource={mockData}
            titles={['Source', 'Target']}
            targetKeys={targetKeys}
            selectedKeys={selectedKeys}
            onChange={this.handleChange}
            onSelectChange={this.handleSelectChange}
            onScroll={this.handleScroll}
            render={item => item.title}
            disabled={disabled}
          />
          <Switch
            unCheckedChildren="enabled"
            checkedChildren="disabled"
            checked={disabled}
            onChange={this.handleDisable}
            style={{ marginTop: 16 }}
          />
        </div>
      </div>
    )
  }
}

export default AntdTransferExample
