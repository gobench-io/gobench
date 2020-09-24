import React from 'react'

const General18v1 = () => {
  return (
    <div>
      <div className="py-3 px-4">
        <div className="d-flex flex-wrap-reverse align-items-center pb-3">
          <div className="mr-auto">
            <div className="text-uppercase font-weight-bold font-size-24 text-dark">
              IBAN 4658-1235-1567-8000
            </div>
            <div className="font-size-18">$12,136.78</div>
          </div>
          <div className="flex-shrink-0 font-size-36 text-gray-4 pl-1">
            <i className="fe fe-server" />
          </div>
        </div>
        <div className="font-italic font-size-14 text-center border-top pt-3">
          Current month charged: 12,136.78
        </div>
      </div>
    </div>
  )
}

export default General18v1
