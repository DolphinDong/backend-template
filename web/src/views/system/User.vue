<!-- eslint-disable vue/no-unused-vars -->
<template>
  <a-card :bordered="false">
    <a-form :form="queryParam" :label-col="{ span: 5 }" :wrapper-col="{ span: 17 }">
      <a-row :gutter="48">
        <a-col :md="8" :sm="24">
          <a-form-item label="搜索">
            <a-input
              v-model="queryParam.query"
              @pressEnter="searchUser"
              placeholder="请输入关键信息"
            />
          </a-form-item>
        </a-col>
        <template v-if="advanced">
          <a-col :md="8" :sm="24">
            <a-form-item label="性别">
              <a-select v-model="queryParam.gender" placeholder="请选择">
                <a-select-option value="3">全部</a-select-option>
                <a-select-option value="1">男</a-select-option>
                <a-select-option value="2">女</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <!-- <a-col :md="8" :sm="24">
            <a-form-item label="管理员">
              <a-select
                placeholder="请选择"
                v-model="queryParam.is_admin"
                :default-value="2"
              >
                <a-select-option value="2">全部</a-select-option>
                <a-select-option value="0">否</a-select-option>
                <a-select-option value="1">是</a-select-option>
              </a-select>
            </a-form-item>
          </a-col> -->
          <a-col :md="8" :sm="24">
            <a-form-item label="状态">
              <a-select placeholder="请选择" v-model="queryParam.status">
                <a-select-option value="2">全部</a-select-option>
                <a-select-option value="0">禁用</a-select-option>
                <a-select-option value="1">启用</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
        </template>
        <a-col :md="(!advanced && 8) || 24" :sm="24">
          <span
            class="table-page-search-submitButtons"
            :style="(advanced && { float: 'right', overflow: 'hidden' }) || {}"
          >
            <a-button type="primary" @click="searchUser">查询</a-button>
            <a-button
              style="margin-left: 8px"
              @click="
                () => (this.queryParam = { gender: '3', status: '2', is_admin: '2' })
              "
            >重置</a-button
            >
            <a @click="toggleAdvanced" style="margin-left: 8px">
              {{ advanced ? "收起" : "展开" }}
              <a-icon :type="advanced ? 'up' : 'down'" />
            </a>
          </span>
        </a-col>
      </a-row>
    </a-form>
    <a-button
      v-if="$auth(userApi + '.post')"
      style="margin-bottom: 15px"
      type="primary"
      icon="user-add"
      @click="addUser"
    >添加用户</a-button
    >
    <a-table
      rowKey="id"
      :columns="columns"
      :pagination="pagination"
      @change="handleTableChange"
      :loading="loadingTable"
      :data-source="data"
    >
      <span slot="gender" slot-scope="gender">
        <a-badge
          v-if="gender"
          :status="getGenderInfo(gender).type"
          :text="getGenderInfo(gender).text"
        />
      </span>
      <span slot="userStatus" slot-scope="userStatus">
        <a-badge
          v-if="userStatus === false || userStatus === true"
          :status="getStatusInfo(userStatus).type"
          :text="getStatusInfo(userStatus).text"
        />
      </span>
      <!-- <span slot="is_admin" slot-scope="is_admin">
        <a-badge
          v-if="is_admin === false || is_admin === true"
          :status="getIsAdminInfo(is_admin).type"
          :text="getIsAdminInfo(is_admin).text"
        />
      </span> -->
      <span slot="last_login_time" slot-scope="last_login_time">
        {{ last_login_time | timeFomaterFilter2 }}
      </span>

      <span slot="last_login_ip" slot-scope="last_login_ip">
        <a-tag color="green">
          {{ last_login_ip }}
        </a-tag>
      </span>

      <span slot="action" slot-scope="text, record">
        <a v-if="$auth(userApi + '.put')" @click="updateUser(record)">编辑</a>
        <a-divider type="vertical" />
        <a-dropdown v-if="$auth(resetPwdAPi + '.put') || $auth(userApi + '.delete') || $auth(updateUserPermissionApi + '.put')">
          <a-menu slot="overlay">
            <!-- <a-menu-item v-if="$auth(userApi + '.put')"><a>编辑</a></a-menu-item> -->
            <a-menu-item
              v-if="$auth(resetPwdAPi + '.put')"
            ><a @click="resetPwd(record)">重置密码</a></a-menu-item
            >
            <a-menu-item
              v-if="$auth(updateUserPermissionApi + '.put')"
            ><a @click="updateUserPermission(record)">修改权限</a></a-menu-item
            >
            <a-menu-item
              v-if="$auth(userApi + '.delete')"
            ><a style="color: red" @click="deleteUser(record)">删除</a></a-menu-item
            >
          </a-menu>
          <a>更多<a-icon type="down" /></a>
        </a-dropdown>
      </span>
    </a-table>

    <a-modal
      :title="ModalText"
      :visible="visible"
      :confirm-loading="confirmLoading"
      @ok="handleOk"
      @cancel="handleCancel"
      width="40%"
      :forceRender="true"
    >
      <a-form :form="form" :label-col="{ span: 5 }" :wrapper-col="{ span: 17 }">
        <a-form-item label="姓名">
          <a-input
            v-decorator="[
              'username',
              {
                rules: [
                  {
                    pattern: /^.{2,10}$/,
                    required: true,
                    message: '请输入正确的姓名：长度为2-10',
                  },
                ],
              },
            ]"
            placeholder="请输入姓名"
          />
        </a-form-item>
        <a-form-item label="登录名">
          <a-input
            :disabled="disableInput"
            v-decorator="[
              'login_name',
              {
                rules: [
                  {
                    pattern: /^[a-zA-Z0-9]{4,20}$/,
                    required: true,
                    message: '请输入正确的登录名：数组字母组成且长度为4-20',
                  },
                ],
              },
            ]"
            placeholder="请输入登录名"
          />
        </a-form-item>
        <a-form-item label="性别">
          <a-radio-group
            v-decorator="[
              'gender',
              { rules: [{ required: true, message: '请选择性别' }] },
            ]"
          >
            <a-radio :value="1"> 男 </a-radio>
            <a-radio :value="2"> 女 </a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item label="手机号">
          <a-input
            v-decorator="[
              'phone_number',
              {
                rules: [
                  {
                    pattern: /^1[3456789]\d{9}$/,
                    required: true,
                    message: '请输入正确的手机号',
                  },
                ],
              },
            ]"
            placeholder="请输入手机号"
          />
        </a-form-item>
        <a-form-item label="邮箱">
          <a-input
            v-decorator="[
              'email',
              { rules: [{ required: true, type: 'email', message: '请输入正确的邮箱' }] },
            ]"
            placeholder="请输入邮箱"
          />
        </a-form-item>
        <!-- <a-form-item label="管理员">
          <a-radio-group
            v-decorator="[
              'is_admin',
              { rules: [{ required: true, message: '请选择是否为管理员' }] },
            ]"
          >
            <a-radio :value="1"> 是 </a-radio>
            <a-radio :value="0"> 否 </a-radio>
          </a-radio-group>
        </a-form-item> -->
        <a-form-item label="状态">
          <a-radio-group
            v-decorator="[
              'status',
              { rules: [{ required: true, message: '请选择用户状态' }] },
            ]"
          >
            <a-radio :value="1"> 启用 </a-radio>
            <a-radio :value="0"> 禁用 </a-radio>
          </a-radio-group>
        </a-form-item>
      </a-form>
    </a-modal>
    <a-drawer
      :title="drawer.title"
      placement="right"
      :closable="false"
      :visible="drawer.visible"
      :after-visible-change="afterVisibleChange"
      @close="onClose"
      :width="400"
    >
      <a-spin :spinning="drawer.spinning">
        <RoleTree ref="roleTree" ></RoleTree>
      </a-spin>
      <div style="width:100%;height:50px">

      </div>
      <div
        :style="{
          position: 'absolute',
          right: 0,
          bottom: 0,
          width: '100%',
          borderTop: '1px solid #e9e9e9',
          padding: '8px 16px',
          background: '#fff',
          textAlign: 'right',
          zIndex: 1,
        }"
      >
        <a-button :style="{ marginRight: '8px' }" @click="onClose">
          取消
        </a-button>
        <a-button :loading="this.drawer.spinning" type="primary" @click="onSubmit">
          确定
        </a-button>
      </div>
    </a-drawer>
  </a-card>
