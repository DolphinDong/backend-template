import Vue from 'vue'
import moment from 'moment'
import 'moment/locale/zh-cn'
moment.locale('zh-cn')

Vue.filter('NumberFormat', function (value) {
  if (!value) {
    return '0'
  }
  const intPartFormat = value.toString().replace(/(\d)(?=(?:\d{3})+$)/g, '$1,') // 将整数部分逢三一断
  return intPartFormat
})

Vue.filter('dayjs', function (dataStr, pattern = 'YYYY-MM-DD HH:mm:ss') {
  return moment(dataStr).format(pattern)
})

Vue.filter('moment', function (dataStr, pattern = 'YYYY-MM-DD HH:mm:ss') {
  return moment(dataStr).format(pattern)
})

Vue.filter('timeFomaterFilter', function (timestamp) {
  timestamp = parseInt(timestamp)
  var date = new Date(timestamp)
  const Y = date.getFullYear()
  const M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1)
  const D = (date.getDate() < 10 ? '0' + (date.getDate()) : date.getDate())
  const h = (date.getHours() < 10 ? '0' + (date.getHours()) : date.getHours())
  const m = (date.getMinutes() < 10 ? '0' + (date.getMinutes()) : date.getMinutes())
  const s = (date.getSeconds() < 10 ? '0' + (date.getSeconds()) : date.getSeconds())
return Y + '-' + M + '-' + D + ' ' + h + ':' + m + ':' + s
})

Vue.filter('timeFomaterFilter2', function (timestamp) {
  timestamp = parseInt(timestamp)
  var date = new Date(timestamp * 1000)
  const Y = date.getFullYear()
  const M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1)
  const D = (date.getDate() < 10 ? '0' + (date.getDate()) : date.getDate())
  const h = (date.getHours() < 10 ? '0' + (date.getHours()) : date.getHours())
  const m = (date.getMinutes() < 10 ? '0' + (date.getMinutes()) : date.getMinutes())
  const s = (date.getSeconds() < 10 ? '0' + (date.getSeconds()) : date.getSeconds())
return Y + '-' + M + '-' + D + ' ' + h + ':' + m + ':' + s
})

Vue.filter('statusTypeFilter', function (status) {
  return status === 1 ? 'success' : 'error'
})
