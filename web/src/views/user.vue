<template>
    <el-table :data="users" style="width: 100%">
        <el-table-column prop="Auth" label="auth" />
        <el-table-column prop="name" label="Name" />
        <el-table-column prop="Email" label="Email" />
        <el-table-column prop="CreatedAt" label="CreatedAt" />
    </el-table>
</template>

<script>
export default {
  name: 'user',
  data: function() {
      return {
          users: []
      }
  },
  mounted: async function() {
      this.users = await this.$store.dispatch('listusers')
  }
}
</script>
