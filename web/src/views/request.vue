<template>
    <el-table :data="this.requests" style="width: 100%">
        <el-table-column prop="name" label="name" />
        <el-table-column prop="email" label="email"/>
        <el-table-column prop="status" label="status" />
    </el-table>
</template>

<script>
export default {
    name: 'Request',
    data: function() {
        return {
            requests: []
        }
    },
    mounted: async function() {
        this.requests = await this.$store.dispatch('listRequests')
    }
}
</script>
