import React, { useEffect, useState, lazy, Suspense, useContext } from 'react'
import { isArray, get } from 'lodash'

import GoBenchAPI from '../../api/gobench'
import { AppContext } from '../../context'

const GroupComponent = lazy(() => import('./Group'))

const loading = () => <p>Loading group...</p>

const Dashboard = () => {
  const appData = useContext(AppContext)
  const [groups, fetchGroups] = useState([])

  useEffect(() => {
    if (appData) {
      GoBenchAPI.getGroups(get(appData, 'id', '')).then(res => {
        fetchGroups(res)
      })
    }
  }, [appData])
  return <div>
    {
      (isArray(groups) && groups.length > 0) &&
      groups.map((group, index) => {
        return <Suspense key={group.id || index} fallback={loading()}>
          <GroupComponent group={group} appData={appData} expandDefault={index === 0} />
        </Suspense>
      })
    }
  </div>
}

export default Dashboard
