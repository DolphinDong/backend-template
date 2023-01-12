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
      <span slot="action" slot-scope="txt,record">
        <a v-if="$auth(menuApi + '.put')" @click="updateMenu(record)">编辑</a>
        <a-divider type="vertical" />
        <a-dropdown v-if="$auth(menuApi + '.delete')">
          <a-menu slot="overlay">
            <a-menu-item
              v-if="$auth(menuApi + '.delete')"
            ><a style="color: red" @click="deleteMenu(record)">删除</a></a-menu-item
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
        <a-form-item label="菜单名称">
          <a-input
            v-decorator="[
              'title',
              {
                rules: [
                  {
                    required: true,
                    message: '请输入正确的菜单名称',
                  },
                ],
              },
            ]"
            placeholder="请输入菜单名称"
          />
        </a-form-item>
        <a-form-item label="菜单标识">
          <!-- :disabled="editRecord!=null && editRecord.type===2" -->
          <a-input
            v-decorator="[
              'name',
              {
                rules: [
                  {
                    required: true,
                    message: '请输入正确的菜单标识',
                  },
                ],
              },
            ]"
            placeholder="请输入菜单标识"
          />
        </a-form-item>

        <a-form-item label="父级菜单">
          <a-select
            show-search
            placeholder="请选择父级菜单"
            option-filter-prop="children"
            :filter-option="filterOption"
            v-decorator="[
              'parentId',
              { rules: [{ required: true, message: '请选择父级菜单' }] },
            ]"
          >
            <a-select-option :value="0">
              顶级菜单
            </a-select-option>
            <a-select-option v-for="(menu,index) in menus" :value="menu.id" :key="menu.id">
              {{ menu.title }}
            </a-select-option>

          </a-select>
        </a-form-item>
        <a-form-item label="菜单类型">
          <a-select
            show-search
            placeholder="请选择菜单类型"
            v-decorator="[
              'type',
              { rules: [{ required: true, message: '请选择菜单类型' }] },
            ]"
            :disabled="editRecord!=null"
            @change="handleChange"
          >
            <a-select-option :value="1">
              菜单
            </a-select-option>
            <a-select-option :value="2">
              权限
            </a-select-option>
          </a-select>
        </a-form-item>
        <template v-if="isMenu===true">
          <a-form-item label="菜单路径">
            <a-input
              v-decorator="[
                'path',
                {
                  rules: [
                    {
                      pattern: /(^[a-zA-Z0-9_/-]{0,}$)|(^$)/,
                      message: '请输入正确的菜单路径：数字字母_-/组成长度',
                    },
                  ],
                },
              ]"
              placeholder="请输入菜单路径"
            />
          </a-form-item>
          <a-form-item label="菜单图标">
            <a-input
              v-decorator="[
                'icon',
              ]"
              placeholder="请输入菜单图标"
            />
          </a-form-item>
          <a-form-item label="组件">
            <a-input
              v-decorator="[
                'component',
                {
                  rules: [
                    {
                      required: true,
                      message: '请输入正确组件'
                    },
                  ],
                }
              ]"
              placeholder="请输入组件"
            />
          </a-form-item>
          <a-form-item label="重定向地址">
            <a-input
              v-decorator="[
                'redirect',
              ]"
              placeholder="请输入重定向地址"
            />
          </a-form-item>
          <a-form-item label="Target">
            <a-select
              show-search
              placeholder="请选择菜单Target"
              v-decorator="[
                'target',
              ]"
            >
              <a-select-option value="">
                _self
              </a-select-option>
              <a-select-option value="_blank">
                _blank
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item label="是否展示">
            <a-radio-group
              v-decorator="[
                'show',
                {
                  rules: [
                    {
                      required: true,
                      message: '请选择是否展示'
                    },
                  ],
                }
              ]"
            >
              <a-radio :value="1"> 是 </a-radio>
              <a-radio :value="0"> 否 </a-radio>
            </a-radio-group>
          </a-form-item>
          <a-form-item label="排序">
            <a-input-number
              :min="1"
              :max="100"
              v-decorator="[
                'sort',
              ]"
              placeholder="顺序"
            />
          </a-form-item>

        </template>
        <template v-if="isMenu===false">
          <a-form-item label="权限标识">
            <!-- :disabled="editRecord!=null && editRecord.type===2" -->
            <a-select
              show-search
              placeholder="请选择权限标识"
              v-decorator="[
                'action',
                { rules: [{ required: true, message: '请选择权限标识' }] },
              ]"
            >
              <a-select-option value="get">
                GET
              </a-select-option>
              <a-select-option value="post">
                POST
              </a-select-option>
              <a-select-option value="put">
                PUT
              </a-select-option>
              <a-select-option value="delete">
                DELETE
              </a-select-option>
            </a-select>
          </a-form-item>
        </template>
      </a-form>
    </a-modal>
  </a-card>
