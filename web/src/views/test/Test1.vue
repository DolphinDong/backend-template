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
    <a-button v-if="$auth('/api/test/test2.delete')" @click="authTest(1)">/api/test/test2.delete</a-button>
    <a-button v-if="$auth(authUrl+'.update')" @click="authTest(2)">err</a-button>
    <a-button v-if="$auth(authUrl+'.update')" @click="authTest(3)">warning</a-button>
    <a-button v-action:delete>v-action1</a-button>
    <a-button v-action:post>v-action2</a-button>
    <a-button v-action:get>v-action3</a-button>
    <a-button v-action:delete>v-action4</a-button>
    <h1>{{ title }}</h1>
  </div>
</template>
  <script>
  import request from '@/utils/request'
  import APIS from '@/api/url'
  export default {
    name: 'Test1',
    data () {
      return {
        size: 'small',
        authUrl: '/api/test/test1',
        title: ''
      }
    },
    methods: {
      async authTest (i) {
      const result = await request({
          url: '/test' + i,
          method: 'get'
        })
        if (result && result.code === 20001) {
        this.title = result.msg
      } else {
        this.title = ''
      }
      }

    },
    mounted () {
      console.log(APIS.BaseUrl, APIS.loginApi.UserInfo)
    }
  }
  </script>
