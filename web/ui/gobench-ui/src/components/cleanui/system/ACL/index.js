import React, { Fragment, useEffect } from 'react'
import { connect } from 'react-redux'
import { notification } from 'antd'
import { history } from 'index'

const mapStateToProps = ({ user }) => ({ user })

const ACL = ({
  user: { role },
  redirect = false,
  defaultRedirect = '/auth/404',
  roles = [],
  children,
}) => {
  useEffect(() => {
    const authorized = roles.includes(role)
    // if user not equal needed role and if component is a page - make redirect to needed route
    if (!authorized && redirect) {
      const url = typeof redirect === 'boolean' ? defaultRedirect : redirect
      notification.error({
        message: 'Unauthorized Access',
        description: (
          <div>
            You have no rights to access this page.
            <br />
            Redirected to &#34;{url}&#34;
          </div>
        ),
      })
      history.push(url)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  const authorized = roles.includes(role)
  const AuthorizedChildren = () => {
    // if user not authorized return null to component
    if (!authorized) {
      return null
    }
    // if access is successful render children
    return <Fragment>{children}</Fragment>
  }

  return AuthorizedChildren()
}

export default connect(mapStateToProps)(ACL)
