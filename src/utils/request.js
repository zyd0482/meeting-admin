import axios from 'axios'
import { Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'
import qs from 'qs'

axios.defaults.baseURL = process.env.VUE_APP_BASE_API
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'
axios.defaults.withCredentials = true
axios.defaults.timeout = 5000

axios.interceptors.request.use(
  config => {
    if (store.getters.token) {
      config.headers['M-Token'] = getToken()
    }
    return config
  },
  error => {
    console.log(error)
    return Promise.reject(error)
  }
)

export function Post({url, data}) {
  if (data) {
    data = qs.stringify(data)
  }
  
  return axios.post(url, data).then(
    response => {
      const res = response.data
      if (res.code !== 1) {
        Message({
          message: res.msg || 'Error',
          type: 'error',
          duration: 5 * 1000
        })
        return Promise.reject(new Error(res.msg))
      } else {
        return res
      }
    }
  ).catch(
    error => {
      console.log('err' + error) // for debug
      Message({
        message: error.message,
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(error)
    } 
  )
}

