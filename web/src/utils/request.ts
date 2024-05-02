
import axios, { type AxiosError, AxiosResponse } from 'axios'

interface ApiReqeuseOptions {
  url: string
  params?: Record<string, string>
  data?: Record<string, string>
}

const instance = axios.create({
  baseURL: '/api',
  timeout: 10 * 1000
})

instance.interceptors.request.use(
  (config) => {
    return config
  },
  (err: AxiosError) => {
    console.log(err)
  }
)

instance.interceptors.response.use(
  (response: AxiosResponse) => {
    return response.data
  },
  (err: AxiosError) => {
    console.log(err)
  }
)

export const apiGet = <T, D = any>({ url, params }: ApiReqeuseOptions) => {
  return instance.get<D, T>(url, {
    params,
  })
}

export const apiPost = <T, D = any>({ url, data, params }: ApiReqeuseOptions) => {
  return instance.post<D, T>(url, data, {
    params
  })
}

export default instance
