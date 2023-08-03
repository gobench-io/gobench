import React, { useState, useEffect } from 'react'
import { connect } from 'react-redux'
import { Link, withRouter } from 'react-router-dom'
import { Menu, Layout } from 'antd'
import classNames from 'classnames'
import store from 'store'
import PerfectScrollbar from 'react-perfect-scrollbar'
import { find } from 'lodash'
import style from './style.module.scss'

const mapStateToProps = ({ menu, settings, user }) => ({
  menuData: menu.menuData,
  isMenuCollapsed: settings.isMenuCollapsed,
  isMobileView: settings.isMobileView,
  isMenuUnfixed: settings.isMenuUnfixed,
  isMenuShadow: settings.isMenuShadow,
  leftMenuWidth: settings.leftMenuWidth,
  menuColor: settings.menuColor,
  logo: settings.logo,
  role: user.role,
})

const MenuLeft = ({
  dispatch,
  menuData = [],
  location: { pathname },

  isMenuCollapsed,
  isMobileView,
  isMenuUnfixed,
  isMenuShadow,
  leftMenuWidth,
  menuColor,
  logo,
  role,
}) => {
  const [selectedKeys, setSelectedKeys] = useState(store.get('app.menu.selectedKeys') || [])
  const [openedKeys, setOpenedKeys] = useState(store.get('app.menu.openedKeys') || [])

  useEffect(() => {
    applySelectedKeys()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [pathname, menuData])

  const applySelectedKeys = () => {
    const flattenItems = (items, key) =>
      items.reduce((flattenedItems, item) => {
        flattenedItems.push(item)
        if (Array.isArray(item[key])) {
          return flattenedItems.concat(flattenItems(item[key], key))
        }
        return flattenedItems
      }, [])
    const selectedItem = find(flattenItems(menuData, 'children'), ['url', pathname])
    setSelectedKeys(selectedItem ? [selectedItem.key] : [])
  }

  const onCollapse = (value, type) => {
    if (type === 'responsive' && isMenuCollapsed) {
      return
    }
    dispatch({
      type: 'settings/CHANGE_SETTING',
      payload: {
        setting: 'isMenuCollapsed',
        value: !isMenuCollapsed,
      },
    })
    setOpenedKeys([])
  }

  const onOpenChange = keys => {
    store.set('app.menu.openedKeys', keys)
    setOpenedKeys(keys)
  }

  const handleClick = e => {
    store.set('app.menu.selectedKeys', [e.key])
    setSelectedKeys([e.key])
  }

  const generateMenuItems = () => {
    const generateItem = item => {
      const { key, title, url, icon, disabled, count } = item
      if (item.category) {
        return <Menu.ItemGroup key={Math.random()} title={item.title} />
      }
      if (item.url) {
        return (
          <Menu.Item key={key} disabled={disabled}>
            {item.target && (
              <a href={url} target={item.target} rel="noopener noreferrer">
                <span className={style.title}>{title}</span>
                {count && <span className="badge badge-success ms-2">{count}</span>}
                {icon && <span className={`${icon} ${style.icon} icon-collapsed-hidden`} />}
              </a>
            )}
            {!item.target && (
              <Link to={url}>
                <span className={style.title}>{title}</span>
                {count && <span className="badge badge-success ms-2">{count}</span>}
                {icon && <span className={`${icon} ${style.icon} icon-collapsed-hidden`} />}
              </Link>
            )}
          </Menu.Item>
        )
      }
      return (
        <Menu.Item key={key} disabled={disabled}>
          <span className={style.title}>{title}</span>
          {count && <span className="badge badge-success ms-2">{count}</span>}
          {icon && <span className={`${icon} ${style.icon} icon-collapsed-hidden`} />}
        </Menu.Item>
      )
    }

    const generateSubmenu = items =>
      items.map(menuItem => {
        if (menuItem.children) {
          const subMenuTitle = (
            <span key={menuItem.key}>
              <span className={style.title}>{menuItem.title}</span>
              {menuItem.count && <span className="badge badge-success ms-2">{menuItem.count}</span>}
              {menuItem.icon && <span className={`${menuItem.icon} ${style.icon}`} />}
            </span>
          )
          return (
            <Menu.SubMenu title={subMenuTitle} key={menuItem.key}>
              {generateSubmenu(menuItem.children)}
            </Menu.SubMenu>
          )
        }
        return generateItem(menuItem)
      })

    return menuData.map(menuItem => {
      if (menuItem.roles && !menuItem.roles.includes(role)) {
        return null
      }
      if (menuItem.children) {
        const subMenuTitle = (
          <span key={menuItem.key}>
            <span className={style.title}>{menuItem.title}</span>
            {menuItem.count && <span className="badge badge-success ms-2">{menuItem.count}</span>}
            {menuItem.icon && <span className={`${menuItem.icon} ${style.icon}`} />}
          </span>
        )
        return (
          <Menu.SubMenu title={subMenuTitle} key={menuItem.key}>
            {generateSubmenu(menuItem.children)}
          </Menu.SubMenu>
        )
      }
      return generateItem(menuItem)
    })
  }

  const menuSettings = isMobileView
    ? {
      width: leftMenuWidth,
      collapsible: false,
      collapsed: false,
      onCollapse,
    }
    : {
      width: leftMenuWidth,
      collapsible: true,
      collapsed: isMenuCollapsed,
      onCollapse,
      breakpoint: 'lg',
    }

  return (
    <Layout.Sider
      {...menuSettings}
      className={classNames(`${style.menu}`, {
        [style.white]: menuColor === 'white',
        [style.gray]: menuColor === 'gray',
        [style.dark]: menuColor === 'dark',
        [style.unfixed]: isMenuUnfixed,
        [style.shadow]: isMenuShadow,
      })}
    >
      <div
        className={style.menuOuter}
        style={{
          width: isMenuCollapsed && !isMobileView ? 80 : leftMenuWidth,
          height: isMobileView || isMenuUnfixed ? 'calc(100% - 64px)' : 'calc(100% - 110px)',
        }}
      >
        <div className={style.logoContainer}>
          <div className={style.logo}>
            <img src="resources/images/logo.svg" className="me-2" alt="gobench" />
            <div className={style.name}>{logo}</div>
          </div>
        </div>
        <PerfectScrollbar>
          <Menu
            onClick={handleClick}
            selectedKeys={selectedKeys}
            openKeys={openedKeys}
            onOpenChange={onOpenChange}
            mode="inline"
            className={style.navigation}
            inlineIndent="15"
          >
            {generateMenuItems()}
          </Menu>
          <div className={style.banner}>
            <br />
            <br />
            <br />
            <br />
            <br />
            <br />
            <br />
            <br />
            <br />
            <br />
          </div>
        </PerfectScrollbar>
      </div>
    </Layout.Sider>
  )
}

export default withRouter(connect(mapStateToProps)(MenuLeft))
