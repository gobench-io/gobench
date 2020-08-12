import React, { useEffect, useState, useContext } from 'react'
import { get } from 'lodash'
import GoBenchAPI from '../../api/gobench'
import { useHistory } from 'react-router-dom'
import { useInterval, INTERVAL } from '../../realtimeHelpers'
import Status from '../../components/Status'
import { ApplicationsListContext, SpinnerContext } from '../../context'

const Applications = () => {
  const history = useHistory()
const	app = useContext(ApplicationsListContext)
	const applications = app.apps || []
	const isFetching = useContext(SpinnerContext)
  return <div >
  <div className='card'>
      <div className='applications-list-header'>
        <h2>Applications</h2>
        <button
          onClick={() => history.push('/application/create')}
          className='btn btn-primary'
        >
          Create application
        </button>
      </div>
      <div className='applications-body'>
        {
          isFetching                             
            ? <div>Loading applications</div>
            : <table className='applications-table'>
              <thead>
                <tr>
                  <th>Application Name</th>
                  <th>Status</th>
                  <th>Day created</th>
                  <th />
                </tr>
              </thead>
              <tbody>
                {
                  (!isFetching && applications.length === 0)
                    ? <tr>
                      <td colSpan={3}>There is no applications</td>
                    </tr>
                    : applications.map((application) => {
                      const status = get(application, 'status', '')
                      return <tr key={application.id}>
                        <td style={{ width: '34%%' }}>
                          {application.name || '-'}
                        </td>
                        <td style={{ width: '15%' }}>
                          <Status status={application.status} />
                        </td>
                        <td style={{ width: '15%' }}>
                          <Status status={application.status} />
                        </td>
                        <td style={{ width: '33%' }}>
                          <div style={{ float: 'right' }}>
                            {['finished', 'running', 'cancel'].includes(status) &&
                              <button
className='btn btn-primary'
                                onClick={() => history.push(`/application/${application.id}`)}
                              >
                                View Details
                              </button>}
                            {['running', 'pending'].includes(status) &&
                              <button className='btn btn-cancel'>
                                 onClick={() => app.cancelRunApplication(application.id)} 
                                >
                                Cancel
                              </button>}
                          </div>
                        </td>
                             </tr>
                    })
                }
              </tbody>
              </table>
        }
      </div>
    </div>
       </div>
};
export default Applications
