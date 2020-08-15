import React, { useState, useEffect, useContext } from 'react'
import { Tag, Input, Tooltip } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import { useParams } from 'react-router-dom'
import { AppContext } from '../../context'

const Tags = (props) => {
  const { appId } = useParams()
  const app = useContext(AppContext)
  const [editable, setEditable] = useState(false)
  const [newTag, setNewTag] = useState('')

  useEffect(() => {
    if (app.tags && props.tags.length === 0) {
      props.setTags(app.tags.split(',').filter(x => x))
    }
  }, [appId])

  const removeTag = removedTag => {
    props.setTags(props.tags.filter(x => x !== removedTag))
  }
  const addTag = () => {
    if (newTag && props.tags.indexOf(newTag) === -1) {
      props.setTags([...props.tags, newTag])
    }
    setEditable(false)
    setNewTag('')
  }
  return (
    <>
      {props.tags.map((tag, index) => {
        const isLongTag = tag.length > 20

        const tagElem = (
          <Tag
            className='edit-tag'
            key={tag}
            closable={index !== 0}
            onClose={() => removeTag(tag)}
          >
            <span>
              {isLongTag ? `${tag.slice(0, 20)}...` : tag}
            </span>
          </Tag>
        )
        return isLongTag ? (
          <Tooltip title={tag} key={tag}>
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
          onPressEnter={() => addTag()}
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
export default Tags
