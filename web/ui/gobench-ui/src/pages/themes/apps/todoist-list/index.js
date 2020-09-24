import React, { useState } from 'react'
import { Helmet } from 'react-helmet'
import PerfectScrollbar from 'react-perfect-scrollbar'
import SortableTree, { changeNodeAtPath } from 'react-sortable-tree'
import { SearchOutlined } from '@ant-design/icons'
import { Input, Tooltip, Checkbox } from 'antd'
import Table6 from 'components/kit/widgets/Tables/6'
import Chart4 from 'components/kit/widgets/Charts/4'
import Chart4v1 from 'components/kit/widgets/Charts/4v1'
import Chart4v2 from 'components/kit/widgets/Charts/4v2'
import style from './style.module.scss'

const ExtraAppsTodoistList = () => {
  const taskInput = React.createRef()
  const treeDataDefault = [
    { name: 'Level 1' },
    {
      name: 'Level 2',
      expanded: true,
      children: [{ name: 'Level 2' }, { name: 'Level 2' }, { name: 'Level 2' }],
    },
  ]
  const completedDefault = [
    { name: 'Level 0', checked: true },
    { name: 'Level 0', checked: true },
    { name: 'Level 0', checked: true },
  ]

  const [treeData, setTreeData] = useState(treeDataDefault)
  const [completed, setCompleted] = useState(completedDefault)
  const [hideInput, setHideInput] = useState(true)

  const toggleInput = () => {
    setHideInput(!hideInput)
  }

  const addTask = e => {
    const task = e.target.value

    if (e.which === 13 && task !== '') {
      const treeDataProcessed = treeData.concat({
        name: task,
      })

      setTreeData(treeDataProcessed)
      taskInput.current.value = ''
    }
  }

  const getNodeKey = ({ treeIndex }) => treeIndex

  return (
    <div>
      <Helmet title="Todoist List" />
      <div className="row">
        <div className="col-12 col-md-3">
          <div className="mb-4">
            <Input
              prefix={<SearchOutlined style={{ color: 'rgba(0,0,0,.25)' }} />}
              placeholder="Search mail..."
            />
          </div>
          <div className={style.categories}>
            <div className="d-flex flex-column">
              <div
                className={` ${style.category} ${style.title} text-dark font-size-18 font-weight-bold`}
              >
                <span className="text-truncate">Overall</span>
              </div>
              <a
                href="#"
                onClick={e => e.preventDefault()}
                className={`${style.category} text-dark font-size-18`}
              >
                <span className="text-truncate">Today</span>
              </a>
              <a
                href="#"
                onClick={e => e.preventDefault()}
                className={`${style.category} text-dark font-size-18`}
              >
                <span className="text-truncate">Next 7 days</span>
              </a>
              <div
                className={` ${style.category} ${style.title} text-dark font-size-18 font-weight-bold`}
              >
                <span className="text-truncate">To Do Lists</span>
              </div>
              <a
                href="#"
                onClick={e => e.preventDefault()}
                className={`${style.category} ${style.current} text-dark font-size-18`}
              >
                <span className="text-truncate">
                  Welcome{' '}
                  <span role="img" aria-label="hello">
                    &#128075;
                  </span>
                </span>
              </a>
              <a
                href="#"
                onClick={e => e.preventDefault()}
                className={`${style.category} text-dark font-size-18`}
              >
                <span className="text-truncate">Homework (1)</span>
              </a>
              <a
                href="#"
                onClick={e => e.preventDefault()}
                className={`${style.category} text-dark font-size-18`}
              >
                <span className="text-truncate">Fitness</span>
              </a>
            </div>
          </div>
        </div>
        <div className="col-md-9">
          <div className="card">
            <div className="card-header card-header-flex align-items-center">
              <div className="d-flex flex-column justify-content-center mr-auto">
                <h5 className="mb-0">
                  Welcome{' '}
                  <span role="img" aria-label="hello">
                    &#128075;
                  </span>
                </h5>
              </div>
              <div>
                <Tooltip placement="top" title="Unlock Account">
                  <a
                    href="#"
                    onClick={e => e.preventDefault()}
                    className="btn btn-sm btn-light mr-2"
                  >
                    <i className="fe fe-unlock" />
                  </a>
                </Tooltip>
                <Tooltip placement="top" title="Mark as important">
                  <a
                    href="#"
                    onClick={e => e.preventDefault()}
                    className="btn btn-sm btn-light mr-2"
                  >
                    <i className="fe fe-star" />
                  </a>
                </Tooltip>
                <Tooltip placement="top" title="Delete user">
                  <a href="#" onClick={e => e.preventDefault()} className="btn btn-sm btn-light">
                    <i className="fe fe-trash" />
                  </a>
                </Tooltip>
              </div>
            </div>
            <div className="card-body">
              <h6 className="text-uppercase text-dark font-size-18 font-weight-bold mb-2">
                Current tasks
              </h6>
              <p className="mb-3">Welocme to Todoist! Let&apos;s get you started with a few tips</p>
              <div className="height-400">
                <PerfectScrollbar>
                  <div className="height-400">
                    <SortableTree
                      treeData={treeData}
                      onChange={tree => setTreeData(tree)}
                      generateNodeProps={({ node, path }) => ({
                        title: !node.children ? (
                          <Checkbox
                            checked={node.checked}
                            onChange={event => {
                              const { checked } = event.target
                              setTreeData(
                                changeNodeAtPath({
                                  treeData,
                                  path,
                                  getNodeKey,
                                  newNode: { ...node, checked },
                                }),
                              )
                            }}
                          >
                            {node.name}
                          </Checkbox>
                        ) : (
                          <span>{node.name}:</span>
                        ),
                      })}
                    />
                  </div>
                </PerfectScrollbar>
              </div>
              <button
                type="button"
                className="btn btn-primary btn-with-addon text-nowrap"
                onClick={toggleInput}
              >
                <span className="btn-addon">
                  <i className="btn-addon-icon fe fe-plus-circle" />
                </span>
                Add Task
              </button>
              <input
                hidden={hideInput}
                className="form-control mt-3"
                placeholder="Add task here and press enter..."
                type="text"
                onKeyPress={e => addTask(e)}
                ref={taskInput}
              />
              <h6 className="text-uppercase text-dark font-size-18 font-weight-bold mb-2 mt-4">
                Completed
              </h6>
              <p className="mb-3">Welocme to Todoist! Let&apos;s get you started with a few tips</p>
              <div className="height-200">
                <PerfectScrollbar>
                  <div className="height-200">
                    <SortableTree
                      treeData={completed}
                      onChange={tree => setCompleted(tree)}
                      generateNodeProps={({ node }) => ({
                        title: !node.children ? (
                          <Checkbox checked={node.checked}>{node.name}</Checkbox>
                        ) : (
                          <span>{node.name}:</span>
                        ),
                      })}
                    />
                  </div>
                </PerfectScrollbar>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-lg-4">
              <div className="card">
                <div className="card-body">
                  <Chart4 />
                </div>
              </div>
            </div>
            <div className="col-lg-4">
              <div className="card">
                <div className="card-body">
                  <Chart4v1 />
                </div>
              </div>
            </div>
            <div className="col-lg-4">
              <div className="card">
                <div className="card-body">
                  <Chart4v2 />
                </div>
              </div>
            </div>
          </div>
          <div className="card">
            <div className="card-body">
              <Table6 />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ExtraAppsTodoistList
