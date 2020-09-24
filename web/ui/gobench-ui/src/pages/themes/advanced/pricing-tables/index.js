import React from 'react'
import { Helmet } from 'react-helmet'
import List21 from 'components/kit/widgets/Lists/21'
import List21v1 from 'components/kit/widgets/Lists/21v1'
import List21v2 from 'components/kit/widgets/Lists/21v2'

const SystemPricingTables = () => {
  return (
    <div>
      <Helmet title="Advanced / Pricing Tables" />
      <div className="kit__utils__heading">
        <h5>Pricing Tables</h5>
      </div>
      <div className="row">
        <div className="col-lg-4">
          <List21 />
        </div>
        <div className="col-lg-4">
          <List21v1 />
        </div>
        <div className="col-lg-4">
          <List21v2 />
        </div>
      </div>
    </div>
  )
}

export default SystemPricingTables
