/* eslint-disable */
import React from 'react'
import { Anchor } from 'antd'

const { Link } = Anchor

class AntdAnchorExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <Anchor>
          <Link title="Basic demo" />
          <Link title="Static demo" />
          <Link title="Basic demo with Target" />
          <Link title="API">
            <Link title="Anchor Props" />
            <Link title="Link Props" />
          </Link>
        </Anchor>
      </div>
    )
  }
}

export default AntdAnchorExample
