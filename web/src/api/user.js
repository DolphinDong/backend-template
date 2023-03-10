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

export function updateUser (data) {
  return request({
    url: APIS.userApi.user,
    method: 'put',
    data: data
  })
}

export function resetUserPwd (data) {
  return request({
    url: APIS.userApi.resetPwd,
    method: 'put',
    data: data
  })
}

export function deleteUser (data) {
  return request({
    url: APIS.userApi.user,
    method: 'delete',
    data: data
  })
}

export function getUserPermission (params) {
  return request({
    url: APIS.userApi.userPermission,
    method: 'get',
    params: params
  })
}
export function updateUserPermission (data) {
  return request({
    url: APIS.userApi.userPermission,
    method: 'put',
    data: data
  })
}

export function getUserRole (params) {
  return request({
    url: APIS.userApi.userRole,
    method: 'get',
    params: params
  })
}

export function updateUserRole (data) {
  return request({
    url: APIS.userApi.userRole,
    method: 'put',
    data: data
  })
}

export function updateUserAvatar (formData) {
  return request({
    url: APIS.userApi.uploadUserAvatar,
    data: formData,
    method: 'post',
    contentType: false,
    processData: false,
    headers: { 'Content-Type': 'application/x-www-form-urlencoded' } })
}
