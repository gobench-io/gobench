/* eslint-disable */
import React from 'react'
import { Tree } from 'antd'

const { TreeNode } = Tree

const treeData = [
  {
    title: 'parent 1',
    key: '0-0',
    children: [
      {
        title: 'parent 1-0',
        key: '0-0-0',
        disabled: true,
        children: [
          {
            title: 'leaf',
            key: '0-0-0-0',
            disableCheckbox: true,
          },
          {
            title: 'leaf',
            key: '0-0-0-1',
          },
        ],
      },
      {
        title: 'parent 1-1',
        key: '0-0-1',
        children: [{ title: <span style={{ color: '#1890ff' }}>sss</span>, key: '0-0-1-0' }],
      },
    ],
  },
]

class AntdTreeExample extends React.Component {
  render() {
    return (
      <div>
        <div className="row">
          <div className="col-lg-4">
            <h5 className="mb-3">
              <strong>Checkable</strong>
            </h5>
            <Tree
              checkable
              defaultExpandedKeys={['0-0-0', '0-0-1']}
              defaultSelectedKeys={['0-0-0', '0-0-1']}
              defaultCheckedKeys={['0-0-0', '0-0-1']}
              treeData={treeData}
            />
          </div>
          <div className="col-lg-4">
            <h5 className="mb-3">
              <strong>Basic</strong>
            </h5>
            <Tree
              defaultExpandedKeys={['0-0-0', '0-0-1']}
              defaultSelectedKeys={['0-0-0', '0-0-1']}
              defaultCheckedKeys={['0-0-0', '0-0-1']}
              treeData={treeData}
            />
          </div>
          <div className="col-lg-4">
            <h5 className="mb-3">
              <strong>With Lines</strong>
            </h5>
            <Tree clasName="component-col" showLine defaultExpandedKeys={['0-0-0']}>
              <TreeNode title="parent 1" key="0-0">
                <TreeNode title="parent 1-0" key="0-0-0">
                  <TreeNode title="leaf" key="0-0-0-0" />
                  <TreeNode title="leaf" key="0-0-0-1" />
                  <TreeNode title="leaf" key="0-0-0-2" />
                </TreeNode>
                <TreeNode title="parent 1-1" key="0-0-1">
                  <TreeNode title="leaf" key="0-0-1-0" />
                </TreeNode>
                <TreeNode title="parent 1-2" key="0-0-2">
                  <TreeNode title="leaf" key="0-0-2-0" />
                  <TreeNode title="leaf" key="0-0-2-1" />
                </TreeNode>
              </TreeNode>
            </Tree>
          </div>
        </div>
      </div>
    )
  }
}

export default AntdTreeExample
