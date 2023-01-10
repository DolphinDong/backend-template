<!-- eslint-disable vue/no-unused-vars -->
<template>
  <div>
    <a-card :bordered="false">
      <a-table
        rowKey="id"
        :columns="columns"
        :data-source="data"
        :expanded-row-keys.sync="expandedRowIds"
        :loading="loadingTable"
        size="small"
        :bordered="false"
      >
        <span slot="icon" slot-scope="icon,record">
          <a-icon :type="icon" /> {{ icon }}
        </span>
        <span slot="name" slot-scope="name, record">
          <a-tag color="green">
            {{ name }}
          </a-tag>
        </span>
        <span slot="type" slot-scope="type, record">
          <a-tag v-if="type===1" color="#108ee9">
            菜单
          </a-tag>
          <a-tag v-if="type===2" color="#f50">
            权限
          </a-tag>
        </span>
        <!-- slot-scope="text, record" -->
        <span slot="action" >
          <a v-if="$auth(menuApi + '.put')" >编辑</a>
          <a-divider type="vertical" />
          <a-dropdown v-if="$auth(menuApi + '.delete')">
            <a-menu slot="overlay">
              <a-menu-item
                v-if="$auth(menuApi + '.delete')"
              ><a style="color: red">删除</a></a-menu-item
              >
            </a-menu>
            <a>更多<a-icon type="down" /></a>
          </a-dropdown>
        </span>

      </a-table>
    </a-card>
  </div>
</template>

<script>
import APIS from '@/api/url'
import { getMenus } from '@/api/menu'

const columns = [
  {
    title: '标题',
    dataIndex: 'title'
  },
  {
    title: '权限标识',
    dataIndex: 'name',
    scopedSlots: { customRender: 'name' }
  },
  {
    title: '路径',
    dataIndex: 'path'
  },
  {
    title: '图标',
    dataIndex: 'icon',
    scopedSlots: { customRender: 'icon' }
  },
  {
    title: '组件',
    dataIndex: 'component'
  },
  {
    title: '重定向地址',
    dataIndex: 'redirect'
  },
  {
    title: 'target',
    dataIndex: 'target'
  },
  {
    title: '是否展示',
    dataIndex: 'show',
    scopedSlots: { customRender: 'show' }
  },
  {
    title: '类型',
    dataIndex: 'type',
    scopedSlots: { customRender: 'type' }
  },
  {
    title: '操作',
    key: 'action',
    // fixed: 'right',
    scopedSlots: { customRender: 'action' }
  }
]
export default {
  name: 'Menu',
  data () {
    return {
        loadingTable: false,
        menuApi: APIS.BaseUrl + APIS.menuApi.menu,
        data: [],
        columns,
        expandedRowIds: [1, 2]
    }
  },
  methods: {
    async queryMenus () {
        try {
            this.loadingTable = true
            const res = await getMenus({})
            const data = res.data
            if (data) {
                this.changePermissionId(data.menu_tree)
                this.data = data.menu_tree
                this.expandedRowIds = data.menu_ids
            }
        } catch (e) {

        } finally {
            this.loadingTable = false
        }
    },
    changePermissionId (menuTree) {
      menuTree.forEach(element => {
        if (element.type === 2) {
          element.id = 'p' + element.id
        }
        if (element.children) {
          this.changePermissionId(element.children)
        }
      })
    }
  },
  mounted () {
    this.queryMenus()
  }

}
</script>

<style>

</style>
