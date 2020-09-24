import apiClient from 'services/axios'
import store from 'store'

export async function login (username, password) {
  return apiClient
    .post('/users/login', {
      username,
      password
    })
    .then(response => {
      if (response) {
        const { id } = response.data
        if (id) {
          store.set('accessToken', id)
        }
        return response.data
      }
      return false
    })
    .catch(err => console.log(err))
}

export async function register (email, password, name) {
  return apiClient
    .post('/auth/register', {
      email,
      password,
      name
    })
    .then(response => {
      if (response) {
        const { accessToken } = response.data
        if (accessToken) {
          store.set('accessToken', accessToken)
        }
        return response.data
      }
      return false
    })
    .catch(err => console.log(err))
}

export async function currentAccount () {
  const accessToken = store.get('accessToken')
  if (!accessToken) {
    return false
  }
  return accessToken
}

export async function logout () {

  // .get('/auth/logout')
  // .then(() => {
  //   store.remove('accessToken')
  //   return true
  // })
  // .catch(err => console.log(err))
}
