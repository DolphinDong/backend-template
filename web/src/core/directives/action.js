import Vue from 'vue'
import store from '@/store'

/**
 * Action 权限指令
 * 指令用法：
 *  - 在需要控制 action 级别权限的组件上使用 v-action:[method] , 如下：
 *    <i-button v-action:add >添加用户</a-button>
 *    <a-button v-action:delete>删除用户</a-button>
 *    <a v-action:edit @click="edit(record)">修改</a>
 *
 *  - 当前用户没有权限时，组件上使用了该指令则会被隐藏
 *  - 当后台权限跟 pro 提供的模式不同时，只需要针对这里的权限过滤进行修改即可
 *
 *  @see https://github.com/vueComponent/ant-design-vue-pro/pull/53
 */
const action = Vue.directive('action', {
  inserted: function (el, binding, vnode) {
    const actionName = binding.arg
    const roles = store.getters.roles
    const elVal = vnode.context.$route.meta.permission
    // console.log('root permission', elVal)
    // 当前页面所在路由的 meta.permission
    const permissionId = Object.prototype.toString.call(elVal) === '[object String]' && [elVal] || elVal
    let isFind = false
    console.log('permissionId', permissionId)
    // 遍历用户拥有的所有的权限
    roles.permissions.forEach(p => {
      // 先找到对应的权限
      if (!permissionId.includes(p.permissionId)) {
        return
      }
      isFind = true
      // console.log('p', p)
      // 再查看用户该权限中的actions 是否包含传入的action，如果没有则移除当前元素
      if (p.actionList && !p.actionList.includes(actionName)) {
        el.parentNode && el.parentNode.removeChild(el) || (el.style.display = 'none')
      }
    })
    // 如果在用户的信息中没有找到该权限则需要将该元素移除
    if (!isFind) {
      el.parentNode && el.parentNode.removeChild(el) || (el.style.display = 'none')
    }
  }
})

export default action
