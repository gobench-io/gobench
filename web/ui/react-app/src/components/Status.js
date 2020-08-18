import React from 'react'
import Icon from '@ant-design/icons'


export const statusColors = {
  running: '#52c41a',
  init: '#ffcc00',
  finished: '#1890ff',
  cancel: '#595959',
  pending: '#faad14',
  provisioning: '#00bcd4',
  error: '#f5222d'
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

const Status = ({ status = '', shortcut = false }) => {
  return <span
    className='application-status'
    style={{ background: statusColors[status] || '#bfbfbf' }}
  >
    {shortcut ? status.slice(0, 1).toUpperCase() : status}
  </span>
}

export default Status