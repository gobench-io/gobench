import React, { useEffect, useState, lazy, Suspense } from 'react';
import { isArray, get } from 'lodash';
import './App.css';
import GoBenchAPI from '../../api/gobench';
import { AppContext } from '../../context';

const GroupComponent = lazy(() => import('./Group'));

const loading = () => <p>Loading group...</p>;

const statusColor = {
  running: '#4dbd74',
  init: '#ffcc00',
  finished: '#0066ff',
  cancel: '#ff0000'
};

const App = () => {
  const [groups, fetchGroups] = useState([]);
  const [appData, fetchAppData] = useState(null);
  const [fetching, setFetching] = useState(false);

  useEffect(() => {
    setFetching(true);
    GoBenchAPI.getAppInfo().then(res => {
      setFetching(false);
      fetchAppData(res);
    })
  }, []);
  useEffect(() => {
    if (appData) {
      GoBenchAPI.getGroups().then(res => {
        fetchGroups(res);
      })
    }
  }, [appData]);
  const appStatus = get(appData, 'status', '');
  if (!appData && !fetching) {
    return <div className="app">
      <p>Cannot load the application. May be your benchmark has been stopped.</p>
    </div>
  }

  return !fetching && (
    <div className="app">
      <div className="app-header">
        <h2>{get(appData, 'name', '') || ''}</h2>
        <span
          style={{
            color: '#FFFFFF',
            background: statusColor[appStatus] || '#bfbfbf',
          }}>{get(appData, 'status', '')}</span>
      </div>
      <div className="container">
        <AppContext.Provider value={appData}>
          {
            (isArray(groups) && groups.length > 0) &&
            groups.map((group, index) => {
              return <Suspense key={group.id || index} fallback={loading()}>
                <GroupComponent group={group} appData={appData}/>
              </Suspense>
            })
          }
        </AppContext.Provider>
      </div>
    </div>
  );
}

export default App;
