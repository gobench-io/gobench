import React, { useState, useEffect } from 'react'
import { Menu, Button } from 'antd'
import { PieChartOutlined } from '@ant-design/icons'
import { Link } from 'react-router-dom'
import { useInterval, INTERVAL } from '../../realtimeHelpers'
import GoBenchAPI from '../../api/gobench'
import { statusColors, iconStatus } from '../Status'

const { Item } = Menu

const MenuLeft = (props) => {
  const [applications, setApplications] = useState([])
  const [menus, setMenus] = useState([])
  const [, setIsFetching] = useState(true)
  console.log('app_data', applications)

  useEffect(() => {
    GoBenchAPI.getApplications().then(apps => {
      const menus = apps.map(x => ({
        label: x.name,
        status: x.status,
        color: statusColors[x.status],
        Icon: iconStatus(x.status),
        link: `/application/${x.id}`
      }))
      setMenus(menus)
      setApplications(apps)
      setIsFetching(false)
    })
  }, [])

  useInterval(() => {
    if (applications && applications.length > 0) {
      GoBenchAPI.getApplications().then(apps => {
        setApplications(apps)
      })
    }
  }, INTERVAL)
  return (
    <Menu theme={props.theme} defaultSelectedKeys={props.defaultSelected} mode={props.mode}>
      <Item key='1' icon={<PieChartOutlined />}>
        Applications
      </Item>
      {menuItem(menus)}
    </Menu>
  )
}
const menuItem = (items) => {
  return items.length > 0
    ? items.map((item, index) => (
      <Item key={`sub-${index}`}>
        <Link className='menu-label' to={item.link}>
          {item.label}
          <Button className='btn-status' style={{ backgroundColor: item.color }}>{item.status}</Button>
        </Link>
      </Item>
    )) : null
}
export default MenuLeft
