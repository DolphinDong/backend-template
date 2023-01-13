<template>
  <div>
    <a-button @click="openClose">全部{{ isOpen?"收起":"展开" }}</a-button>
    <a-button style="margin:0 8px;" type="primary" @click="selectAll">全部勾选</a-button>
    <a-button @click="cancelSelect" type="primary">全部取消</a-button>
    <a-tree
      v-model="checkedKeys"
      checkable
      :auto-expand-parent="autoExpandParent"
      :tree-data="treeData"
      :defaultExpandAll="true"
      :checkStrictly="true"
      :replaceFields="replaceFields"
      @expand="onExpand"
      @select="onSelect"
      :expanded-keys="expandedKeys"
    />
  </div>
  <!-- :expanded-keys="defaultExpandedKeys" -->
</template>
  <script>
  import APIS from '@/api/url'
import { getMenus, addMenu, updateMenu, deleteMenu } from '@/api/menu'
  export default {
    props: ['defaultCheck'],
    data () {
      return {
        autoExpandParent: true,
        checkedKeys: this.defaultCheck,
        treeData: [],
        replaceFields: { key: 'id' },
        expandedKeys: [],
        isOpen: true
      }
    },
    watch: {
      checkedKeys (val) {
        // console.log('onCheck', val)
      }
    },
    methods: {
      onExpand (expandedKeys) {
        // console.log('onExpand', expandedKeys)
        // if not set autoExpandParent to false, if children expanded, parent can not collapse.
        // or, you can remove all expanded children keys.
        this.expandedKeys = expandedKeys
        this.autoExpandParent = false
      },
      onCheck (checkedKeys) {
        // console.log('onCheck', checkedKeys)
        this.checkedKeys = checkedKeys
      },
      onSelect (selectedKeys, info) {
        // console.log('onSelect', info)
        this.selectedKeys = selectedKeys
      },
    async queryMenus () {
        this.selectedKeys = []
        try {
            this.loadingTable = true
            const res = await getMenus({})
           const data = res.data
            if (data) {
                this.changePermissionId(data.menu_tree)
                this.listMenu(data.menu_tree)
                this.treeData = data.menu_tree
            }
        } catch (e) {
                return
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
    listMenu (dataList) {
      dataList.forEach(element => {
          if (element.type !== 1) {
            return
          }
          this.expandedKeys.push(element.id)
          if (element.children && element.children.length > 0) {
            this.listMenu(element.children)
          }
      })
    },
    cancelSelect () {
      this.checkedKeys = []
    },
    selectAll () {
      this.checkedKeys = []
      this.selectAllMenu(this.treeData)
    },
    selectAllMenu (dataList) {
      dataList.forEach(element => {
          this.checkedKeys.push(element.id)
          if (element.children && element.children.length > 0) {
            this.selectAllMenu(element.children)
          }
      })
    },
    openClose () {
      this.isOpen = !this.isOpen
      if (this.isOpen) {
        this.listMenu(this.treeData)
      } else {
        this.expandedKeys = []
      }
    }
    },
    mounted () {
        this.queryMenus()
        // this.$on('sendChecks', this.checkedKeys)
    }
  }
  </script>
