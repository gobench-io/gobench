import React, { useEffect, useState } from 'react';
import { get } from 'lodash';
import GoBenchAPI from '../../api/gobench';
import { useHistory } from 'react-router-dom';
import { useInterval } from '../../realtimeHelpers';
import Status from '../../components/Status';

const Applications = () => {
  const history = useHistory();
  const [applications, setApplications] = useState([]);
  const [isFetching, setIsFetching] = useState(true);


  useEffect(() => {
    GoBenchAPI.getApplications().then(apps => {
      setApplications(apps);
      setIsFetching(false);
    })
  }, []);

  useInterval(() => {
    if (applications && applications.length > 0) {
      GoBenchAPI.getApplications().then(apps => {
        setApplications(apps);
      })
    }
  }, 5000)

  const cancelRunApplication = (id) => {
    GoBenchAPI.cancelApplication(id).then(() => {
      GoBenchAPI.getApplications().then(apps => {
        setApplications(apps);
      })
    })
  }
  return <div className="container">
    <div className="card">
      <div className="applications-list-header">
        <h2>Applications</h2>
        <button
          onClick={() => history.push('/application/create')}
          className="btn btn-primary">
          Create application
            </button>
      </div>
      <div className="applications-body">
        {
          isFetching ?
            <div>Loading applications</div>
            : <table className="applications-table">
              <thead>
                <tr>
                  <th>Application Name</th>
                  <th>Status</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                {
                  (!isFetching && applications.length === 0) ?
                    <tr>
                      <td colSpan={3}>There is no applications</td>
                    </tr>
                    : applications.map((app) => {
                      const status = get(app, 'status', '');
                      return <tr key={app.id}>
                        <td style={{ width: '34%%' }}>
                          {app.name || '-'}
                        </td>
                        <td style={{ width: '33%' }}>
                          <Status status={app.status} />
                        </td>
                        <td style={{ width: '33%' }}>
                          <div style={{ float: 'right' }}>
                            {['finished', 'running', 'cancel'].includes(status) &&
                              <button className="btn btn-primary"
                                onClick={() => history.push(`/application/${app.id}`)}>
                                View Details
                        </button>}
                            {['running', 'pending'].includes(status) &&
                              <button className="btn btn-cancel"
                                onClick={() => cancelRunApplication(app.id)}>
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
  </div >;
};
export default Applications;