<template>
    <el-table :data="this.commodities" style="width: 100%">
        <el-table-column prop="name" label="name" />
        <el-table-column prop="shop_id" label="shop_id"/>
        <el-table-column prop="description" label="description" />
    </el-table>
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
