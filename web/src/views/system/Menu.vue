<!-- eslint-disable vue/no-unused-vars -->
<template>
  <a-card :bordered="false">
    <a-form :form="queryParam" :label-col="{ span: 5 }" :wrapper-col="{ span: 17 }">
      <a-row :gutter="48">

        <a-col :md="8" :sm="24">
          <a-form-item label="搜索">
            <a-input
              v-model.trim="queryParam.query"
              @pressEnter="searchMenu"
              placeholder="请输入关键信息"
            />
          </a-form-item>
        </a-col>
        <a-col :md="8" :sm="24">
          <a-button type="primary" @click="searchMenu">查询</a-button>
          <a-button @click="openClose" style="margin:0px 15px">{{ treeOpen?"折叠":"展开" }}菜单</a-button>
          <a-button
            v-if="$auth(menuApi + '.post')"
            type="primary"
            icon="plus"
            @click="addMenu"
          >添加菜单</a-button>
        </a-col>
      </a-row>
    </a-form>
    <a-table
      rowKey="id"
      :columns="columns"
      :data-source="data"
      :expanded-row-keys.sync="expandedRowIds"
      :loading="loadingTable"
      size="small"
      :scroll="{ x: 1600 }"
      :bordered="false"
    >
      <span slot="icon" slot-scope="icon,record">
        <a-icon v-if="icon" :type="icon" /> {{ icon }}
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
      <span slot="show" slot-scope="show,record">

        <a-badge
          v-if="show===true&&record.type==1"
          status="success"
          text="是"
        />
        <a-badge
          v-else-if="show===false&&record.type==1"
          status="error"
          text="否"
        />
      </span>

      <span slot="sort" slot-scope="sort,record">
        <span v-if="record.type===1">{{ sort }}</span>
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
</template>

<script>
import APIS from '@/api/url'
import { getMenus } from '@/api/menu'

const columns = [
  {
    title: '标题',
    dataIndex: 'title',
    fixed: 'left',
    width: 180
  },
  {
    title: '权限标识',
    dataIndex: 'name',
    scopedSlots: { customRender: 'name' },
    width: 230
  },
  {
    title: '路径',
    dataIndex: 'path',
    width: 190
  },
  {
    title: '图标',
    dataIndex: 'icon',
    scopedSlots: { customRender: 'icon' },
    width: 150
  },
  {
    title: '组件',
    dataIndex: 'component',
    width: 180
  },
  {
    title: '重定向地址',
    dataIndex: 'redirect',
    width: 190
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
    title: '排序',
    dataIndex: 'sort',
    width: 70,
    scopedSlots: { customRender: 'sort' }
  },
  {
    title: '类型',
    dataIndex: 'type',
    fixed: 'right',
    scopedSlots: { customRender: 'type' }
  },
  {
    title: '操作',
    key: 'action',
    fixed: 'right',
    scopedSlots: { customRender: 'action' }
  }
]
export default {
  name: 'Menu',
  data () {
    return {
        loadingTable: false,
        menuApi: APIS.BaseUrl + APIS.menuApi.menu,
        menuTree: [],
        columns,
        data: [],
        expandedRowIds: [],
        queryParam: { query: '' },
        treeOpen: true
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
                this.menuTree = data.menu_tree
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
    },
    async searchMenu (expandAll) {
        let result = []
       await this.queryMenus()
        if (this.queryParam.query) {
            result = this.findMenu(this.menuTree)
            this.data = result
            // 展开所有
            if (expandAll !== false) {
              this.expandedRow(this.data)
            }
        } else {
            this.data = this.menuTree
        }
    },
    expandedRow (dataList) {
      dataList.forEach(element => {
          this.expandedRowIds.push(element.id)
          if (element.children && element.children.length > 0) {
            this.expandedRow(element.children)
          }
      })
    },
    // 递归查找菜单，如果子菜单符合条件同时把父级菜单给展示出来，如果只有父级菜单符合条件则不展示器子菜单
    findMenu (dataList) {
      if (!dataList || dataList.length === 0) {
        return null
      }
        const result = []
        dataList.forEach(element => {
            if ((element.name && element.name.indexOf(this.queryParam.query) !== -1) ||
                (element.path && element.path.indexOf(this.queryParam.query) !== -1) ||
             (element.component && element.component.indexOf(this.queryParam.query) !== -1) ||
             (element.redirect && element.redirect.indexOf(this.queryParam.query) !== -1) ||
             (element.title && element.title.indexOf(this.queryParam.query) !== -1)) {
                result.push(element)
                if (element.children && element.children.length > 0) {
                  const children = this.findMenu(element.children)
                  // console.log(element.name, '---->', children)
                  if (children && children.length > 0) {
                    element.children = children
                  } else { // 如果这个菜单的没有符合条件的子菜单则清空这个菜单的子菜单
                    element.children = null
                  }
                }
            } else if (element.children && element.children.length > 0) { // 直接找儿子，如果儿子中符合条件的菜单则把当前这个菜单添加进来
              const children = this.findMenu(element.children)
              if (children && children.length > 0) {
                element.children = children
                result.push(element)
                console.log('element', element)
              }
            }
        })
        return result
    },
    addMenu () {

    },
    openClose () {
      if (this.treeOpen) {
        this.expandedRowIds = []
      } else {
        this.expandedRow(this.data)
      }
      this.treeOpen = !this.treeOpen
    }
  },
  mounted () {
    this.searchMenu(false)
  }

}
</script>

<style>

</style>
