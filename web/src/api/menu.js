import APIS from './url.js'
import request from '@/utils/request'
export function getMenus (params) {
    return request({
        url: APIS.menuApi.menu,
        method: 'get',
        params: params
    })
}
