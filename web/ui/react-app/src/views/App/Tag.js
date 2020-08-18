import React, { useState, useEffect, useContext } from 'react'
import { Tag, Input, Tooltip } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import { colorFull } from '../../components/Status'
import { AppContext } from '../../context'

const Tags = (props) => {
  const app = useContext(AppContext)
  const [editable, setEditable] = useState(false)
  const [newTag, setNewTag] = useState('')
  const [tags, setTags] = useState([])

  useEffect(() => {
    if (app.tags) {
      const _tags = app.tags.split(',')
      setTags(_tags)
    }
  }, [app])

  const removeTag = removedTag => {
    const _tags = tags.filter(x => x !== removedTag)

    setTags(_tags)
    props.saveTags(_tags)
  }
  const addTag = (saved) => {
    if (newTag && tags.indexOf(newTag) === -1) {
      const _tags = [...tags, newTag]
      setTags(_tags)
      if (saved) {
        props.saveTags(_tags)
      }
    } else {
      setEditable(false)
    }

    setNewTag('')
  }
  return (
    <>
      {tags && tags.map((tag, index) => {
        const isLongTag = tag.length > 20

        const tagElem = (
          <Tag
            className='edit-tag'
            key={tag}
            closable
            onClose={() => removeTag(tag)}
            color={colorFull()}
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
export default Tags
