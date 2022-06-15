<template>
    <div>
        <div class="view-header">
            <div>商铺列表</div>
            <el-button type="primary" @click="createShop">创建商铺</el-button>
        </div>
        <el-table :data="this.shops" style="width: 100%">
            <el-table-column label="名称">
                <template #default="scope">
                     <router-link to="/shops/1">{{scope.row.name}}</router-link>
                </template>
            </el-table-column>
            <el-table-column prop="address" label="地址" />
            <el-table-column prop="description" label="描述信息" />
        </el-table>
    </div>
</template>

<script>
export default {
    data: function() {
        return {
            shops: []
        }
    },
    mounted: async function() {
        this.shops = await this.$store.dispatch('listShop')
    },
    methods: {
        createShop() {
           this.$router.push('/shop/create') 
        }
    }
}
</script>