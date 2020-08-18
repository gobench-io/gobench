import React, { useContext } from 'react'
import { get } from 'lodash'
import TimeAgo from 'react-timeago'
import { useHistory } from 'react-router-dom'
import Status from '../../components/Status'
import { RootContext, SpinnerContext } from '../../context'
import { Button, Popconfirm } from 'antd'

const Applications = () => {
  const history = useHistory()
  const app = useContext(RootContext)
  const applications = app.apps || []
  const isFetching = useContext(SpinnerContext)
  const now = new Date()
  const timestamp = `${now.getUTCFullYear()}-${now.getUTCMonth()}-${now.getUTCDate()}-${now.getUTCHours()}-${now.getUTCMinutes()}-${now.getUTCSeconds()}`
  return (
    <div>
      <div className='card'>
        <div className='applications-list-header'>
          <h2>Applications</h2>
          <button
            onClick={() => history.push('/application-create')}
            className='btn btn-primary'
          >
            Create application
          </button>
        </div>
        <div className='applications-body'>
          {isFetching ? (
            <div>Loading applications</div>
          ) : (
            <table className='applications-table'>
              <thead>
                <tr>
                  <th>Application Name</th>
                  <th>Status</th>
                  <th>Created At</th>
                  <th />
                </tr>
              </thead>
              <tbody>
                {!isFetching && applications.length === 0 ? (
                  <tr>
                    <td colSpan={3}>There is no applications</td>
                  </tr>
                ) : (
                  applications.map((application) => {
                    const status = get(application, 'status', '')
                    return (
                      <tr key={application.id}>
                        <td style={{ width: '34%%' }}>
                          {application.name || '-'}
                        </td>
                        <td style={{ width: '15%' }}>
                          <Status status={application.status} />
                        </td>
                        <td style={{ width: '15%' }}>
                          <TimeAgo date={application.created_at} />
                        </td>
                        <td style={{ width: '33%' }}>
                          <div style={{ float: 'right' }}>
                            {['finished', 'running', 'cancel'].includes(
                              status
                            ) && (
                              <Button
                                type='primary'
                                onClick={() =>
                                  history.push(`/application/${application.id}`)}
                              >
                                Detail
                              </Button>
                            )}
                            {['finished'].includes(status) && (
                              <Button
                                style={{ marginLeft: 5 }}
                                type='default'
                                onClick={() =>
                                  app.submitCreate({
                                    ...application,
                                    name: `${application.name}-${timestamp}`
                                  })}
                              >
                                Clone
                              </Button>
                            )}
                            {['running', 'pending'].includes(status) && (
                              <Popconfirm
                                title={`Are you sure cancel application ${application.name}?`}
                                onConfirm={() =>
                                  app.cancelRunApplication(application.id)}
                                okText='Yes'
                                cancelText='No'
                              >
                                <Button
                                  type='dashed'
                                  style={{ marginLeft: 5 }}
                                  danger
                                >
                                Cancel
                                </Button>
                              </Popconfirm>
                            )}
                            {['finished', 'pending', 'error', 'cancel'].includes(status) && (
                              <Popconfirm
                                title={`Are you sure delete application ${application.name}?`}
                                onConfirm={() =>
                                  app.deleteApplication({
                                    ...application,
                                    name: `${application.name}-${timestamp}`
                                  })}
                                okText='Yes'
                                cancelText='No'
                              >
                                <Button
                                  type='primary'
                                  style={{ marginLeft: 5 }}
                                  danger
                                >
                                Delete
                                </Button>
                              </Popconfirm>
                            )}
                          </div>
                        </td>
                      </tr>
                    )
                  })
                )}
              </tbody>
            </table>
          )}
        </div>
      </div>
    </div>
  )
}
export default Applications
