
const loginApi = {
  Login: '/system/login',
  Logout: '/auth/logout',
  ForgePassword: '/auth/forge-password',
  Register: '/auth/register',
  twoStepCode: '/auth/2step-code',
  SendSms: '/account/sms',
  SendSmsErr: '/account/sms_err',
  // get my info
  // UserInfo: '/user/info',
  // UserMenu: '/user/nav'

  UserInfo: '/system/userInfo',
  UserMenu: '/system/menu'
}
const manageApi = {
  user: '/user',
  role: '/role',
  service: '/service',
  permission: '/permission',
  permissionNoPager: '/permission/no-pager',
  orgTree: '/org/tree'
}
const userApi = {
  user: '/system/user',
  resetPwd: '/system/user/resetPwd'
}
const roleApi = {
  role: '/system/role'
}

const APIS = {
  BaseUrl: process.env.VUE_APP_API_BASE_URL,
  loginApi,
  userApi,
  manageApi,
  roleApi
}
export default APIS
