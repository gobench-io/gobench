import React, { useState } from 'react';
import Editor from 'react-simple-code-editor';
import { highlight, languages } from 'prismjs/components/prism-core';
import 'prismjs/components/prism-clike';
import 'prismjs/components/prism-go';
import { useHistory } from 'react-router-dom';

import GoBenchAPI from '../../api/gobench';
import '../../css/editor.css'


const CreateApplicationForm = () => {
  const history = useHistory();
  const [name, onChangeName] = useState('');
  const [code, onChangeEditor] = useState('');
  const [errorMessage, setErrorMessage] = useState(null);

  const submitCreate = ({ name, scenario }) => {
    if (!name || name.trim() === '') {
      return setErrorMessage('name is required');
    }
    if (!scenario || scenario.trim() === '') {
      return setErrorMessage('scenario is required');
    }
    setErrorMessage(null);
    GoBenchAPI.createApplication({
      name,
      scenario: btoa(unescape(encodeURIComponent(scenario)))
    }).then(result => {
      console.log('result', result);
      history.push('/')
    })
  }

  return <div className="container">
    <div className="card">
      <div className="create-form-header">
        <h2>Create new application</h2>
        <div className="form-submit">
          <button
            className="btn btn-primary"
            type="button"
            onClick={() => {
              submitCreate({ name, scenario: code });
            }}>
            Create Application
            </button>
          <button
            className="btn btn-cancel"
            type="button"
            onClick={() => history.goBack()}>
            Cancel
            </button>
        </div>
      </div>
      <div className="form-group">
        <label className="form-label" htmlFor="name">Application name:</label>
        <input
          type="text"
          name="name"
          className="form-input"
          placeholder="Enter application name"
          onChange={e => onChangeName(e.target.value)}
        />
      </div>
      <div className="form-group">
        <label className="form-label" htmlFor="scenario">Scenario:</label>
        <div className="editor-container">
          <Editor
            value={code}
            onValueChange={c => onChangeEditor(c)}
            highlight={code => highlight(code, languages.go, 'go')}
            padding={16}
            tabSize={4}
            insertSpaces={true}
            className="editor"
            autoFocus
            style={{
              fontFamily: '"Arial", "Open Sans", monospace',
              fontSize: 14,
            }}
          />
        </div>
      </div>
      <span className="text-danger">{errorMessage}</span>
    </div>
  </div>
};

export default CreateApplicationForm;