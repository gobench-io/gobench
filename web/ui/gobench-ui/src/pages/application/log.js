import React, { } from 'react'
import { Helmet } from 'react-helmet'
import { connect } from 'react-redux'
import { withRouter } from 'react-router-dom'
import 'prismjs/components/prism-clike'
import 'prismjs/components/prism-go'
import 'css/editor.css'

const mapStateToProps = ({ application, dispatch }) => ({ detail: application.detail, groups: application.groups, dispatch })

const DefaultPage = ({ detail }) => {
  return (
    <>
      <div className='application-log'>
        <Helmet title='Application| log' />
        <h5>Log</h5>
        Log
      </div>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))
