import APIS from './url.js'
import request from '@/utils/request'
export function getMenus (params) {
    return request({
        url: APIS.menuApi.menu,
        method: 'get',
        params: params
    })
}

export function addMenu (data) {
    return request({
      url: APIS.menuApi.menu,
      method: 'post',
      data: data
    })
  }

  export function updateMenu (data) {
    return request({
      url: APIS.menuApi.menu,
      method: 'put',
      data: data
    })
  }

  export function deleteMenu (data) {
    return request({
      url: APIS.menuApi.menu,
      method: 'delete',
      data: data
    })
  }
