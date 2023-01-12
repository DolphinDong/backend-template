<template>
  <a-card :bordered="false">
    <RoleTree></RoleTree>

  </a-card>
</template>

<script>
import RoleTree from '@/components/RoleTree'
export default {
  name: 'Test2',
  components: {
    RoleTree
  }

}
</script>

<style>

</style>
