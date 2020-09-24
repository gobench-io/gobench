import React from 'react'

const General7 = () => {
  return (
    <div>
      <div className="d-flex flex-wrap align-items-center mb-2">
        <div className="flex-shrink-0 kit__utils__avatar mr-4 mb-2">
          <img src="resources/images/avatars/5.jpg" alt="Mary Stanform" />
        </div>
        <div className="mb-2">
          <div className="text-dark font-size-18 font-weight-bold text-nowrap">
            Helen Maggie
            <i className="align-text-bottom fe fe-check-square text-success ml-2 font-size-24 " />
          </div>
          <div className="text-uppercase">Support team</div>
        </div>
      </div>
      <div className="mb-3">
        <a className="btn btn-outline-primary mr-2">Chat</a>
        <a className="btn btn-outline-danger">Unfollow</a>
      </div>
      <div className="table-responsive">
        <table className="table table-borderless">
          <tbody>
            <tr>
              <td className="text-gray-6 pl-0">Location</td>
              <td className="pr-0 text-right text-dark">New York</td>
            </tr>
            <tr>
              <td className="text-gray-6 pl-0">Phone</td>
              <td className="pr-0 text-right text-dark">+1 800 367 4784</td>
            </tr>
            <tr>
              <td className="text-gray-6 pl-0">Email</td>
              <td className="pr-0 text-right text-dark">mail@google.com</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  )
}

export default General7
