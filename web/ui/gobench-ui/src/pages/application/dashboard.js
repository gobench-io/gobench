import React, { useEffect, useState, lazy, Suspense } from 'react'
import { Helmet } from 'react-helmet'
import { connect } from 'react-redux'
import { withRouter, useParams } from 'react-router-dom'
import { isArray } from 'lodash'
import { INTERVAL } from 'constant'

const Group = lazy(() => import('./chart/group'))
const loading = () => <p>Loading group...</p>

const mapStateToProps = ({ application, dispatch }) => ({ detail: application.detail, groups: application.groups, dispatch })

const DefaultPage = ({ detail, groups, dispatch }) => {
  const [fetching, setFetching] = useState(false)
  const { id } = useParams()
  const { status } = detail

  useEffect(() => {
    if (!fetching && id) {
      dispatch({
        type: 'application/GROUPS',
        payload: { id }
      })
      setFetching(true)
    }
  }, [groups, id])
  useEffect(() => {
    if (status === 'running' && groups.length === 0) {
      const interval = setInterval(() => {
        dispatch({
          type: 'application/GROUPS',
          payload: { id }
        })
      }, INTERVAL / 3)
      // destroy interval on unmount
      return () => clearInterval(interval)
    }
  })
  return (
    <>
      <div className='application-dashboard'>
        <Helmet title='Application| Dashboard' />
        {(isArray(groups) && groups.length > 0) &&
      groups.map((group, index) => {
        return (
          <Suspense key={group.id || index} fallback={loading()}>
            <Group group={group} expandDefault={index === 0} />
          </Suspense>
        )
      })}
      </div>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))