</template>

<script>
import APIS from '@/api/url'
import {
  getUsers,
  addUser,
  updateUser,
  resetUserPwd,
  deleteUser,
  getUserPermission,
  updateUserPermission } from '@/api/user'

import RoleTree from '@/components/RoleTree'
const genderMap = {
  1: {
    type: 'success',
    text: '男'
  },
  2: {
    type: 'error',
    text: '女'
  }
}

const statusMap = {
  true: {
    type: 'success',
    text: '启用'
  },
  false: {
    type: 'error',
    text: '禁用'
  }
}
const isAdminMap = {
  true: {
    type: 'success',
    text: '是'
  },
  false: {
    type: 'error',
    text: '否'
  }
}

const columns = [
  {
    dataIndex: 'username',
    title: '姓名',
    slots: { title: 'customTitle' },
    scopedSlots: { customRender: 'username' }
  },
  {
    title: '登录名',
    dataIndex: 'login_name'
  },
  {
    title: '性别',
    dataIndex: 'gender',
    scopedSlots: { customRender: 'gender' }
  },
  {
    title: '手机号',
    dataIndex: 'phone_number'
  },
  {
    title: '邮箱',
    dataIndex: 'email'
  },
  // {
  //   title: '管理员',
  //   dataIndex: 'is_admin',
  //   scopedSlots: { customRender: 'is_admin' }
  // },
  {
    title: '状态',
    dataIndex: 'status',
    scopedSlots: { customRender: 'userStatus' }
  },
  {
    title: '上次登录时间',
    dataIndex: 'last_login_time',
    scopedSlots: { customRender: 'last_login_time' }
  },
  {
    title: '上次登录IP',
    dataIndex: 'last_login_ip',
    scopedSlots: { customRender: 'last_login_ip' }
  },
  {
    title: '操作',
    key: 'action',
    // fixed: 'right',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  name: 'User',
  data: function () {
    return {
      ModalText: '新增用户',
      visible: false,
      confirmLoading: false,
      data: [],
      disableInput: false,
      columns,
      userApi: APIS.BaseUrl + APIS.userApi.user,
      resetPwdAPi: APIS.BaseUrl + APIS.userApi.resetPwd,
      updateUserPermissionApi: APIS.BaseUrl + APIS.userApi.userPermission,
      loadingTable: false,
      queryParam: {
        gender: '3',
        status: '2',
        is_admin: '2'
      },
      advanced: false,
      editRecord: {},
      pagination: {
        pageSizeOptions: ['10', '20', '30', '40', '50'],
        current: 1,
        pageSize: 10,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
        onChange: (page, size) => {
          this.handPageChange(page, size)
        },
        onShowSizeChange: (current, size) => {
          this.handPageChange(current, size)
        }
      },
      form: this.$form.createForm(this, { name: 'coordinated' }),
      drawer: {
        visible: false,
        editRecord: {},
        defaultCheck: [],
        spinning: false,
        title: ''
      }
    }
  },
  methods: {
    onShowSizeChange (current, pageSize) {
      this.pagination.pageSize = pageSize
      this.queryUser()
    },
    handleTableChange (pagination, filters, sorter) {
      // console.log(pagination)
      // const pager = { ...this.pagination }
      // pager.current = pagination.current
      // this.pagination = pager
      // this.fetch({
      //   results: pagination.pageSize,
      //   page: pagination.current,
      //   sortField: sorter.field,
      //   sortOrder: sorter.order,
      //   ...filters
      // })
    },
    handPageChange (page, size) {
      this.pagination.current = page
      this.pagination.pageSize = size
      this.queryUser()
    },
    getGenderInfo (gender) {
      return genderMap[gender] || { type: '', text: '' }
    },
    getStatusInfo (status) {
      return statusMap[status] || { type: '', text: '' }
    },
    getIsAdminInfo (isAdmin) {
      return isAdminMap[isAdmin] || { type: '', text: '' }
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    searchUser () {
      this.pagination.current = 1
      this.queryUser()
    },
    handleOk (e) {
      this.form.validateFields(async (err, values) => {
        if (err) {
          return
        }
        values.status = values.status === 1
        // values.is_admin = values.is_admin === 1

        this.confirmLoading = true
        let data = {}
        try {
          // 编辑
          if (this.editRecord.id) {
            values.id = this.editRecord.id
            data = await updateUser(values)
          } else {
            // 新增
            data = await addUser(values)
          }
        } catch (e) {
          return
        } finally {
          this.confirmLoading = false
        }

        if (data.code && data.code === 20001) {
          this.$message.success(this.editRecord.id ? '编辑成功' : '添加成功')
          this.visible = false
          this.queryUser()
        }
        // this.confirmLoading = false
      })
    },
    handleCancel (e) {
      this.form.resetFields()
      this.visible = false
    },
    addUser () {
      this.form.resetFields()
      this.editRecord = {}
      this.disableInput = false
      this.initFormData()
      this.ModalText = '新增用户'
      this.visible = true
    },
    updateUser (record) {
      this.editRecord = record
      this.disableInput = true
      this.ModalText = '编辑用户信息'
      // this.$nextTick(() => {
        this.form.setFieldsValue({
          gender: record.gender,
          //         is_admin: record.is_admin === true ? 1 : 0,
          status: record.status === true ? 1 : 0,
          username: record.username,
          login_name: record.login_name,
          phone_number: record.phone_number,
          email: record.email
        })
     // })

      this.visible = true
    },
    resetPwd (record) {
      this.$confirm({
        title: '是否确认重置密码？',
        content: '重置后密码为: 123456',
        okText: '确认',
        cancelText: '取消',
        onOk: async () => {
          let data = {}
          try {
            data = await resetUserPwd({ id: record.id })
          } catch (e) {
            return
          }
          if (data.code && data.code === 20001) {
            this.$message.success('重置成功')
            // this.queryUser()
          }
        }
      })
    },
    deleteUser (record) {
      this.$confirm({
        title: '是否确认要删除该用户？',
        content: '删除之后将无法恢复',
        okText: '确认',
        okType: 'danger',
        cancelText: '取消',
        onOk: async () => {
          let data = {}
          try {
            data = await deleteUser({ id: record.id })
          } catch (e) {
            return
          }
          if (data.code && data.code === 20001) {
            this.$message.success('删除成功')
            // 如果为本业最后一个则返回到上一页
            if (this.data.length === 1) {
              // this.pagination.current = this.pagination.current > 1 ? this.pagination.current - 1 : 1
              this.pagination.current = 1
            }
            this.queryUser()
          }
        }
      })
    },
    async queryUser () {
      this.loadingTable = true
      try {
      const res = await getUsers({
        page: this.pagination.current,
        page_size: this.pagination.pageSize,
        search: this.queryParam.query,
        gender: this.queryParam.gender,
        //      is_admin: this.queryParam.is_admin,
        status: this.queryParam.status
      })
      this.data = res.data.data
      this.pagination.total = res.data.total
      this.loadingTable = false
    } catch (e) {

      } finally {
        this.loadingTable = false
      }
    },
    initFormData () {
     this.$nextTick(() => {
        this.form.setFieldsValue({
          gender: 1,
          //       is_admin: 0,
          status: 1
        })
     })
    },
    onClose () {
      this.drawer.visible = false
      this.drawer.editRecord = {}
      this.drawer.defaultCheck = []
      this.$refs.roleTree.checkedKeys = []
      this.drawer.spinning = false
      this.drawer.title = ''
    },
    async onSubmit () {
      this.drawer.spinning = true
      let checkedKes = []
      if (this.$refs.roleTree.checkedKeys) {
         checkedKes = this.$refs.roleTree.checkedKeys.checked ? this.$refs.roleTree.checkedKeys.checked : this.$refs.roleTree.checkedKeys
      }
     let data = {}
     try {
       data = await updateUserPermission({ id: this.drawer.editRecord.id, permissions: checkedKes })
     } catch (e) {
      return
     } finally {
     this.drawer.spinning = false
     this.drawer.visible = false
     }
     if (data.code && data.code === 20001) {
          this.$message.success('修改成功')
        }
    },
    afterVisibleChange () {

    },
    async updateUserPermission (record) {
      this.drawer.title = '修改' + record.username + '的权限'
      this.loadingTable = true
      this.drawer.editRecord = record
      this.drawer.defaultCheck = []
      try {
        const data = await getUserPermission({ id: record.id })
        this.drawer.defaultCheck = data.data
        this.drawer.visible = true
        this.drawer.spinning = true
        if (this.$refs.roleTree) {
          this.$refs.roleTree.queryMenus()
        }
        setTimeout(() => {
          this.$refs.roleTree.checkedKeys = this.drawer.defaultCheck
          this.drawer.spinning = false
        }, 1000)
      } catch (e) {
        console.log(e)
        setTimeout(() => {
          this.drawer.visible = false
        }, 2000)
      } finally {
        this.loadingTable = false
      }
    }
  },
  mounted () {
    this.queryUser()
    this.initFormData()
  },
  components: {
    RoleTree

  }
}
</script>

<style scoped>
.numInput {
  width: 100%;
  width: -moz-available;
  width: -webkit-fill-available;
  width: fill-available;
}
</style>
