import React from 'react'
import { connect } from 'react-redux'
import { Input, Button, Form } from 'antd'
import style from '../style.module.scss'

const mapStateToProps = ({ user, settings, dispatch }) => ({
  dispatch,
  user,
  authProvider: settings.authProvider,
  logo: settings.logo
})

const Login = ({ dispatch, user, logo }) => {
  const onFinish = values => {
    dispatch({
      type: 'user/LOGIN',
      payload: { ...values, username: 'admin' }
    })
  }

  const onFinishFailed = errorInfo => {
    console.log('Failed:', errorInfo)
  }

  return (
    <div>
      <div className='text-center mb-5'>
        <h1 className='mb-5 px-3'>
          <strong>{logo}</strong>
        </h1>
      </div>
      <div className={`card ${style.container}`}>
        <div className='text-dark font-size-24 mb-3'>
          <strong>Enter your password</strong>
        </div>
        <Form
          layout='vertical'
          hideRequiredMark
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          className='mb-4'
        >
          <Form.Item
            name='username'
            rules={[{ required: false }]}
          >
            <Input type='hidden' value='admin' />
          </Form.Item>
          <Form.Item
            name='password'
            rules={[{ required: true, message: 'Please input password' }]}
          >
            <Input size='large' type='password' placeholder='Password' />
          </Form.Item>
          <Button
            type='primary'
            size='large'
            className='text-center w-100'
            htmlType='submit'
            loading={user.loading}
          >
            <strong>Sign in</strong>
          </Button>
        </Form>
      </div>

    </div>
  )
}

export default connect(mapStateToProps)(Login)
