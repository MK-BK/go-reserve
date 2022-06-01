<template>
    <el-form :model="form" label-position="top" label-width="120px">
        <el-form-item label="user name">
            <el-input v-model="this.user.name" />
        </el-form-item>
        <el-form-item label="password">
            <el-input v-model="this.user.password" />
        </el-form-item>
        <el-form-item label="email">
            <el-input v-model="this.user.email" />
        </el-form-item>
        <el-form-item>
            <el-radio label="1" size="large">普通用户</el-radio>
            <el-radio label="2" size="large">店铺管理员</el-radio>
        </el-form-item>
        <el-form-item>
            <el-button>登录</el-button>
            <el-button type="primary">注册</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
export default {
  data: function() {
      return {
          user: {
              name: '',
              password: '',
              email: '',
              auth: ''
          }
      }
  },
  method: {
      login() {
          this.$store.commit('', user)
      },

      register() {

      }
  }
}
</script>
