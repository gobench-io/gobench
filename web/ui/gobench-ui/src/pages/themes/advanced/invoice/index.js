import React from 'react'
import { Helmet } from 'react-helmet'
import Table8 from 'components/kit/widgets/Tables/8'

const SystemInvoice = () => {
  return (
    <div>
      <Helmet title="Advanced / Invoice" />
      <div className="kit__utils__heading">
        <h5>Invoice</h5>
      </div>
      <Table8 />
    </div>
  )
}

export default SystemInvoice
