import React from 'react'
import { connect } from 'react-redux'
import { Layout } from 'antd'
import { withRouter } from 'react-router-dom'
import classNames from 'classnames'
import Footer from 'components/cleanui/layout/Footer'
import style from './style.module.scss'

const mapStateToProps = ({ settings }) => ({
  logo: settings.logo,
  isGrayTopbar: settings.isGrayTopbar,
  isCardShadow: settings.isCardShadow,
  isSquaredBorders: settings.isSquaredBorders,
  isBorderless: settings.isBorderless,
  authPagesColor: settings.authPagesColor
})

const AuthLayout = ({
  children,
  logo,
  isGrayTopbar,
  isCardShadow,
  isSquaredBorders,
  isBorderless,
  authPagesColor
}) => {
  return (
    <Layout>
      <Layout.Content>
        {/* <Sidebar />
        <SupportChat /> */}
        <div
          className={classNames(`${style.container}`, {
            cui__layout__squaredBorders: isSquaredBorders,
            cui__layout__cardsShadow: isCardShadow,
            cui__layout__borderless: isBorderless,
            [style.white]: authPagesColor === 'white',
            [style.gray]: authPagesColor === 'gray'
          })}
          style={{
            backgroundImage:
              authPagesColor === 'image' ? 'url(resources/images/content/photos/7.jpg)' : ''
          }}
        >
          <div
            className={classNames(`${style.topbar}`, {
              [style.topbarGray]: isGrayTopbar
            })}
          >
            <div className={style.logoContainer}>
              <div className={style.logo}>
                <img src='resources/images/logo.svg' className='mr-2' alt='Gobench' />
                <div className={style.name}>{logo}</div>
              </div>
            </div>
          </div>
          <div className={style.containerInner}>{children}</div>
        </div>
      </Layout.Content>
      <Layout.Footer>
        <Footer />
      </Layout.Footer>
    </Layout>
  )
}

export default withRouter(connect(mapStateToProps)(AuthLayout))
