import React, { useContext } from 'react'
import { get } from 'lodash'
import { Divider } from 'antd'
import Editor from 'react-simple-code-editor'
import { AppContext } from '../../context'
import { highlight, languages } from 'prismjs/components/prism-core'
import 'prismjs/components/prism-clike'
import 'prismjs/components/prism-go'
import '../../css/editor.css'

const Scenario = () => {
  const appData = useContext(AppContext)
  return <div>
    <div className='application-scenario editor-container'>
      <Editor
        value={get(appData, 'scenario', '')}
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
        value={get(appData, 'gomod', '')}
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
        value={get(appData, 'gosum', '')}
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
}
export default Scenario
