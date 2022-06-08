<template>
    <el-table :data="this.auditLogs" style="width: 100%">
        <el-table-column prop="user_id" label="user_id" />
        <el-table-column prop="action" label="action" />
        <el-table-column prop="created_at" label="created_at" />
        <el-table-column prop="status" label="status" />
    </el-table>
</template>

<script>
export default {
    data: function() {
        return {
            auditLogs: []
        }
    },
    mounted: async function() {
        this.auditLogs = await this.$store.dispatch('listAuditLogs')
    }
}
</script>
