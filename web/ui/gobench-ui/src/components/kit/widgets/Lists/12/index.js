import React from 'react'
import style from './style.module.scss'

const List12 = () => {
  return (
    <div>
      <div className="mb-3">
        <div className="bg-success text-white text-uppercase px-3 py-1 mb-2">
          Today - 7 may 2019
        </div>
        <div className="table-responsive">
          <table className="table table-borderless text-gray-6 mb-0">
            <tbody>
              <tr>
                <td className="text-nowrap">
                  <div className={`${style.donut} ${style.danger} mr-3`} />
                  California
                </td>
                <td className="text-right">
                  <strong>+78,366,263.00$</strong>
                </td>
              </tr>
              <tr>
                <td className="text-nowrap">
                  <div className={`${style.donut} ${style.primary} mr-3`} />
                  Texas
                </td>
                <td className="text-right">
                  <strong>+58,165,000.00$</strong>
                </td>
              </tr>
              <tr>
                <td className="text-nowrap">
                  <div className={`${style.donut} ${style.success} mr-3`} />
                  Wyoming
                </td>
                <td className="text-right">
                  <strong>+26,156,267.00$</strong>
                </td>
              </tr>
              <tr>
                <td className="text-nowrap">
                  <div className={`${style.donut} ${style.info} mr-3`} />
                  Florida
                </td>
                <td className="text-right">
                  <strong>+18,823,026.00$</strong>
                </td>
              </tr>
              <tr>
                <td className="text-nowrap">
                  <div className={`${style.donut} ${style.orange} mr-3`} />
                  New York
                </td>
                <td className="text-right">
                  <strong>+8,125,642.00$</strong>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <div>
        <div className="bg-light text-gray-6 text-uppercase px-3 py-1 mb-2">6 may 2019</div>
        <div className="table-responsive">
          <table className="table table-borderless text-muted mb-0">
            <tbody>
              <tr>
                <td>No Items</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  )
}

export default List12
