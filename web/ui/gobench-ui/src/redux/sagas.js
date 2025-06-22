import { all, put, takeEvery } from 'redux-saga/effects'
import user from './user/sagas'
import menu from './menu/sagas'
import settings from './settings/sagas'
import application from './application/sagas'

function* handleApiError(error) {
  if (error.response?.status === 401) {
    yield put({ type: 'system/REDIRECT_TO_LOGIN' }); // Trigger redirect
  } else {
    yield put({ type: 'system/API_ERROR', error }); // Log other errors
  }
}

export default function* rootSaga() {
  yield all([user(), menu(), settings(), application(),
  yield takeEvery('system/API_CALL_FAILURE', handleApiError)
  ])
}
