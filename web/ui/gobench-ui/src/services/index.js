import apiClient from 'services/axios'
let API
export const init = (api) => {
  API = api
}
export const listApi = async (limit, offset, order, isAsc, keyword) => {
  if (typeof keyword !== 'string') {
    keyword = ''
  }
  const res = await apiClient.get(`${API.count}?keyword=${keyword}`)
  if (!res) {
    return
  }
  const { count } = res.data
  const url = `${API.list}?limit=${limit}&offset=${offset}&order=${order}&isAsc=${isAsc}&keyword=${keyword}`
  const response = await apiClient.get(url)
  if (response) {
    return {
      total: count,
      list: response.data
    }
  }
}
export const detailApi = async (id) => {
  const response = await apiClient.get(API.detail.format(id))
  if (response) {
    return response.data
  }
}
export const createApi = async (data) => {
  const response = await apiClient.post(API.create, data)
  if (response) {
    return response.data
  }
}
export const updateApi = async (id, data) => {
  const response = await apiClient.put(API.update.format(id), data)
  if (response) {
    return response.data
  }
}
export const destroyApi = async (id) => {
  await apiClient.delete(API.delete.format(id))
  return {}
}
