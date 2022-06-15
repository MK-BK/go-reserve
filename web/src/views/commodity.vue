<template>
    <div>
        <div class="view-header">商品列表</div>
        <el-table :data="this.commodities" style="width: 100%">
            <el-table-column prop="name" label="商品名称" />
            <el-table-column prop="shopID" label="商铺名称"/>
            <el-table-column prop="Description" label="商品描述信息" />
            <el-table-column prop="CreatedAt" label="创建时间" />
            <el-table-column label="操作">
                <template #default="scope">
                   <el-button @click="createJob(scope.row)">下单</el-button>
                   <el-button @click="updateCommdity(scope.row)">更新</el-button>
                </template>
            </el-table-column>
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
    },
    methods: {
        async createJob(commodity) {
            await this.$store.dispatch('createJob', {
                commodityID: commodity.ID + '',
                UserID: "1",
            })
        },
        async updateCommdity(commodity) {
            // await this.$store.dispatch('updateCommdity', {
            //     commodityID: commodity.ID + '',
            //     UserID: "1",
            // })
        }
    }
}
</script>
