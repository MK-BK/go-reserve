<template>
    <div>
         <div class="view-header">
            <a>商铺列表></a>
            <div>{{id}}</div>
            <el-button type="primary" @click="createCommodity">创建商品</el-button>
        </div>
        <div class="view-content">
            <el-avatar shape="square" :size="100" fit="fill" src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg" />
            <div class="content">
                <div>商铺名称: 测试</div>
                <div class="profile"> 
                    <span>商铺评分:</span>
                    <el-rate v-model="rate" allow-half />
                </div>
            </div>
        </div>
        <div class="view-profile">
            <el-tabs type="border-card" v-model="tab">
                <el-tab-pane name="view" label="商品列表">
                    <el-table :data="this.commodities" style="width: 100%">
                        <el-table-column prop="name" label="Name" />
                        <el-table-column prop="shopID" label="shopID"/>
                        <el-table-column prop="Description" label="Description" />
                        <el-table-column prop="CreatedAt" label="CreatedAt" />
                        <el-table-colum fixed="right" label="操作" width="100%">
                            <template #default="scope">
                                <el-button @click="down(scope.row)">商品下架</el-button>
                            </template>
                        </el-table-colum>
                    </el-table>
                </el-tab-pane>
                <el-tab-pane name="content" label="商铺评论">
                    <el-table :data="this.commodities" style="width: 100%">
                        <el-table-column prop="name" label="Name" />
                        <el-table-column prop="shopID" label="shopID"/>
                        <el-table-column prop="Description" label="Description" />
                        <el-table-column prop="CreatedAt" label="CreatedAt" />
                           <el-table-colum fixed="right" label="操作" width="100">
                            <template #default="scope">
                                <el-button @click="down(scope.row)">商品下架</el-button>
                            </template>
                        </el-table-colum>
                    </el-table>
                </el-tab-pane>
            </el-tabs>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            id: this.$router.path,
            commodities: [],
            rate: 3.5,
            tab: 'view'
        }
    },
    methods: {
        createCommodity() {
             this.$router.push('/commodity/create') 
        },
        change(value) {
            console.log(value)
        }
    },
    mounted: async function() {
        this.commodities = await this.$store.dispatch('listCommodity')
    }
}
</script>

<style lang="less" scoped>
.view-header {
    line-height: 50px;
    display: flex;
    align-items: center;
}

.view-content {
    padding: 20px 20px;
    border: 1px solid #00000059;
    display: flex;
    align-items: center;

    .content {
        margin-left: 20px;
        text-align: left;
        
        .profile {
            margin-top: 10px;
        }
    }
}

.view-profile {
    text-align: left;

    .profile-header {
        padding: 10px 0;
    }
}
</style>