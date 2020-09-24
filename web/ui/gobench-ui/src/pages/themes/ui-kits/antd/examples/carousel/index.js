/* eslint-disable */
import React from 'react'
import { Carousel } from 'antd'

function onChange(a, b, c) {
  console.log(a, b, c)
}

class AntdCarouselExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <Carousel afterChange={onChange}>
          <div>
            <div
              style={{
                height: '160px',
                textAlign: 'center',
                lineHeight: '160px',
                background: '#364d79',
              }}
            >
              <h3 style={{ lineHeight: '160px', color: '#fff' }}>1</h3>
            </div>
          </div>
          <div>
            <div
              style={{
                height: '160px',
                textAlign: 'center',
                lineHeight: '160px',
                background: '#364d79',
              }}
            >
              <h3 style={{ lineHeight: '160px', color: '#fff' }}>2</h3>
            </div>
          </div>
          <div>
            <div
              style={{
                height: '160px',
                textAlign: 'center',
                lineHeight: '160px',
                background: '#364d79',
              }}
            >
              <h3 style={{ lineHeight: '160px', color: '#fff' }}>3</h3>
            </div>
          </div>
          <div>
            <div
              style={{
                height: '160px',
                textAlign: 'center',
                lineHeight: '160px',
                background: '#364d79',
              }}
            >
              <h3 style={{ lineHeight: '160px', color: '#fff' }}>4</h3>
            </div>
          </div>
        </Carousel>
      </div>
    )
  }
}

export default AntdCarouselExample
