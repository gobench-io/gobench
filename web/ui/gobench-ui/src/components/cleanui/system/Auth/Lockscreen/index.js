import React from 'react'
import { Input, Button, Form } from 'antd'
import { Link } from 'react-router-dom'
import style from '../style.module.scss'

const Lockscreen = () => {
  const onFinish = values => {
    console.log('Success:', values)
  }

  const onFinishFailed = errorInfo => {
    console.log('Failed:', errorInfo)
  }

  return (
    <div>
      <div className={`card ${style.container}`}>
        <div className="text-dark text-center font-size-24 mb-4">
          <strong>Account Locked</strong>
        </div>
        <div className="text-center">
          <div className="kit__utils__avatar kit__utils__avatar--size64 d-inline-block mb-2">
            <img src="resources/images/avatars/2.jpg" alt="Mary Stanform" />
          </div>
          <div className="font-size-18 text-dark mb-4">
            <strong>Mary Stanform</strong>
          </div>
        </div>
        <Form
          layout="vertical"
          hideRequiredMark
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          className="mb-4"
        >
          <Form.Item
            name="password"
            rules={[{ required: true, message: 'Please input your e-mail address' }]}
          >
            <Input size="large" placeholder="Password" />
          </Form.Item>
          <Button type="primary" htmlType="submit" size="large" className="text-center w-100">
            <strong>Unlock Account</strong>
          </Button>
        </Form>
      </div>
      <div className="text-center pt-2 mb-auto">
        <span className="mr-2">Already have an account?</span>
        <Link to="/auth/login" className="kit__utils__link font-size-16">
          Sign in
        </Link>
      </div>
    </div>
  )
}

export default Lockscreen
