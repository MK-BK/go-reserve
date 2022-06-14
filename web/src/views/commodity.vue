<template>
    <div>
        <div class="view-header">商品列表</div>
        <el-table :data="this.commodities" style="width: 100%">
            <el-table-column prop="name" label="Name" />
            <el-table-column prop="shopID" label="shopID"/>
            <el-table-column prop="Description" label="Description" />
            <el-table-column prop="CreatedAt" label="CreatedAt" />
        </el-table>
    </div>
</template>

<script>
export default {
    data: function() {
        return {
            commodities: []
        }
    },
    mounted: async function() {
        this.commodities = await this.$store.dispatch('listCommodity')
    }
}
</script>
