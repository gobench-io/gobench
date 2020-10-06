import React, { useState, useEffect } from 'react'
import { Tag, Input, Tooltip } from 'antd'
import { connect } from 'react-redux'
import { PlusOutlined } from '@ant-design/icons'
import { colorFull } from 'utils/status'
import { withRouter, useParams } from 'react-router-dom'

const mapStateToProps = ({ application, dispatch }) => {
  const { detail, tags } = application
  return { detail, tags, dispatch }
}
const DefaultPage = ({ tags, dispatch }) => {
  const { id } = useParams()

  const [editable, setEditable] = useState(false)
  const [newTag, setNewTag] = useState('')

  useEffect(() => {
    dispatch({
      type: 'application/TAGS',
      payload: { id }
    })
  }, [id])

  const removeTag = tagId => {
    dispatch({
      type: 'application/TAG_REMOVE',
      payload: { id, tagId }
    })
  }
  const addTag = (saved) => {
    if (newTag && tags.indexOf(newTag) === -1) {
      if (saved) {
        dispatch({
          type: 'application/TAG_ADD',
          payload: { id, name: newTag }
        })
      }
    } else {
      setEditable(false)
    }

    setNewTag('')
  }
  return (
    <>
      {tags && tags.map(({ id, name }, index) => {
        const isLongTag = name.length > 20

        const tagElem = (
          <Tag
            className='edit-tag'
            key={id}
            closable
            onClose={() => { removeTag(id) }}
            color={colorFull(index)}
          >
            <span>
              {isLongTag ? `${name.slice(0, 20)}...` : name}
            </span>
          </Tag>
        )
        return isLongTag ? (
          <Tooltip title={name} key={id}>
            {tagElem}
          </Tooltip>
        ) : (
          tagElem
        )
      })}
      {editable && (
        <Input
          autoFocus
          type='text'
          size='small'
          className='tag-input'
          value={newTag}
          onChange={(e) => setNewTag(e.target.value)}
          onBlur={() => addTag()}
          onPressEnter={() => addTag(true)}
        />
      )}
      {!editable && (
        <Tag className='site-tag-plus' onClick={() => setEditable(true)}>
          <PlusOutlined /> New Tag
        </Tag>
      )}
    </>
  )
}
export default withRouter(connect(mapStateToProps)(DefaultPage))
