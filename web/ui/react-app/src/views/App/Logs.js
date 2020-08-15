import React, { useContext, useEffect } from 'react'
import { Tag, Table } from 'antd'
import { RootContext, EventLogContext } from '../../context'
import { statusColors } from '../../components/Status'
import { useParams } from 'react-router-dom'
import TimeAgo from 'react-timeago'

const Logs = (props) => {
  const eventLogs = useContext(EventLogContext)
  const app = useContext(RootContext)
  const { appId } = useParams()
  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
      render: text => <a>{text}</a>
    },
    {
      title: 'Level',
      dataIndex: 'level',
      key: 'level',
      render: x => (
        <Tag color={statusColors[x]}>
          {x}
        </Tag>
      )
    },
    {
      title: 'Message',
      dataIndex: 'message',
      key: 'message'
    },
    {
      title: 'Source',
      key: 'source',
      dataIndex: 'source'
    },
    {
      title: 'Logged',
      key: 'created_at',
      dataIndex: 'created_at',
      render: x => <TimeAgo date={x} />
    }
  ]


  useEffect(() => {
    if (app.getEventLogs) {
      app.getEventLogs(appId)
    }
  }, [appId])

  return (
    <Table columns={columns} dataSource={eventLogs.data} />
  )
}
export default Logs
