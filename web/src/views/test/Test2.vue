<template>
  <a-card :bordered="false">
    <a-button @click="getCheck">获取选中</a-button>
    <RoleTree ref="roleTree" :defaultCheck="defaultCheck"></RoleTree>

  </a-card>
</template>

<script>
import RoleTree from '@/components/RoleTree'
export default {
  name: 'Test2',
  data () {
    return {
      defaultCheck: [3, 'p511', 'p37']
    }
  },
  components: {
    RoleTree
  },
  methods: {
    sendChecks (checks) {
      console.log(checks)
    },
    getCheck () {
      console.log('1111111', this.$refs.roleTree.checkedKeys)
    }
  }

}
</script>

<style>

</style>
