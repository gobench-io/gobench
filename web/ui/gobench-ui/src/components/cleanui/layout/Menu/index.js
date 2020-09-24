import React, { useEffect } from 'react'
import { Drawer } from 'antd'
import { connect } from 'react-redux'
import MenuLeft from './MenuLeft'
import MenuTop from './MenuTop'
import style from './style.module.scss'

const mapStateToProps = ({ settings }) => ({
  menuLayoutType: settings.menuLayoutType,
  isMobileMenuOpen: settings.isMobileMenuOpen,
  isMobileView: settings.isMobileView,
  leftMenuWidth: settings.leftMenuWidth,
})

let touchStartPrev = 0
let touchStartLocked = false

const Menu = ({ dispatch, isMobileMenuOpen, isMobileView, menuLayoutType, leftMenuWidth }) => {
  useEffect(() => {
    // mobile menu touch slide opener
    const unify = e => {
      return e.changedTouches ? e.changedTouches[0] : e
    }
    document.addEventListener(
      'touchstart',
      e => {
        const x = unify(e).clientX
        touchStartPrev = x
        touchStartLocked = x > 70
      },
      { passive: false },
    )
    document.addEventListener(
      'touchmove',
      e => {
        const x = unify(e).clientX
        const prev = touchStartPrev
        if (x - prev > 50 && !touchStartLocked) {
          toggleMobileMenu()
          touchStartLocked = true
        }
      },
      { passive: false },
    )
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  const toggleMobileMenu = () => {
    dispatch({
      type: 'settings/CHANGE_SETTING',
      payload: {
        setting: 'isMobileMenuOpen',
        value: !isMobileMenuOpen,
      },
    })
  }

  const GetMenu = () => {
    if (isMobileView) {
      return (
        <div>
          <div
            className={style.handler}
            onClick={toggleMobileMenu}
            onFocus={e => {
              e.preventDefault()
            }}
            onKeyPress={toggleMobileMenu}
            role="button"
            tabIndex="0"
          >
            <div className={style.handlerIcon} />
          </div>
          <Drawer
            closable={false}
            visible={isMobileMenuOpen}
            placement="left"
            className={style.mobileMenu}
            onClose={toggleMobileMenu}
            maskClosable
            getContainer={null}
            width={leftMenuWidth}
          >
            <MenuLeft />
          </Drawer>
        </div>
      )
    }
    if (menuLayoutType === 'top') {
      return <MenuTop />
    }
    if (menuLayoutType === 'nomenu') {
      return null
    }
    return <MenuLeft />
  }

  return GetMenu()
}

export default connect(mapStateToProps)(Menu)
