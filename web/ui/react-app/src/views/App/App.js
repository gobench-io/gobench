import React, { useEffect, useState } from 'react';
import { get } from 'lodash';
import GoBenchAPI from '../../api/gobench';
import { AppContext } from '../../context';
import { useParams, useHistory } from 'react-router-dom';

import { statusColors } from '../../components/Status';
import Dashboard from './Dashboard';
import Scenario from './Scenario';

const App = () => {
  const [appData, fetchAppData] = useState(null);
  const [fetching, setFetching] = useState(false);
  const [activeTab, setActiveTab] = useState('dashboard');
  const { appId } = useParams();
  const history = useHistory();

  useEffect(() => {
    setFetching(true);
    GoBenchAPI.getAppInfo(appId).then(res => {
      setFetching(false);
      fetchAppData(res);
    })
  }, []);
  const appStatus = get(appData, 'status', '');
  if (!appData && !fetching) {
    return <div className="app">
      <p>Loading application benchmark</p>
    </div>
  }

  return !fetching && (
    <div className="card">
      <div className="app-header">
        <div className="app-header-left">
          <h2 className="application-title">
            {get(appData, 'name', '') || ''} application benchmark
          </h2>
          <span
            className="application-status"
            style={{
              color: '#FFFFFF',
              background: statusColors[appStatus] || '#bfbfbf',
            }}>{get(appData, 'status', '')}</span>
        </div>
        <button className="btn btn-cancel" onClick={() => history.goBack()}>&lt; Back to Applications</button>
      </div>
      <div className="">
        <ul className="tabs">
          <li
            className={`tab-nav-item ${activeTab === 'dashboard' ? 'tab-active' : ''}`}
            onClick={() => setActiveTab('dashboard')}>
            Dashboard
            </li>
          <li
            className={`tab-nav-item ${activeTab === 'scenario' ? 'tab-active' : ''}`}
            onClick={() => setActiveTab('scenario')}>
            Scenario
             </li>
        </ul>
        <AppContext.Provider value={appData}>
          <div className="tab-content">
            {
              activeTab === 'dashboard' &&
              <div className="tab-item">
                <Dashboard />
              </div>
            }
            {
              activeTab === 'scenario' &&
              <div className="tab-item">
                <Scenario />
              </div>
            }
          </div>
        </AppContext.Provider>
      </div>
    </div>
  );
}

export default App;
