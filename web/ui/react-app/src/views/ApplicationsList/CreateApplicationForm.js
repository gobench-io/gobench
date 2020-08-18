import React, { useState, useContext } from 'react'
import Editor from 'react-simple-code-editor'
import { highlight, languages } from 'prismjs/components/prism-core'
import 'prismjs/components/prism-clike'
import 'prismjs/components/prism-go'
import { useHistory } from 'react-router-dom'

import '../../css/editor.css'
import { RootContext } from '../../context'

const CreateApplicationForm = () => {
  const history = useHistory()
  const app = useContext(RootContext)
  const [name, onChangeName] = useState('')
  const [code, onChangeEditor] = useState('')

  return <div className='create-application-form'>
    <div className='card'>
      <div className='create-form-header'>
        <h2>Create new application</h2>
        <div className='form-submit'>
          <button
            className='btn btn-primary'
            type='button'
            onClick={() => {
              app.submitCreate({ name, scenario: code })
            }}
          >
            Create Application
          </button>
          <button
            className='btn btn-cancel'
            type='button'
            onClick={() => history.goBack()}
          >
            Cancel
          </button>
        </div>
      </div>
      <div className='form-group'>
        <label className='form-label' htmlFor='name'>Application name:</label>
        <input
          type='text'
          name='name'
          className='form-input'
          placeholder='Enter application name'
          onChange={e => onChangeName(e.target.value)}
        />
      </div>
      <div className='form-group'>
        <label className='form-label' htmlFor='scenario'>Scenario:</label>
        <div className='editor-container'>
          <Editor
            value={code}
            onValueChange={c => onChangeEditor(c)}
            highlight={code => highlight(code, languages.go, 'go')}
            padding={16}
            tabSize={4}
            insertSpaces
            className='editor'
            autoFocus
            style={{
              fontFamily: '"Arial", "Open Sans", monospace',
              fontSize: 14
            }}
          />
        </div>
      </div>
    </div>
  </div>
}

export default CreateApplicationForm
