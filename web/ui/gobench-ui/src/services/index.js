import apiClient from 'services/axios'
let API
export const init = (api) => {
  API = api
}
export const listApi = async (limit = 10, skip = 0, name, tag) => {
  let url = `${API.list}?limit=${limit}&skip=${skip}`
  if (name) {
    url += `&filter[where][name][like]=${name}`
  }
  if (tag) {
    url += `&filter[where][tag][like]=${tag}`
  }
  const response = await apiClient.get(url)
  if (response) {
    return response.data
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
