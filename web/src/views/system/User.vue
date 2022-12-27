<!-- eslint-disable vue/no-unused-vars -->
<template>
  <div class="div-container">
    <a-button v-if="$auth(userApi+'.post')" icon="user-add" type="primary" style="margin:0 0 10px 0 ;">添加用户</a-button>
    <a-table
      rowKey="id"
      :columns="columns"
      :pagination="pagination"
      @change="handleTableChange"
      :loading="loadingTable"
      :data-source="data">
      <span slot="gender" slot-scope="gender">
        <a-badge v-if="gender" :status="getGenderInfo(gender).type" :text="getGenderInfo(gender).text" />
      </span>

      <span slot="status" slot-scope="status">
        <a-badge v-if="status" :status="getGenderInfo(status).type" :text="getStatusInfo(status).text" />
      </span>
      <span slot="is_admin" slot-scope="is_admin">
        <a-badge v-if="is_admin" :status="getIsAdminInfo(is_admin).type" :text="getIsAdminInfo(is_admin).text" />
      </span>
      <span slot="last_login_time" slot-scope="last_login_time">
        {{ last_login_time | timeFomaterFilter }}
      </span>

      <span slot="last_login_ip" slot-scope="last_login_ip">
        <a-tag color="green" >
          {{ last_login_ip }}
        </a-tag>
      </span>

      <span slot="action" slot-scope="text,record">
        <a-button size="small" v-if="$auth(userApi+'.put')" type="primary" style="margin-right:5px;">编辑</a-button>
        <a-button size="small" v-if="$auth(userApi+'.delete')" type="danger"> 删除</a-button>
      </span>

    </a-table>

  </div>
</template>

<script>
import APIS from '@/api/url'
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
  1: {
    type: 'success',
    text: '启用'
  },
  2: {
    type: 'error',
    text: '禁用'
  }
}
const isAdminMap = {
  1: {
    type: 'success',
    text: '是'
  },
  2: {
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
  {
    title: '管理员',
    dataIndex: 'is_admin',
    scopedSlots: { customRender: 'is_admin' }
  },
  {
    title: '状态',
    dataIndex: 'status',
    scopedSlots: { customRender: 'status' }
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

const data = [
  {
    id: '123456789',
    username: '刘冬',
    gender: 1,
    login_name: 'liudong',
    phone_number: '12345678520',
    email: 'liudong@ones.ai',
    is_admin: 1,
    status: 1,
    last_login_time: 1671874280000,
    last_login_ip: '192.168.66.12'
  },
  {
    id: '1234567891',
    username: '李四',
    gender: 2,
    login_name: 'lisi',
    phone_number: '12345678521',
    email: 'lisi@ones.ai',
    is_admin: 2,
    status: 2,
    last_login_time: '1671875535000',
    last_login_ip: '192.168.66.13'
  }
]
export default {
    name: 'User',
    data: function () {
      return {
          data,
          columns,
          userApi: APIS.BaseUrl + APIS.userApi.user,
          loadingTable: false,
          pagination: {
            pageSizeOptions: ['10', '20', '30', '40', '50'],
            current: 1,
            pageSize: 10,
            total: 1000,
            showSizeChanger: true,
            showTotal: (total) => `共${total}条`,
            onChange: (page, size) => {
              this.handPageChange(page, size)
            },
            onShowSizeChange: (current, size) => {
              this.handPageChange(current, size)
            }
          }

      }
    },
    methods: {
      onShowSizeChange (current, pageSize) {
        this.pageSize = pageSize
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
    },
    getGenderInfo (gender) {
      return genderMap[gender] || { type: '', text: '' }
    },
    getStatusInfo (status) {
     return statusMap[status] || { type: '', text: '' }
    },
    getIsAdminInfo (isAdmin) {
     return isAdminMap[isAdmin] || { type: '', text: '' }
    }
  },
  mounted () {
    console.log(this.userApi)
  }

}
</script>

<style >

</style>
