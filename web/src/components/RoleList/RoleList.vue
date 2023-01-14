<template>
  <div>
    <a-form :form="queryParam" :label-col="{ span: 5 }" :wrapper-col="{ span: 17 }">
      <a-row :gutter="48">
        <a-col :md="12" :sm="24">
          <a-form-item label="搜索">
            <a-input v-model.trim="queryParam.query" @pressEnter="searchRole" placeholder="请输入关键信息" />
          </a-form-item>
        </a-col>
        <a-col :md="8" :sm="24">
          <a-button type="primary" @click="searchRole">查询</a-button>
        </a-col>
      </a-row>
    </a-form>
    <a-table
      rowKey="id"
      :columns="columns"
      :pagination="pagination"
      :loading="loadingTable"
      :data-source="data"
      size="small"
      :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }">

    </a-table>
  </div>
</template>

<script>
import { getUserRole } from '@/api/user'
import { getRoles } from '@/api/role'
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
  }
]
export default {
    name: 'RoleList',
    data () {
        return {
            queryParam: { query: '' },
            loadingTable: false,
            data: [],
            columns: columns,
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
            selectedRowKeys: []
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
        onSelectChange (selectedRowKeys) {
      // console.log('selectedRowKeys changed: ', selectedRowKeys)
      this.selectedRowKeys = selectedRowKeys
    }
    },
    mounted () {
        this.queryRole()
    }

}
</script>

<style></style>
