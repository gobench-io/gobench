import { combineReducers } from 'redux'
import { connectRouter } from 'connected-react-router'
import user from './user/reducers'
import menu from './menu/reducers'
import settings from './settings/reducers'
import application from './application/reducers'

export default history =>
  combineReducers({
    router: connectRouter(history),
    user,
    menu,
    settings,
    application,
    system: (state = {}, action) => {
      switch (action.type) {
        case 'system/API_CALL_FAILURE':
          window.location.href = '#/auth/login'
          return state
        case 'system/API_ERROR':
          console.log('TODO: handle error')
          return state
        default:
          return state
      }
    }
  })
