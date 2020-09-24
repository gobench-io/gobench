import actions from './actions'

const initialState = {
  id: '',
  name: '',
  role: '',
  email: '',
  avatar: '',
  authorized: process.env.REACT_APP_AUTHENTICATED || false, // false is default value
  loading: false
}

export default function userReducer (state = initialState, action) {
  switch (action.type) {
    case actions.SET_STATE:
      return { ...state, ...action.payload }
    default:
      return state
  }
}
