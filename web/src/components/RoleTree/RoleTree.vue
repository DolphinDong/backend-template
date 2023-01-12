<template>
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
  <!-- :expanded-keys="defaultExpandedKeys" -->
</template>
  <script>
  import APIS from '@/api/url'
import { getMenus, addMenu, updateMenu, deleteMenu } from '@/api/menu'
  export default {
    data () {
      return {
        autoExpandParent: true,
        checkedKeys: [3, 5],
        treeData: [],
        replaceFields: { key: 'id' },
        expandedKeys: []
      }
    },
    watch: {
      checkedKeys (val) {
        console.log('onCheck', val)
      }
    },
    methods: {
      onExpand (expandedKeys) {
        console.log('onExpand', expandedKeys)
        // if not set autoExpandParent to false, if children expanded, parent can not collapse.
        // or, you can remove all expanded children keys.
        this.expandedKeys = expandedKeys
        this.autoExpandParent = false
      },
      onCheck (checkedKeys) {
        console.log('onCheck', checkedKeys)
        this.checkedKeys = checkedKeys
      },
      onSelect (selectedKeys, info) {
        console.log('onSelect', info)
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
    }
    },
    mounted () {
        this.queryMenus()
    }
  }
  </script>
