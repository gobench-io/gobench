import React, { useContext } from 'react'
import { get } from 'lodash'

import { AppContext } from '../../context'

const Logs = () => {
  const appData = useContext(AppContext)
  return <div>
    <textarea
      disabled
      className='application-scenario'
      value='TODO'
    />
         </div>
}
export default Logs
