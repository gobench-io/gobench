import { all, takeEvery, put, call, select } from 'redux-saga/effects'
import { notification } from 'antd'
import { history } from 'index'
import * as jwt from 'services/jwt'
import actions from './actions'

const mapAuthProviders = {
  jwt: {
    login: jwt.login,
    register: jwt.register,
    currentAccount: jwt.currentAccount,
    logout: jwt.logout
  }
}

export function * LOGIN ({ payload }) {
  const { username, password } = payload
  yield put({
    type: 'user/SET_STATE',
    payload: {
      loading: true
    }
  })
  const { authProvider: autProviderName } = yield select(state => state.settings)
  const success = yield call(mapAuthProviders[autProviderName].login, username, password)
  if (success) {
    yield put({
      type: 'user/LOAD_CURRENT_ACCOUNT'
    })
    yield history.push('/')
    // notification.success({
    //   message: 'Logged In',
    //   description: 'You have successfully logged in!'
    // })
  }
  if (!success) {
    yield put({
      type: 'user/SET_STATE',
      payload: {
        loading: false
      }
    })
  }
}

export function * REGISTER ({ payload }) {
  const { email, password, name } = payload
  yield put({
    type: 'user/SET_STATE',
    payload: {
      loading: true
    }
  })
  const { authProvider } = yield select(state => state.settings)
  const success = yield call(mapAuthProviders[authProvider].register, email, password, name)
  if (success) {
    yield put({
      type: 'user/LOAD_CURRENT_ACCOUNT'
    })
    yield history.push('/')
    notification.success({
      message: 'Succesful Registered',
      description: 'You have successfully registered!'
    })
  }
  if (!success) {
    yield put({
      type: 'user/SET_STATE',
      payload: {
        loading: false
      }
    })
  }
}

export function * LOAD_CURRENT_ACCOUNT () {
  yield put({
    type: 'user/SET_STATE',
    payload: {
      loading: true
    }
  })
  const { authProvider } = yield select(state => state.settings)
  const response = yield call(mapAuthProviders[authProvider].currentAccount)
  if (response) {
    const { id, email, name, avatar, role } = response
    yield put({
      type: 'user/SET_STATE',
      payload: {
        id,
        name,
        email,
        avatar,
        role,
        authorized: true
      }
    })
  }
  yield put({
    type: 'user/SET_STATE',
    payload: {
      loading: false
    }
  })
}

export function * LOGOUT () {
  const { authProvider } = yield select(state => state.settings)
  yield call(mapAuthProviders[authProvider].logout)
  yield put({
    type: 'user/SET_STATE',
    payload: {
      id: '',
      name: '',
      role: '',
      email: '',
      avatar: '',
      authorized: false,
      loading: false
    }
  })
}

export default function * rootSaga () {
  yield all([
    takeEvery(actions.LOGIN, LOGIN),
    takeEvery(actions.REGISTER, REGISTER),
    takeEvery(actions.LOAD_CURRENT_ACCOUNT, LOAD_CURRENT_ACCOUNT),
    takeEvery(actions.LOGOUT, LOGOUT),
    LOAD_CURRENT_ACCOUNT() // run once on app load to check user auth
  ])
}
