import APIS from './url.js'
import request from '@/utils/request'

export function getRoles (params) {
    return request({
        url: APIS.roleApi.role,
        method: 'get',
        params: params
    })
}

export function updateRole (data) {
    return request({
        url: APIS.roleApi.role,
        method: 'put',
        data: data
    })
}

export function addRole (data) {
    return request({
        url: APIS.roleApi.role,
        method: 'post',
        data: data
    })
}

export function deleteRole (data) {
    return request({
        url: APIS.roleApi.role,
        method: 'delete',
        data: data
    })
}

export function getRolePermission (params) {
    return request({
      url: APIS.roleApi.rolePermission,
      method: 'get',
      params: params
    })
  }
  export function updateRolePermission (data) {
    return request({
      url: APIS.roleApi.rolePermission,
      method: 'put',
      data: data
    })
  }
