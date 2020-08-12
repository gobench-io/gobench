import { notification } from 'antd'
import React, { useEffect, useState } from 'react'
import { ErrorContext } from '../../context'

const Notification = (props) => {
  const setError = (errorMessage) => {
    setEm({
      ...em,
      ...errorMessage
    })
  }
  const [em, setEm] = useState({ setError })
  useEffect(() => {
    if (em.type) {
      notification[em.type]({ ...em })
      // reset error message
      setEm({ setEm })
    }
  }, [em, setError])
  return (
    <ErrorContext.Provider value={em}>
      {props.children}
    </ErrorContext.Provider>
  )
}
export default Notification
