import React from 'react'
import { connect } from 'react-redux'
import { Input, Button, Radio, Form, Tooltip } from 'antd'
import { Link } from 'react-router-dom'
import style from '../style.module.scss'

const mapStateToProps = ({ user, settings, dispatch }) => ({
  dispatch,
  user,
  authProvider: settings.authProvider,
  logo: settings.logo,
})

const Login = ({ dispatch, user, authProvider, logo }) => {
  const onFinish = values => {
    dispatch({
      type: 'user/LOGIN',
      payload: {...values,username:'admin'},
    })
  }

  const onFinishFailed = errorInfo => {
    console.log('Failed:', errorInfo)
  }


  return (
    <div>
      <div className="text-center mb-5">
        <h1 className="mb-5 px-3">
          <strong>{logo}</strong>
        </h1>
        <p>
          A distributed benchmark tool with Golang
          <br />
          Supporting more than HTTP like MQTT, Websocket, graphQL. 
          <br />
          It can scale to support up to 1 million connection concurrently. 
          <br />
          It could support scriptable tool.
        </p>
      </div>
      <div className={`card ${style.container}`}>
        <div className="text-dark font-size-24 mb-3">
          <strong>Enter your passphrase</strong>
        </div>
        <Form
          layout="vertical"
          hideRequiredMark
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          className="mb-4"
        >
          <Form.Item
            name="username"
            rules={[{ required: false }]}
          >
            <Input type='hidden' value='admin' />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[{ required: true, message: 'Please input password' }]}
          >
            <Input size="large" type="password" placeholder="Password" />
          </Form.Item>
          <Button
            type="primary"
            size="large"
            className="text-center w-100"
            htmlType="submit"
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
