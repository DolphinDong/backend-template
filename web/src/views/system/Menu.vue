<!-- eslint-disable vue/no-unused-vars -->
<template>
  <div>
    <a-card :bordered="false">
      <a-table
        rowKey="id"
        :columns="columns"
        :data-source="data"
        :expanded-row-keys.sync="expandedRowKeys"
      >
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
const data = [
  {
    id: 1,
    title: '菜单管理',
    name: 'menu',
    type: 1,
    children: [
        {
            id: 2,
            title: '删除菜单',
            name: '/api/system/menu : delete',
            type: 2
        }
    ]
  }
]
export default {
    name: 'Menu',
    data () {
    return {
      menuApi: APIS.BaseUrl + APIS.menuApi.menu,
      data,
      columns,
      expandedRowKeys: [1, 2]
    }
  }

}
</script>

<style>

</style>
