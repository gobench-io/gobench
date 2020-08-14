import React, { useContext, useEffect, useState } from 'react'
import { Menu } from 'antd'
import { PieChartOutlined, AreaChartOutlined } from '@ant-design/icons'
import { Link } from 'react-router-dom'
import { RootContext } from '../../context'
import Status, { statusColors, iconStatus } from '../Status'
const { Item } = Menu

const MenuLeft = (props) => {
  const app = useContext(RootContext)
  const [menus, setMenus] = useState([])
  useEffect(() => {
    if (app.apps) {
      const menus = app.apps.map(x => ({
        label: x.name,
        status: x.status,
        color: statusColors[x.status],
        icon: iconStatus(x.status),
        link: `/application/${x.id}`
      }))
      setMenus(menus)
    }
  }, [app.apps])
  return (
    <Menu theme={props.theme} defaultSelectedKeys={props.defaultSelected} mode={props.mode}>
      <Item key='1' icon={<PieChartOutlined />}>
        <Link to='/'>  Applications</Link>
      </Item>
      {menuItem(menus)}
    </Menu>
  )
}
const menuItem = (items) => {
  return items.length > 0
    ? items.map((item, index) => (
      <Item
        key={`sub-${index}`}
      >
        <Link className='menu-label' to={item.link}>
          <AreaChartOutlined style={{ color: item.color }} />
          <span className="menu-label-title">{item.label}</span>
          <div style={{ marginLeft: 'auto' }}>
            <Status status={item.status} shortcut />
          </div>
        </Link>
      </Item>
    )) : null
}
export default MenuLeft
