<template>
    <el-table :data="this.jobs" style="width: 100%">
        <el-table-column prop="user_id" label="user_id" />
        <el-table-column prop="product_id" label="product_id"/>
        <el-table-column prop="commodity_id" label="commodity_id" />
        <el-table-column prop="price" label="price" />
        <el-table-column prop="status" label="status" />
    </el-table>
</template>

<script>
export default {
    data: function() {
        return {
            jobs: []
        }
    },
    mounted: async function() {
        this.jobs = await this.$store.dispatch('listJob')
    }
}
</script>
