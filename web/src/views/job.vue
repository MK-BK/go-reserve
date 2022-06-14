<template>
    <div>
        <div class="view-header">用户订单</div>
        <el-table :data="this.jobs" style="width: 100%">
            <el-table-column prop="userID" label="userID" />
            <el-table-column prop="Price" label="Price"/>
            <el-table-column prop="commodityID" label="commodityID" />
            <el-table-column prop="Status" label="Status" />
            <el-table-column prop="CreatedAt" label="CreatedAt" />
        </el-table>
    </div>
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
