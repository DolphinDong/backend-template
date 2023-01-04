import axios from 'axios'
import store from '@/store'
import storage from 'store'
import notification from 'ant-design-vue/es/notification'
import { message } from 'ant-design-vue'
import { VueAxios } from './axios'
import { ACCESS_TOKEN } from '@/store/mutation-types'

// 创建 axios 实例
const request = axios.create({
  // API 请求的默认前缀
  baseURL: process.env.VUE_APP_API_BASE_URL,
  timeout: 6000 // 请求超时时间
})

// 异常拦截处理器
const errorHandler = (error) => {
  if (error.response) {
    const data = error.response.data
    // 从 localstorage 获取 token
    const token = storage.get(ACCESS_TOKEN)
    if (error.response.status === 403) {
      notification.error({
        message: '403 Forbidden',
        description: data.msg
      })
      return Promise.reject(error)
    }

    if (error.response.status === 400 || (error.response.status === 500 && data.code === 50001)) {
      message.error(data.msg)
      return Promise.reject(error)
    }

    if (error.response.status === 401 && !(data.result && data.result.isLogin)) {
      notification.error({
        message: '登录提醒',
        description: '会话已过期请重新登录'
      })
      if (token) {
        store.dispatch('Logout').then(() => {
          setTimeout(() => {
            window.location.reload()
          }, 1000)
        })
      }
      return Promise.reject(error)
    }
    if (error.response.data && error.response.data.msg) {
      notification.error({
        message: 'Error',
        description: data.msg
      })
    } else {
      notification.error({
        message: 'Error',
        description: 'Unknown Error'
      })
    }
  } else {
    notification.error({
      message: 'Error',
      description: 'Unknown Error'
    })
  }
  return Promise.reject(error)
}

// request interceptor
request.interceptors.request.use(config => {
  const token = storage.get(ACCESS_TOKEN)
  // 如果 token 存在
  // 让每个请求携带自定义 token 请根据实际情况自行修改
  if (token) {
    config.headers[ACCESS_TOKEN] = token
  }
  return config
}, errorHandler)

// response interceptor
request.interceptors.response.use((response) => {
  if (response && response.headers && response.headers['new-token']) {
    const newToken = response.headers['new-token']
    store.dispatch('ReplaceToken', newToken)
  }

  if (response.data && response.data.msg && response.data.code === 20003) {
    message.warning(response.data.msg)
  }

  return response.data
}, errorHandler)

const installer = {
  vm: {},
  install (Vue) {
    Vue.use(VueAxios, request)
  }
}

export default request

export {
  installer as VueAxios,
  request as axios
}
