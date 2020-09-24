export const statusColors = {
  running: '#73d13d',
  init: '#ffcc00',
  finished: '#1890ff',
  cancel: '#bfbfbf',
  pending: '#faad14',
  provisioning: '#00bcd4',
  error: '#f5222d',
  info: 'default'
}
export const formatTag = (str) => {
  if (!str || str.length === 0) {
    return []
  }
  const tags = str.split(',').map((text, index) => ({
    color: colorFull(index),
    index,
    text
  }))
  return tags
}
export const colorFull = (index) => {
  const colors = [
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
  if (!index) {
    return colors[0]
  }

  if (index >= colors.length) {
    index = index % colors.length
  }

  return colors[index]
}
