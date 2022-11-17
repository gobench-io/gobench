import axios from 'axios'
import store from 'store'
import { notification } from 'antd'
import { history } from 'index'

const API_ENDPOINT = process.env.NODE_ENV === 'production' ? '/api' : 'http://localhost:8080/api'
const apiClient = axios.create({
  baseURL: API_ENDPOINT,
  timeout: 5 * 60 * 1000,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json'
  }
})

apiClient.interceptors.request.use(request => {
  const accessToken = store.get('accessToken')
  if (accessToken) {
    request.headers.Authorization = `Bearer ${accessToken}`
    // request.headers.AccessToken = accessToken
  }
  return request
})

apiClient.interceptors.response.use(undefined, error => {
  // Errors handling
  const { response } = error
  if (!response) {
    notification.warning({
      message: 'request failed!'
    })
    return
  }
  let { data, statusText, status } = response
  console.log('res', response)
  if (status === 401) {
    history.push('/auth/login')
  }
  if (data) {
    if (typeof data === 'object') {
      data = statusText
    }
    notification.warning({
      message: data
    })
  }
})

export default apiClient