</template>

<script>
import APIS from '@/api/url'
import { getMenus, addMenu, updateMenu, deleteMenu } from '@/api/menu'

const columns = [
  {
    title: '标题',
    dataIndex: 'title',
    fixed: 'left',
    width: 180
  },
  {
    title: '唯一标识',
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
      ModalText: '新增菜单',
      visible: false,
      confirmLoading: false,
      form: this.$form.createForm(this, { name: 'coordinated' }),
      loadingTable: false,
      isMenu: null,
      menuApi: APIS.BaseUrl + APIS.menuApi.menu,
      menuTree: [],
      columns,
      menus: [],
      data: [],
      expandedRowIds: [],
      queryParam: { query: '' },
      treeOpen: false,
      editRecord: null
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
    listMenu (dataList) {
      dataList.forEach(element => {
          if (element.type !== 1) {
            return
          }
          this.menus.push({ id: element.id, title: element.title })
          if (element.children && element.children.length > 0) {
            this.listMenu(element.children)
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
      this.ModalText = '新增菜单'
      this.editRecord = null
      this.menus = []
      this.visible = true
      this.listMenu(this.menuTree)

        this.initForm()
    },
    updateMenu (record) {
      this.ModalText = '编辑菜单'
      this.menus = []
      this.listMenu(this.menuTree)
      this.isMenu = record.type
      this.initForm()
      this.visible = true
      this.editRecord = record
      this.$nextTick(() => {
          this.form.setFieldsValue({
            parentId: this.editRecord.parentId,
            name: this.editRecord.name.split(' : ')[0],
            title: this.editRecord.title,
            type: this.editRecord.type
          })
          })
      this.handleChange(record.type)
        },
    openClose () {
      if (this.treeOpen) {
        this.expandedRowIds = []
      } else {
        this.expandedRow(this.data)
      }
      this.treeOpen = !this.treeOpen
    },
    async handleOk (e) {
      this.form.validateFields(async (err, values) => {
        if (err) {
          return
        }
        values.show = values.show === 1
        this.confirmLoading = true
        let data = null
        try {
        if (this.editRecord) {
          if (this.editRecord.type === 1) {
            values.id = this.editRecord.id
          } else {
            values.id = parseInt(this.editRecord.id.replaceAll('p', ''))
          }
          data = await updateMenu(values)
        } else {
          data = await addMenu(values)
        }
      } catch (e) {
        return
      } finally {
        this.confirmLoading = false
      }
      if (data.code && data.code === 20001) {
          this.$message.success(this.editRecord ? '编辑成功' : '添加成功')
          this.visible = false
          this.searchMenu()
        }
        this.visible = false
      this.initForm()
      this.editRecord = null
      this.isMenu = null
      })
    },
    handleCancel (e) {
      this.visible = false
      this.initForm()
      this.editRecord = null
      this.isMenu = null
    },
    filterOption (input, option) {
      return (
        option.componentOptions.children[0].text.toLowerCase().indexOf(input.toLowerCase()) >= 0
      )
    },
    handleChange (value) {
      this.isMenu = value === 1
      // 菜单
      if (this.isMenu === true) {
        if (this.editRecord) {
          this.$nextTick(() => {
          this.form.setFieldsValue({
            component: this.editRecord.component,
            show: this.editRecord.show ? 1 : 0,
            redirect: this.editRecord.redirect,
            sort: this.editRecord.sort,
            icon: this.editRecord.icon,
            target: this.editRecord.target,
            path: this.editRecord.path
          })
          })
        } else {
          this.$nextTick(() => {
          this.form.setFieldsValue({
            component: 'RouteView',
            show: 1
          })
      })
        }
      } else { // 权限
        if (this.editRecord) { // 编辑
          this.$nextTick(() => {
          this.form.setFieldsValue({
            action: this.editRecord.action
          })
        })
        }
      }
    },
    deleteMenu (record) {
      this.$confirm({
        title: '是否确认要删除该菜单？',
        content: '删除之后将无法恢复',
        okText: '确认',
        okType: 'danger',
        cancelText: '取消',
        onOk: async () => {
          let data = {}
          try {
            const d = { type: record.type }
            if (record.type === 2) {
              d.id = parseInt(record.id.replaceAll('p', ''))
            } else {
              d.id = record.id
            }
            data = await deleteMenu(d)
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
            this.searchMenu()
          }
        }
      })
    },
    initForm () {
      this.form.resetFields()
      this.$nextTick(() => {
      this.form.setFieldsValue({
        // component: 'RouteView'
      })
    })
  }
  },
  mounted () {
    this.searchMenu(false)
  }

}
</script>

<style>

</style>
