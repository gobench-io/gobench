import React, { useState } from 'react'
import { Table } from 'antd'
import data from './data.json'
import style from './style.module.scss'

const columns = [
  {
    title: 'PRODUCT',
    dataIndex: 'product',
    className: 'text-gray-6',
    key: 'product',
  },
  {
    title: 'LOCATION',
    dataIndex: 'location',
    key: 'location',
    className: 'text-gray-6',
    render: text => {
      return (
        <a href="" className="text-blue">
          {text}
        </a>
      )
    },
  },
  {
    title: 'DESCRIPTION',
    dataIndex: 'description',
    className: 'text-gray-6',
    key: 'description',
  },
  {
    title: 'QUANTITY',
    dataIndex: 'quantity',
    key: 'quantity',
    className: 'text-right text-gray-6',
    render: text => {
      return <div className="font-weight-bold">{text}</div>
    },
  },
  {
    title: 'UNIT COST',
    dataIndex: 'cost',
    key: 'cost',
    className: 'text-right text-gray-6',
    render: text => {
      return <div className="font-weight-bold">{text}</div>
    },
  },
  {
    title: 'SUMMARY',
    dataIndex: 'overall',
    key: 'overall',
    className: 'text-right text-gray-6',
    render: text => {
      return <div className="font-weight-bold">{text}</div>
    },
  },
]

const Table8 = () => {
  const [selectedRowKeys, setSelectedRowKeys] = useState(['1', '2', '3'])

  const onSelectChange = keys => {
    setSelectedRowKeys(keys)
  }

  const rowSelection = {
    selectedRowKeys,
    onChange: onSelectChange,
  }

  return (
    <div>
      <div className="d-flex flex-wrap">
        <div className="flex-grow-1 d-flex flex-column flex-sm-row mb-4">
          <div className="font-size-18 font-weight-bold text-uppercase mb-4">
            <div>From:</div>
            <div className="text-dark mb-3">Amazon delivery</div>
            <img
              className="d-block"
              src="resources/images/content/amazon-logo.jpg"
              alt="Amazon logo"
            />
          </div>
          <div className="ml-sm-auto mr-lg-auto pr-3">
            795 Folsom Ave, Suite 600
            <br />
            San Francisco, CA, 94107
            <br />
            E-mail: example@amazon.com
            <br />
            Phone: (123) 456-7890
            <br />
            Fax: 800-692-7753
          </div>
        </div>
        <div className="flex-grow-1 d-flex flex-column flex-sm-row mb-4">
          <div className="font-size-18 font-weight-bold text-uppercase pb-4">
            <div>To:</div>
            <div className="text-dark mb-3">Invoice info</div>
            <div className="text-dark">W32567-2352-4756</div>
            <div className="text-dark">Artour Arteezy</div>
          </div>
          <div className="mt-auto mt-sm-0 ml-sm-auto pr-3 mr-lg-auto">
            795 Folsom Ave, Suite 600
            <br />
            San Francisco, CA, 94107
            <br />
            P: (123) 456-7890
            <br />
            Invoice Date: January 20, 2016
            <br />
            Due Date: January 22, 2016
          </div>
        </div>
      </div>
      <div className="mb-4">
        <Table
          className={style.table}
          columns={columns}
          dataSource={data}
          pagination={false}
          rowSelection={rowSelection}
        />
      </div>
      <div className="text-right font-size-18 text-dark p-4 rounded bg-light">
        <div>
          Sub - Total amount: <span className="font-weight-bold">$406,472.50</span>
        </div>
        <div>
          VAT: <span className="font-weight-bold">$81,294.50</span>
        </div>
        <div>
          Grand Total: <span className="font-weight-bold">$487,767.00</span>
        </div>
        <a href="" className="btn btn-outline-success mt-3 mr-3">
          Print
        </a>
        <a href="" className="btn btn-success mt-3">
          Proceed Payment
        </a>
      </div>
    </div>
  )
}

export default Table8
