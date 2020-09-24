import React, { useState } from 'react'
import { Helmet } from 'react-helmet'
import { SearchOutlined } from '@ant-design/icons'
import { Input, Tooltip } from 'antd'
import PerfectScrollbar from 'react-perfect-scrollbar'
import dialogs from './data.json'
import style from './style.module.scss'

const AppsMessaging = () => {
  const [activeIndex, setActiveIndex] = useState(0)
  const { name, position, dialog, avatar } = dialogs[activeIndex]

  const changeDialog = (e, index) => {
    e.preventDefault()
    setActiveIndex(index)
  }

  return (
    <div>
      <Helmet title="Messaging" />
      <div className="row">
        <div className="col-12 col-md-3">
          <div className="mb-4">
            <Input
              prefix={<SearchOutlined style={{ color: 'rgba(0,0,0,.25)' }} />}
              placeholder="Search users..."
            />
          </div>
          <div className={style.dialogs}>
            <PerfectScrollbar>
              {dialogs.map((item, index) => (
                <a
                  href="#"
                  onClick={e => changeDialog(e, index)}
                  key={item.name}
                  className={`${style.item} ${
                    index === activeIndex ? style.current : ''
                  } d-flex flex-nowrap align-items-center`}
                >
                  <div className="kit__utils__avatar kit__utils__avatar--size46 mr-3 flex-shrink-0">
                    <img src={item.avatar} alt={item.name} />
                  </div>
                  <div className={`${style.info} flex-grow-1`}>
                    <div className="text-uppercase font-size-12 text-truncate text-gray-6">
                      {item.position}
                    </div>
                    <div className="text-dark font-size-18 font-weight-bold text-truncate">
                      {item.name}
                    </div>
                  </div>
                  <div
                    hidden={!item.unread}
                    className={`${style.unread} flex-shrink-0 align-self-start`}
                  >
                    <div className="badge badge-success">{item.unread}</div>
                  </div>
                </a>
              ))}
            </PerfectScrollbar>
          </div>
        </div>
        <div className="col-12 col-md-9">
          <div className="card">
            <div className="card-header card-header-flex align-items-center">
              <div className="d-flex flex-column justify-content-center mr-auto">
                <h5 className="mb-0 mr-2 font-size-18">
                  {name} <span className="font-size-14 text-gray-6">({position})</span>
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
              <div className="height-700">
                <PerfectScrollbar>
                  <div className="d-flex flex-column justify-content-end height-100p">
                    {dialog.map(message => (
                      <div
                        key={Math.random()}
                        className={`${style.message} ${
                          message.owner !== 'you' ? style.answer : ''
                        }`}
                      >
                        <div className={style.messageContent}>
                          <div className="text-gray-4 font-size-12 text-uppercase">
                            {message.owner}, {message.time}
                          </div>
                          <div>{message.content}</div>
                        </div>
                        <div className={`${style.messageAvatar} kit__utils__avatar`}>
                          <img
                            src={`${
                              message.owner !== 'you'
                                ? avatar
                                : 'resources/images/avatars/avatar-2.png'
                            }`}
                            alt={name}
                          />
                        </div>
                      </div>
                    ))}
                  </div>
                </PerfectScrollbar>
              </div>
              <div className="pt-2 pb-2">{name} is typing...</div>
              <div className="input-group mb-3">
                <input type="text" className="form-control" placeholder="Send message..." />
                <div className="input-group-append">
                  <button className="btn btn-primary" type="button">
                    <i className="fe fe-send align-middle" />
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default AppsMessaging
