import React, { useState, useEffect } from 'react'
import { Form, Button, Input } from 'antd'
import { connect } from 'react-redux'
import Editor from 'react-simple-code-editor'
import { highlight, languages } from 'prismjs/components/prism-core'
import 'prismjs/components/prism-clike'
import 'prismjs/components/prism-go'
import { withRouter, useHistory } from 'react-router-dom'
import 'css/editor.css'

const { Item } = Form

const mapStateToProps = ({ application, dispatch }) => {
  const { loading, clone } = application
  return {
    loading,
    clone,
    dispatch
  }
}
const DefaultPage = ({ loading, clone, dispatch }) => {
  const [form] = Form.useForm()
  const history = useHistory()
  const [name, setName] = useState('')
  const [scenario, setSceraio] = useState('')
  const [gomod, setGomod] = useState('')
  const [gosum, setGosum] = useState('')

  useEffect(() => {
    // init data if clone
    if (!form.getFieldValue('name') && clone) {
      form.setFieldsValue({
        ...clone
      })
      setName(clone.name)
      setSceraio(clone.scenario)
      setGomod(clone.gomod)
      setGosum(clone.gosum)
    }
  })
  const onFinish = values => {
    if (!values.gomod) {
      values.gomod = ''
    }
    if (!values.gosum) {
      values.gosum = ''
    }
    dispatch({
      type: 'application/CREATE',
      payload: values
    })
  }
  const onChange = (field, value) => {
    form.setFieldsValue({ [field]: value })
    switch (field) {
      case 'name':
        setName(value)
        break
      case 'scenario':
        setSceraio(value)
        break
      case 'gomod':
        setGomod(value)
        break
      case 'gosum':
        setGosum(value)
        break
      default:
        break
    }
  }

  const onFinishFailed = errorInfo => {
    console.log('Failed:', errorInfo)
  }
  return (
    <>
      <div className='container'>
        <div className='card'>
          <Form
            layout='vertical'
            hideRequiredMark
            onFinish={onFinish}
            onFinishFailed={onFinishFailed}
            className='mb-4'
            form={form}
          >
            <div className='card-header row'>
              <div className='col-md-6'>
                <div className='cui__utils__heading mb-0'>
                  <h3>Create new Application</h3>
                </div>
                <div className='text-muted'>Tips: You can clone another application and run again!</div>
              </div>
              <div className='col-md-6'>
                <div className='text-end'>
                  <Button
                    type='primary'
                    size='large'
                    htmlType='submit'
                    loading={loading}
                  >
                    <strong>Create</strong>
                  </Button>

                  <Button
                    size='large'
                    style={{ marginLeft: 5 }}
                    onClick={() => history.push('/applications')}
                  >
                    <strong>Cancel</strong>
                  </Button>
                </div>
              </div>
            </div>
            <div className='card-body'>
              <Item
                name='name'
                label={<h4 className='mb-2'>
                  <strong>Application Name</strong>
                </h4>}
                rules={[{ required: true, message: 'Application name is required' }]}
              >
                <Input value={name} onChange={e => onChange('name', e.target.value)} placeholder='Please input the application name.' />
              </Item>
              <Item
                name='scenario'
                label={
                  <h4 className='mb-2'>
                    <strong>Scenario</strong>
                  </h4>}
                rules={[{ required: true, message: 'Scenario is required' }]}
              >
                <Editor
                  value={scenario}
                  onValueChange={c => onChange('scenario', c)}
                  highlight={code => highlight((code || ''), languages.go, 'go')}
                  padding={16}
                  tabSize={4}
                  insertSpaces
                  className='editor editor-container'
                  autoFocus
                  style={{
                    fontFamily: '"Arial", "Open Sans", monospace',
                    fontSize: 14
                  }}
                />
              </Item>
              <Item
                name='gomod'
                label={<h4 className='mb-2'>
                  <strong>gomod</strong>
                </h4>}
                rules={[{ required: false }]}
              >
                <Editor
                  value={gomod}
                  onValueChange={c => onChange('gomod', c)}
                  highlight={code => highlight((code || ''), languages.go, 'go')}
                  padding={16}
                  tabSize={4}
                  insertSpaces
                  className='editor editor-container'
                  style={{
                    fontFamily: '"Arial", "Open Sans", monospace',
                    fontSize: 14
                  }}
                />
              </Item>
              <Item
                name='gosum'
                label={<h4 className='mb-2'>
                  <strong>gosum</strong>
                </h4>}
                rules={[{ required: false }]}
              >
                <Editor
                  value={gosum}
                  onValueChange={c => onChange('gosum', c)}
                  highlight={code => highlight((code || ''), languages.go, 'go')}
                  padding={16}
                  tabSize={4}
                  insertSpaces
                  className='editor editor-container'
                  style={{
                    fontFamily: '"Arial", "Open Sans", monospace',
                    fontSize: 14
                  }}
                />
              </Item>
            </div>
          </Form>
        </div>
      </div>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))
