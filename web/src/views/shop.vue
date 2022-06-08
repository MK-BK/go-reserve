<template>
    <el-table :data="this.shops" style="width: 100%">
        <el-table-column prop="name" label="name" />
        <el-table-column prop="address" label="address" />
    </el-table>
</template>

<script>
export default {
    data: function() {
        return {
            shops: []
        }
    },
    mounted: async function() {
        this.shops = await this.$store.dispatch('listJob')
    }
}
</script>
