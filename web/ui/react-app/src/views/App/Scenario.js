import React, { useContext } from 'react'
import { get } from 'lodash'
import { Divider } from 'antd'

import { AppContext } from '../../context'

const Scenario = () => {
  const appData = useContext(AppContext)
  return <div>
    <textarea
      disabled
      className='application-scenario'
      value={get(appData, 'scenario', '')}
    />
    <Divider orientation='left' plain>
     gomod
    </Divider>
    <textarea
      disabled
      className='application-gomodule application-scenario'
      value={get(appData, 'gomod', '')}
    />
    <Divider orientation='left' plain>
     gosum
    </Divider>
    <textarea
      disabled
      className='aapplication-gomodule application-scenario'
      value={get(appData, 'gosum', '')}
    />
  </div>
}
export default Scenario
