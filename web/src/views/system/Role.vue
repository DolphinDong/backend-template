<template>
  <a-card :bordered="false">
    <a-form :form="queryParam" :label-col="{ span: 5 }" :wrapper-col="{ span: 17 }">
      <a-row :gutter="48">

        <a-col :md="8" :sm="24">
          <a-form-item label="搜索">
            <a-input
              v-model="queryParam.query"
              @pressEnter="searchRole"
              placeholder="请输入关键信息"
            />
          </a-form-item>
        </a-col>
        <a-col :md="8" :sm="24">
          <a-button type="primary" @click="searchRole">查询</a-button>
          <a-button
            style="margin-left: 15px"
            v-if="$auth(roleApi + '.post')"
            type="primary"
            icon="usergroup-add"
            @click="addRole"
          >添加角色</a-button>
        </a-col>

      </a-row>
    </a-form>
    <a-table
      rowKey="id"
      :columns="columns"
      :pagination="pagination"
      :loading="loadingTable"
      :data-source="data"
    >
      <span slot="action" slot-scope="text,record">
        <a v-if="$auth(roleApi + '.put')" @click="updateRole(record)">编辑</a>
        <a-divider type="vertical" />
        <a-dropdown v-if="$auth(roleApi + '.delete')">
          <a-menu slot="overlay">
            <a-menu-item
              v-if="$auth(roleApi + '.delete')"
            ><a style="color: red" @click="deleteRole(record)">删除</a></a-menu-item
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
        <a-form-item label="角色名称">
          <a-input
            v-decorator="[
              'role_name',
              {
                rules: [
                  {
                    pattern: /^.{2,30}$/,
                    required: true,
                    message: '请输入正确的角色名称：长度为2-30',
                  },
                ],
              },
            ]"
            placeholder="请输入角色名称"
          />
        </a-form-item>
        <a-form-item label="角色标识">
          <a-input
            :disabled="disableInput"
            v-decorator="[
              'role_identify',
              {
                rules: [
                  {
                    pattern: /^[a-zA-Z0-9_]{4,30}$/,
                    required: true,
                    message: '请输入正确的角色标识：数组字母下划线组成且长度为4-30',
                  },
                ],
              },
            ]"
            placeholder="请输入角色标识"
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </a-card>
</template>

<script>
import APIS from '@/api/url'
import { getRoles, updateRole, addRole, deleteRole } from '@/api/role'
const columns = [
  {
    dataIndex: 'role_name',
    title: '角色名称',
    scopedSlots: { customRender: 'role_name' }
  },
  {
    dataIndex: 'role_identify',
    title: '角色标识',
    scopedSlots: { customRender: 'role_identify' }
  },
  {
    title: '操作',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
    name: 'Role',
    data () {
        return {
            queryParam: { query: '' },
            loadingTable: false,
            roleApi: APIS.BaseUrl + APIS.roleApi.role,
            data: [],
            columns: columns,
            confirmLoading: false,
            ModalText: '',
            disableInput: false,
            visible: false,
            editRole: {},
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
            form: this.$form.createForm(this, { name: 'coordinated' })
        }
    },
    methods: {
      searchRole () {
        this.pagination.current = 1
        this.queryRole()
      },
      handPageChange (page, size) {
      this.pagination.current = page
      this.pagination.pageSize = size
      this.queryRole()
    },
    async queryRole () {
      this.loadingTable = true
      try {
        const res = await getRoles({
        page: this.pagination.current,
        page_size: this.pagination.pageSize,
        search: this.queryParam.query
      })
        this.data = res.data.data
        this.pagination.total = res.data.total
      } catch (e) {
        // this.$message.error('获取角色信息失败')
      } finally {
        this.loadingTable = false
      }
    },
    updateRole (record) {
        this.visible = true
        this.ModalText = '编辑角色'
        this.disableInput = true
        this.editRole = record
        this.form.setFieldsValue({
            role_name: record.role_name,
            role_identify: record.role_identify
        })
    },
    deleteRole (record) {
        this.$confirm({
        title: '是否确认要删除该角色？',
        content: '删除之后将无法恢复',
        okText: '确认',
        okType: 'danger',
        cancelText: '取消',
        onOk: async () => {
          let data = {}
          try {
            data = await deleteRole({ id: record.id })
          } catch (e) {
            return
          }
          if (data.code && data.code === 20001) {
            this.$message.success('删除成功')
            // 如果为本页最后一个则返回到第一页
            if (this.data.length === 1) {
              // this.pagination.current = this.pagination.current > 1 ? this.pagination.current - 1 : 1
              this.pagination.current = 1
            }
            this.queryRole()
          }
        }
      })
    },
    addRole () {
        this.form.resetFields()
        this.editRole = {}
        this.visible = true
        this.ModalText = '添加角色'
        this.disableInput = false
    },
    handleCancel () {
        this.visible = false
        this.form.resetFields()
        this.editRole = {}
    },
    handleOk () {
        this.form.validateFields(async (err, values) => {
        if (err) {
          return
        }
        this.confirmLoading = true
        let data = {}
        try {
          // 编辑
          if (this.editRole.id) {
            values.id = this.editRole.id
            data = await updateRole(values)
          } else {
            // 新增
            data = await addRole(values)
          }
        } catch (e) {
          return
        } finally {
          this.confirmLoading = false
        }

        if (data.code && data.code === 20001) {
          this.$message.success(this.editRole.id ? '编辑成功' : '添加成功')
          this.visible = false
          this.queryRole()
        }
    })
    }
    },
    mounted () {
        this.queryRole()
    }

}
</script>

<style>

</style>
