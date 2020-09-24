import React, { useState, useEffect } from 'react'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'
import { injectIntl } from 'react-intl'
import { Dropdown, Input, Tooltip, message } from 'antd'
import PerfectScrollbar from 'react-perfect-scrollbar'
import store from 'store'
import style from './style.module.scss'

const mapStateToProps = ({ menu }) => ({
  menuData: menu.menuData,
})

const FavPages = ({ menuData = [], intl: { formatMessage } }) => {
  const [searchText, setSearchText] = useState('')
  const [favs, setFavs] = useState(store.get('app.topbar.favs') || [])
  const [pagesList, setPagesList] = useState([])

  useEffect(() => {
    const getPagesList = () => {
      const menuDataProcessed = JSON.parse(JSON.stringify(menuData))
      const flattenItems = (items, key) =>
        items.reduce((flattenedItems, item) => {
          if (item.category) {
            return flattenedItems
          }
          if (item.key === 'nestedItem1' || item.disabled) {
            // skip unwanted items
            return flattenedItems
          }
          if (Array.isArray(item[key])) {
            const itemsProcessed = item[key].map(child => {
              child.icon = item.icon
              return child
            })
            return flattenedItems.concat(flattenItems(itemsProcessed, key))
          }
          flattenedItems.push(item)
          return flattenedItems
        }, [])
      return flattenItems(menuDataProcessed, 'children')
    }
    setPagesList(getPagesList())
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [menuData])

  const changeSearchText = e => {
    setSearchText(e.target.value)
  }

  const setFav = (e, item) => {
    e.preventDefault()
    e.stopPropagation()
    const isActive = favs.some(child => child.url === item.url)
    if (isActive) {
      const filtered = favs.filter(child => child.url !== item.url)
      store.set('app.topbar.favs', filtered)
      setFavs(filtered)
      return
    }
    if (favs.length >= 3) {
      message.info('Only three pages can be added to your bookmarks.')
      return
    }
    const items = [...favs]
    items.push(item)
    store.set('app.topbar.favs', items)
    setFavs(items)
  }

  const generatePageList = () => {
    const searchTextProcessed = searchText ? searchText.toUpperCase() : ''
    return pagesList.map(item => {
      const isActive = favs.some(child => child.url === item.url)
      if (!item.title.toUpperCase().includes(searchTextProcessed) && searchTextProcessed) {
        return null
      }
      return (
        <Link to={item.url} className={style.link} key={item.key}>
          <div
            className={`${style.setIcon} ${isActive ? style.setIconActive : ''}`}
            onClick={e => setFav(e, item)}
            role="button"
            tabIndex="0"
            onFocus={e => {
              e.preventDefault()
            }}
            onKeyPress={e => setFav(e, item)}
          >
            <i className="fe fe-star" />
          </div>
          <span>
            <i className={`mr-2 fe ${item.icon}`} />
            {item.title}
          </span>
        </Link>
      )
    })
  }

  const menu = (
    <React.Fragment>
      <div className="card cui__utils__shadow width-300">
        <div className="card-body p-1 ">
          <div className="p-2">
            <Input
              placeholder={formatMessage({ id: 'topBar.findPages' })}
              value={searchText}
              onChange={changeSearchText}
              allowClear
            />
          </div>
          <div className="height-200">
            <PerfectScrollbar>
              <div className="px-2 pb-2">{generatePageList(searchText)}</div>
            </PerfectScrollbar>
          </div>
        </div>
      </div>
    </React.Fragment>
  )
  return (
    <div className={style.container}>
      {favs.map(item => {
        return (
          <Tooltip key={item.key} placement="bottom" title={item.title}>
            <Link to={item.url} className={`${style.item} mr-3`}>
              <i className={`${style.icon} fe ${item.icon}`} />
            </Link>
          </Tooltip>
        )
      })}
      <Tooltip placement="bottom" title="Bookmarks">
        <Dropdown overlay={menu} trigger={['click']} placement="bottomLeft">
          <span className={style.item}>
            <i className={`${style.icon} fe fe-star`} />
          </span>
        </Dropdown>
      </Tooltip>
    </div>
  )
}

export default injectIntl(connect(mapStateToProps)(FavPages))
