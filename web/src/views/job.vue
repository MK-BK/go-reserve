<template>
    <div>
        <div class="view-header">用户订单</div>
        <el-table :data="this.jobs" style="width: 100%">
            <el-table-column prop="ID" label="订单ID" />
            <el-table-column prop="userID" label="用户ID" />
            <el-table-column prop="Price" label="Price"/>
            <el-table-column prop="commodityID" label="商品ID" />
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
