<template>
    <div>
        <div class="view-header">日志列表</div>
        <el-table :data="this.auditLogs" style="width:100%" max-height="100%">
            <el-table-column prop="UserID" label="UserID" />
            <el-table-column prop="Action" label="Action" />
            <el-table-column prop="Status" label="status" />
            <el-table-column prop="CreatedAt" label="CreatedAt" />
            <el-table-column :prop="getDescription()" label="Description" />
        </el-table>
    </div>
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
    },
    methods: {
        getDescription(log) {
            return "description"
        }
    }
}
</script>

