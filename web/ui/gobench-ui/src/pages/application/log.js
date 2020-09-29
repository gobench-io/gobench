import React, { useEffect, useState } from 'react'
import { Helmet } from 'react-helmet'
import { connect } from 'react-redux'
import { withRouter } from 'react-router-dom'
import 'prismjs/components/prism-clike'
import 'prismjs/components/prism-go'
import { INTERVAL } from 'constant'
import 'css/editor.css'
import { Input } from 'antd'

const { TextArea } = Input
const mapStateToProps = ({ application, dispatch }) => ({ detail: application.detail, logs: application.logs, dispatch })

const DefaultPage = ({ detail, logs, dispatch }) => {
  const [fetching, setFetching] = useState(false)
  const { id } = detail
  useEffect(() => {
    if (!fetching) {
      dispatch({
        type: 'application/LOG',
        payload: { id }
      })
      setFetching(true)
    }
  }, [id])
  useEffect(() => {
    const interval = setInterval(() => {
      dispatch({
        type: 'application/LOG',
        payload: { id }
      })
    }, INTERVAL / 2)
    // destroy interval on unmount
    return () => {
      console.log('unmout log')
      clearInterval(interval)
    }
  })
  return (
    <>
      <div className='application-log'>
        <Helmet title='Application| log' />
        <TextArea readOnly value={logs} rows={50} />
      </div>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))
