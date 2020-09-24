import React, { } from 'react'
import { Helmet } from 'react-helmet'
import { connect } from 'react-redux'
import { withRouter } from 'react-router-dom'
import { get } from 'lodash'
import { Divider } from 'antd'
import Editor from 'react-simple-code-editor'
import { highlight, languages } from 'prismjs/components/prism-core'
import 'prismjs/components/prism-clike'
import 'prismjs/components/prism-go'
import 'css/editor.css'

const mapStateToProps = ({ application, dispatch }) => ({ detail: application.detail, groups: application.groups, dispatch })

const DefaultPage = ({ detail }) => {
  return (
    <>
      <div className='application-scenario'>
        <Helmet title='Application| scenario' />
        <h5>Scenario</h5>
        <div>
          <div className='application-scenario editor-container'>
            <Editor
              value={get(detail, 'scenario', '')}
              highlight={code => highlight(code, languages.go, 'go')}
              padding={16}
              tabSize={4}
              insertSpaces
              className='editor'
              disabled
              style={{
                fontFamily: '"Arial", "Open Sans", monospace',
                fontSize: 14
              }}
            />
          </div>
          <Divider orientation='left' plain>
     gomod
          </Divider>
          <div className='application-scenario editor-container'>
            <Editor
              value={get(detail, 'gomod', '')}
              highlight={code => highlight(code, languages.go, 'go')}
              padding={16}
              tabSize={4}
              insertSpaces
              className='editor'
              disabled
              style={{
                fontFamily: '"Arial", "Open Sans", monospace',
                fontSize: 14
              }}
            />
          </div>
          <Divider orientation='left' plain>
     gosum
          </Divider>
          <div className='application-scenario editor-container'>
            <Editor
              value={get(detail, 'gosum', '')}
              highlight={code => highlight(code, languages.go, 'go')}
              padding={16}
              tabSize={4}
              insertSpaces
              className='editor'
              disabled
              style={{
                fontFamily: '"Arial", "Open Sans", monospace',
                fontSize: 14
              }}
            />
          </div>
        </div>
      </div>
    </>
  )
}

export default withRouter(connect(mapStateToProps)(DefaultPage))
