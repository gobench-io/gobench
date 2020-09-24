import { all, put, call } from 'redux-saga/effects'
import getMenuData from 'services/menu'

export function * GET_DATA () {
  const menuData = yield call(getMenuData)
  yield put({
    type: 'menu/SET_STATE',
    payload: {
      menuData
    }
  })
}

export default function * rootSaga () {
  yield all([
    GET_DATA() // run once on app load to fetch menu data
  ])
}
