<template>
  <div>
    <a-radio-group v-model="size">
      <a-radio value="small">Small</a-radio>
      <a-radio value="middle">Middle</a-radio>
      <a-radio value="large">Large</a-radio>
    </a-radio-group>
    <br />
    <br />
    <a-space :size="size">
      <a-button type="primary">Primary</a-button>
      <a-button>Default</a-button>
      <a-button type="dashed">Dashed</a-button>
      <a-button type="link">Link</a-button>
    </a-space>
    <a-button v-if="$auth('/test/test1.delete')">$auth</a-button>
    <a-button v-action:delete>v-action</a-button>
  </div>
</template>
  <script>
  export default {
    name: 'Test1',
    data () {
      return {
        size: 'small'
      }
    }
  }
  </script>
