import React from 'react'
import { Helmet } from 'react-helmet'
import { Checkbox, Calendar, Badge } from 'antd'
import style from './style.module.scss'

const AppsCalendar = () => {
  const getListData = value => {
    let listData
    switch (value.date()) {
      case 8:
        listData = [
          { type: 'warning', content: 'This is warning event.' },
          { type: 'success', content: 'This is usual event.' },
        ]
        break
      case 10:
        listData = [
          { type: 'warning', content: 'This is warning event.' },
          { type: 'success', content: 'This is usual event.' },
          { type: 'error', content: 'This is error event.' },
        ]
        break
      case 15:
        listData = [
          { type: 'warning', content: 'This is warning event' },
          { type: 'success', content: 'This is very long usual event。。....' },
          { type: 'error', content: 'This is error event 1.' },
          { type: 'error', content: 'This is error event 2.' },
          { type: 'error', content: 'This is error event 3.' },
          { type: 'error', content: 'This is error event 4.' },
        ]
        break
      default:
    }
    return listData || []
  }

  const dateCellRender = value => {
    const listData = getListData(value)
    return (
      <ul className="events">
        {listData.map(item => (
          <li key={item.content}>
            <Badge status={item.type} text={item.content} />
          </li>
        ))}
      </ul>
    )
  }

  const getMonthData = value => {
    let num
    if (value.month() === 8) {
      num = 1394
    }
    return num
  }

  const monthCellRender = value => {
    const num = getMonthData(value)
    return num ? (
      <div className="notes-month">
        <section>{num}</section>
        <span>Backlog number</span>
      </div>
    ) : null
  }

  return (
    <div>
      <Helmet title="Calendar" />
      <div className="card">
        <div className="card-body">
          <div className="d-flex flex-wrap mb-3">
            <div className="mr-5 flex-grow-1">
              <div className="text-uppercase text-dark font-weight-bold mb-3">Calendars</div>
              <div className="d-flex flex-wrap">
                <Checkbox className={style.event} checked>
                  Gmail
                </Checkbox>
                <Checkbox className={style.event}>Birthdays</Checkbox>
              </div>
            </div>
            <div className="flex-grow-1">
              <div className="text-uppercase text-dark font-weight-bold mb-3">Events</div>
              <div className="d-flex flex-wrap">
                <div className={style.event}>
                  <div className="kit__utils__donut kit__utils__donut--danger" />
                  Meeting
                </div>
                <div className={style.event}>
                  <div className="kit__utils__donut kit__utils__donut--primary" />
                  Holidays
                </div>
                <div className={style.event}>
                  <div className="kit__utils__donut kit__utils__donut--orange" />
                  Milestones
                </div>
                <div className={style.event}>
                  <div className="kit__utils__donut kit__utils__donut--success" />
                  Conference
                </div>
              </div>
            </div>
          </div>
          <Calendar dateCellRender={dateCellRender} monthCellRender={monthCellRender} />
        </div>
      </div>
    </div>
  )
}

export default AppsCalendar
