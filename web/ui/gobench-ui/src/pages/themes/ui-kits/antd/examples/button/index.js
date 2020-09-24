/* eslint-disable */
import React from 'react'
import { DownloadOutlined, LeftOutlined, RightOutlined, SearchOutlined } from '@ant-design/icons'
import { Button, Radio } from 'antd'

class AntdButtonExample extends React.Component {
  state = {
    size: 'large',
  }

  handleSizeChange = e => {
    this.setState({ size: e.target.value })
  }

  render() {
    const { size } = this.state
    return (
      <div className="row" id="example-buttons">
        <div className="col-lg-6">
          <h5 className="mb-3">
            <strong>Default</strong>
          </h5>
          <div className="mb-5">
            <Button type="primary" className="mr-3 mb-3">
              Primary
            </Button>
            <Button className="mr-3 mb-3">Default</Button>
            <Button type="dashed" className="mr-3 mb-3">
              Dashed
            </Button>
            <Button type="danger" className="mr-3 mb-3">
              Danger
            </Button>
            <Button type="link" className="mr-3 mb-3">
              Link
            </Button>
          </div>
          <h5 className="mb-3">
            <strong>Size</strong>
          </h5>
          <div className="mb-5">
            <Radio.Group value={size} onChange={this.handleSizeChange}>
              <Radio.Button value="large">Large</Radio.Button>
              <Radio.Button value="default">Default</Radio.Button>
              <Radio.Button value="small">Small</Radio.Button>
            </Radio.Group>
            <br />
            <br />
            <Button type="primary" size={size} className="mr-3 mb-3">
              Primary
            </Button>
            <Button size={size} className="mr-3 mb-3">
              Normal
            </Button>
            <Button type="dashed" size={size} className="mr-3 mb-3">
              Dashed
            </Button>
            <Button type="danger" size={size} className="mr-3 mb-3">
              Danger
            </Button>
            <Button type="link" size={size} className="mr-3 mb-3">
              Link
            </Button>
            <br />
            <Button type="primary" icon={<DownloadOutlined />} size={size} className="mr-3 mb-3" />
            <Button
              type="primary"
              shape="circle"
              icon={<DownloadOutlined />}
              size={size}
              className="mr-3 mb-3"
            />
            <Button
              type="primary"
              shape="round"
              icon={<DownloadOutlined />}
              size={size}
              className="mr-3 mb-3"
            />
            <Button
              type="primary"
              shape="round"
              icon={<DownloadOutlined />}
              size={size}
              className="mr-3 mb-3"
            >
              Download
            </Button>
            <Button type="primary" icon={<DownloadOutlined />} size={size} className="mr-3 mb-3">
              Download
            </Button>
            <br />
            <Button.Group size={size} className="mr-3 mb-3">
              <Button type="primary">
                <LeftOutlined />
                Backward
              </Button>
              <Button type="primary">
                Forward
                <RightOutlined />
              </Button>
            </Button.Group>
          </div>
          <h5 className="mb-3">
            <strong>Disabled</strong>
          </h5>
          <div className="mb-5">
            <Button type="primary" className="mr-3 mb-3">
              Primary
            </Button>
            <Button type="primary" disabled className="mr-3 mb-3">
              Primary(disabled)
            </Button>
            <br />
            <Button className="mr-3 mb-3">Default</Button>
            <Button disabled className="mr-3 mb-3">
              Default(disabled)
            </Button>
            <br />
            <Button type="dashed" className="mr-3 mb-3">
              Dashed
            </Button>
            <Button type="dashed" disabled className="mr-3 mb-3">
              Dashed(disabled)
            </Button>
            <br />
            <Button type="link" className="mr-3 mb-3">
              Link
            </Button>
            <Button type="link" disabled className="mr-3 mb-3">
              Link(disabled)
            </Button>
          </div>
        </div>
        <div className="col-lg-6">
          <h5 className="mb-3">
            <strong>Default</strong>
          </h5>
          <div className="mb-5">
            <Button type="primary" shape="circle" icon={<SearchOutlined />} className="mr-3 mb-3" />
            <Button type="primary" shape="circle" className="mr-3 mb-3">
              A
            </Button>
            <Button type="primary" icon={<SearchOutlined />} className="mr-3 mb-3">
              Search
            </Button>
            <Button shape="circle" icon={<SearchOutlined />} className="mr-3 mb-3" />
            <Button icon={<SearchOutlined />} className="mr-3 mb-3">
              Search
            </Button>
            <br />
            <Button shape="circle" icon={<SearchOutlined />} className="mr-3 mb-3" />
            <Button icon={<SearchOutlined />} className="mr-3 mb-3">
              Search
            </Button>
            <Button type="dashed" shape="circle" icon={<SearchOutlined />} className="mr-3 mb-3" />
            <Button type="dashed" icon={<SearchOutlined />} className="mr-3 mb-3">
              Search
            </Button>
          </div>
          <h5 className="mb-3">
            <strong>Loading</strong>
          </h5>
          <div className="mb-5">
            <Button type="primary" loading className="mr-3 mb-3">
              Loading
            </Button>
            <Button type="primary" size="small" loading className="mr-3 mb-3">
              Loading
            </Button>
            <br />
            <Button type="primary" loading className="mr-3 mb-3" />
            <Button type="primary" shape="circle" loading className="mr-3 mb-3" />
            <Button type="danger" shape="round" loading className="mr-3 mb-3" />
          </div>
          <h5 className="mb-3">
            <strong>Block Buttons</strong>
          </h5>
          <div className="mb-5">
            <Button type="primary" block className="mb-3">
              Primary
            </Button>
            <Button block className="mb-3">
              Default
            </Button>
            <Button type="dashed" block className="mb-3">
              Dashed
            </Button>
            <Button type="danger" block className="mb-3">
              Danger
            </Button>
            <Button type="link" block className="mb-3">
              Link
            </Button>
          </div>
        </div>
      </div>
    )
  }
}

export default AntdButtonExample
