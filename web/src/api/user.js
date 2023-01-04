import APIS from './url.js'
import request from '@/utils/request'

export function getUsers (params) {
    return request({
      url: APIS.userApi.user,
      method: 'get',
      params: params
    })
  }

export function addUser (data) {
    return request({
      url: APIS.userApi.user,
      method: 'post',
      data: data
    })
}
