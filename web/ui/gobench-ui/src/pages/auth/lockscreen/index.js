import React from 'react'
import { Helmet } from 'react-helmet'
import Lockscreen from 'components/cleanui/system/Auth/Lockscreen'

const SystemLockscreen = () => {
  return (
    <div>
      <Helmet title='Lockscreen' />
      <Lockscreen />
    </div>
  )
}

export default SystemLockscreen
