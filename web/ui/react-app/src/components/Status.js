import React from 'react';

const Status = ({ status = '' }) => {
  return <span
    className="application-status"
    style={{ background: statusColors[status] || '#bfbfbf' }}>
    {status}
  </span>
}

export default Status;

export const statusColors = {
  running: '#4dbd74',
  init: '#ffcc00',
  finished: '#0066ff',
  cancel: '#555555',
  pending: '#f9b115',
  provisioning: '#00bcd4',
  error: '#ff0000'
};
