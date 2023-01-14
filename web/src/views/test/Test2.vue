<template>
  <a-card :bordered="false">
    <a-button @click="getCheck">获取选中</a-button>
    <a-button @click="reflash">刷新列表</a-button>
    <RoleList ref="roleList"></RoleList>
  </a-card>
</template>

<script>
import RoleList from '@/components/RoleList'
import { getUserRole } from '@/api/user'
export default {
  name: 'Test2',
  data () {
    return {
      userId: 2
    }
  },
  components: {
    RoleList
  },
  methods: {
    sendChecks (checks) {
      console.log(checks)
    },
    getCheck () {
      console.log('1111111', this.$refs.roleList.selectedRowKeys)
    },
    async reflash () {
      this.userId = 'e672996e-ca24-4313-a72a-a20b48bf225a'
      const data = await getUserRole({ user_id: this.userId })
     this.$refs.roleList.selectedRowKeys = data.data
      this.$refs.roleList.queryRole()
    }
  }

}
</script>

<style>

</style>
