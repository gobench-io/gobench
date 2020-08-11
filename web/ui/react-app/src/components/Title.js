import React from 'react'
import { Link } from 'react-router-dom'

const Title = () => {
  return (
    <div className='gobench-title'>
      <Link className='app-title' to='/'>
        <h1> Gobench</h1>
      </Link>
    </div>)
}

export default Title
