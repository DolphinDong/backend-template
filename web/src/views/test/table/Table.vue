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
  </div>
</template>
  <script>
  export default {
    data () {
      return {
        size: 'small'
      }
    }
  }
  </script>
