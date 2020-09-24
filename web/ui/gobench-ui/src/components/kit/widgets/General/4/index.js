import React from 'react'

const General4 = () => {
  return (
    <div>
      <div className="text-dark text-uppercase font-weight-bold mb-1">Work in progress</div>
      <p>Lorem ipsum dolor...</p>
      <div className="progress">
        <div
          className="progress-bar bg-success"
          style={{
            width: '60%',
          }}
          role="progressbar"
          aria-valuenow={60}
          aria-valuemin={0}
          aria-valuemax={100}
        />
      </div>
    </div>
  )
}

export default General4
