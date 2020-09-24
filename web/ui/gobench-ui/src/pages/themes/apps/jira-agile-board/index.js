import React, { useState } from 'react'
import { Helmet } from 'react-helmet'
import { ReactSortable } from 'react-sortablejs'
import style from './style.module.scss'

const ExtraAppsJiraAgileBoard = () => {
  const [backlogItems, setBacklogItems] = useState([
    <div className={style.card} key={1}>
      <div className={style.content}>
        <div className={`${style.flag} bg-primary`} />
        <div className="d-flex flex-wrap-reverse align-items-center">
          <h5 className="text-dark font-size-18 mt-2 mr-auto">Backend Issue</h5>
          <i className="fe fe-arrow-right-circle text-danger font-size-30 flex-shrink-0" />
        </div>
        <div className="text-gray-5 mb-2">Deadline 7:00</div>
        <div className="d-flex align-items-center flex-wrap pb-2">
          <div className="kit__utils__avatar kit__utils__avatar--rounded kit__utils__avatar--size27 flex-shrink-0 mr-2">
            <img src="resources/images/avatars/2.jpg" alt="User" />
          </div>
          <div className="text-gray-5">
            <del>AWS-200</del>
          </div>
        </div>
      </div>
    </div>,
    <div className={style.card} key={2}>
      <div className={style.content}>
        <div className={`${style.flag} bg-success`} />
        <div className="d-flex flex-wrap-reverse align-items-center">
          <h5 className="text-dark font-size-18 mt-2 mr-auto">Prepare Activity List</h5>
          <i className="fe fe-arrow-left-circle text-success font-size-30 flex-shrink-0" />
        </div>
        <div className="text-gray-5 mb-2">Deadline 7:00</div>
        <div className="d-flex align-items-center flex-wrap pb-2">
          <div className="kit__utils__avatar kit__utils__avatar--rounded kit__utils__avatar--size27 flex-shrink-0 mr-2">
            <img src="resources/images/avatars/1.jpg" alt="User" />
          </div>
          <div className="text-gray-5">
            <del>AWS-200</del>
          </div>
        </div>
      </div>
    </div>,
    <div className={style.card} key={3}>
      <div className={style.content}>
        <div className={`${style.flag} bg-danger`} />
        <div className="d-flex flex-wrap-reverse align-items-center">
          <h5 className="text-dark font-size-18 mt-2 mr-auto">Prepare Activity List</h5>
          <i className="fe fe-arrow-left-circle text-success font-size-30 flex-shrink-0" />
        </div>
        <div className="text-gray-5 mb-2">Deadline 7:00</div>
        <div className="d-flex align-items-center flex-wrap pb-2">
          <div className="kit__utils__avatar kit__utils__avatar--rounded kit__utils__avatar--size27 flex-shrink-0 mr-2">
            <img src="resources/images/avatars/4.jpg" alt="Mary Stanform" />
          </div>
          <div className="text-gray-5">
            <del>AWS-200</del>
          </div>
        </div>
      </div>
    </div>,
  ])

  const [todoItems, setTodoItems] = useState([
    <div className={style.card} key={4}>
      <div className={style.content}>
        <div className={`${style.flag} bg-danger`} />
        <div className="d-flex flex-wrap-reverse align-items-center">
          <h5 className="text-dark font-size-18 mt-2 mr-auto">Prepare Activity List</h5>
          <i className="fe fe-arrow-left-circle text-success font-size-30 flex-shrink-0" />
        </div>
        <div className="text-gray-5 mb-2">Deadline 7:00</div>
        <div className="d-flex align-items-center flex-wrap pb-2">
          <div className="kit__utils__avatar kit__utils__avatar--rounded kit__utils__avatar--size27 flex-shrink-0 mr-2">
            <img src="resources/images/avatars/4.jpg" alt="Mary Stanform" />
          </div>
          <div className="text-gray-5">
            <del>AWS-200</del>
          </div>
        </div>
      </div>
    </div>,
  ])

  const [inprogressItems, setInprogressItems] = useState([
    <div className={style.card} key={1}>
      <div className={style.content}>
        <div className={`${style.flag} bg-success`} />
        <div className="d-flex flex-wrap-reverse align-items-center">
          <h5 className="text-dark font-size-18 mt-2 mr-auto">Prepare Activity List</h5>
          <i className="fe fe-arrow-left-circle text-success font-size-30 flex-shrink-0" />
        </div>
        <div className="text-gray-5 mb-2">Deadline 7:00</div>
        <div className="d-flex align-items-center flex-wrap pb-2">
          <div className="kit__utils__avatar kit__utils__avatar--rounded kit__utils__avatar--size27 flex-shrink-0 mr-2">
            <img src="resources/images/avatars/1.jpg" alt="User" />
          </div>
          <div className="text-gray-5">
            <del>AWS-200</del>
          </div>
        </div>
      </div>
    </div>,
  ])

  const [completed, setCompleted] = useState([])

  return (
    <div>
      <Helmet title="Jira Agile Board" />
      <div className="d-flex align-items-center mb-4">
        <div className="kit__utils__avatarGroup mr-4 flex-shrink-0">
          <div className="kit__utils__avatar kit__utils__avatar--size46 kit__utils__avatar--rounded">
            <img src="resources/images/avatars/1.jpg" alt="User 1" />
          </div>
          <div className="kit__utils__avatar kit__utils__avatar--size46 kit__utils__avatar--rounded">
            <img src="resources/images/avatars/2.jpg" alt="User 2" />
          </div>
          <div className="kit__utils__avatar kit__utils__avatar--size46 kit__utils__avatar--rounded">
            <img src="resources/images/avatars/3.jpg" alt="User 3" />
          </div>
          <div className="kit__utils__avatar kit__utils__avatar--size46 kit__utils__avatar--rounded">
            <img src="resources/images/avatars/4.jpg" alt="User 4" />
          </div>
          <button type="button" className="kit__utils__avatarGroupAdd">
            <i className="fe fe-plus" />
          </button>
        </div>
        <a className="mr-4" href="#" onClick={e => e.preventDefault()}>
          Only My Issues
        </a>
        <div>Recently Updated</div>
      </div>
      <div className="row">
        <div className="col-lg-3 col-md-6">
          <div className="card bg-light py-3 px-2">
            <h3 className="font-weight-bold text-dark font-size-18 mb-3">Backlog</h3>
            <ReactSortable
              group="issues"
              style={{ minHeight: 500 }}
              list={backlogItems}
              setList={setBacklogItems}
            >
              {backlogItems.map(item => item)}
            </ReactSortable>
          </div>
        </div>
        <div className="col-lg-3 col-md-6">
          <div className="card bg-light py-3 px-2">
            <h3 className="font-weight-bold text-dark font-size-18 mb-3">To Do</h3>
            <ReactSortable
              group="issues"
              style={{ minHeight: 500 }}
              list={todoItems}
              setList={setTodoItems}
            >
              {todoItems.map(item => item)}
            </ReactSortable>
          </div>
        </div>
        <div className="col-lg-3 col-md-6">
          <div className="card bg-light py-3 px-2">
            <h3 className="font-weight-bold text-dark font-size-18 mb-3">In Progress</h3>
            <ReactSortable
              group="issues"
              style={{ minHeight: 500 }}
              list={inprogressItems}
              setList={setInprogressItems}
            >
              {inprogressItems.map(item => item)}
            </ReactSortable>
          </div>
        </div>
        <div className="col-lg-3 col-md-6">
          <div className="card bg-light py-3 px-2">
            <h3 className="font-weight-bold text-dark font-size-18 mb-3">Completed</h3>
            <ReactSortable
              group="issues"
              style={{ minHeight: 500 }}
              list={completed}
              setList={setCompleted}
            >
              {completed.map(item => item)}
            </ReactSortable>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ExtraAppsJiraAgileBoard
