import React from 'react'
import { Helmet } from 'react-helmet'
import Error500 from 'components/cleanui/system/Errors/500'

const System500 = () => {
  return (
    <div>
      <Helmet title='Page 500' />
      <Error500 />
    </div>
  )
}

export default System500
