import React from 'react'
import Icon from '@ant-design/icons'

export const statusColors = {
  running: '#52c41a',
  init: '#ffcc00',
  finished: '#1890ff',
  cancel: '#595959',
  pending: '#faad14',
  provisioning: '#00bcd4',
  error: '#f5222d',
  info: 'default'
}
export const iconStatus = (status) => {
  switch (status) {
    case 'running':
      return Icon.CheckCircleOutlined
    case 'init':
      return Icon.RedoOutlined
    case 'finished':
      return Icon.AreaChartOutlined
    case 'cancel':
      return Icon.MinusCircleOutlined
    case 'pending':
      return Icon.IssuesCloseOutlined
    case 'provisioning':
      return Icon.InfoCircleOutlined
    case 'error':
      return Icon.CloseCircleOutlined
    default:
      return Icon.WarningOutlined
  }
}
export const colorFull = () => {
  const allColor = [
    '#ff4d4f',
    '#cf1322',
    '#a8071a',
    '#fa541c',
    '#ff7a45',
    '#ad2102',
    '#ffa940',
    '#fa8c16',
    '#ad4e00',
    '#ffc53d',
    '#faad14',
    '#d48806',
    // '#fff566', vang choa mat wa
    '#ffec3d',
    '#fadb14',
    '#d4b106',
    '#d3f261',
    '#bae637',
    '#a0d911',
    '#7cb305',
    '#5b8c00',
    '#73d13d',
    '#52c41a',
    '#389e0d',
    '#237804',
    '#135200',
    '#5cdbd3',
    '#13c2c2',
    '#1890ff',
    '#597ef7',
    '#b37feb',
    '#ff85c0',
    '#722ed1',
    '#eb2f96',
    '#595959'
  ]

  return allColor[Math.floor(Math.random() * allColor.length)]
}
const Status = ({ status = '', shortcut = false }) => {
  return <span
    className='application-status'
    style={{ background: statusColors[status] || '#bfbfbf' }}
         >
    {shortcut ? status.slice(0, 1).toUpperCase() : status}
         </span>
}

export default Status
