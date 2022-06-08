<template>
    <div>
        <div>用户列表</div>
        <el-table :data="users" style="width: 100%">
            <el-table-column prop="Auth" label="auth" />
            <el-table-column prop="name" label="Name" />
            <el-table-column prop="Email" label="Email" />
            <el-table-column prop="CreatedAt" label="CreatedAt" />
        </el-table>
    </div>
</template>

<script>
import App from '../App.vue'
export default {
  components: { App },
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
