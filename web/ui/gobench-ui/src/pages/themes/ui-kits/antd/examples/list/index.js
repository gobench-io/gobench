/* eslint-disable */
import React from 'react'
import { List, Avatar, Typography } from 'antd'

const dataList = [
  'Racing car sprays burning fuel into crowd.',
  'Japanese princess to wed commoner.',
  'Australian walks 100km after outback crash.',
  'Man charged over missing wedding girl.',
  'Los Angeles battles huge wildfires.',
]

const data = [
  {
    title: 'Ant Design Title 1',
  },
  {
    title: 'Ant Design Title 2',
  },
  {
    title: 'Ant Design Title 3',
  },
  {
    title: 'Ant Design Title 4',
  },
]

class AntdListExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <List
            header={<div>Header</div>}
            footer={<div>Footer</div>}
            bordered
            dataSource={dataList}
            renderItem={item => (
              <List.Item>
                <Typography.Text code>[#]</Typography.Text> {item}
              </List.Item>
            )}
          />
        </div>
        <h5 className="mb-3">
          <strong>With Avatar</strong>
        </h5>
        <div className="mb-5">
          <List
            itemLayout="horizontal"
            dataSource={data}
            renderItem={item => (
              <List.Item>
                <List.Item.Meta
                  avatar={
                    <Avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
                  }
                  title={<a href="https://ant.design">{item.title}</a>}
                  description="Ant Design, a design language for background applications, is refined by Ant UED Team"
                />
              </List.Item>
            )}
          />
        </div>
      </div>
    )
  }
}

export default AntdListExample
