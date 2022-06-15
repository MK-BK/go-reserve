<template>
    <div>
        <div class="view-header">商铺请求列表</div>
        <el-table :data="this.requests" style="width: 100%">
            <el-table-column prop="Name" label="名称" />
            <el-table-column prop="Email" label="邮箱"/>
            <el-table-column prop="Status" label="状态" />
            <el-table-column fixed="right" label="操作" width="100">
                <template #default="scope">
                    <template v-if="scope.row.Status==='new'">
                        <el-button @click="update(scope.row, 'approve')">允许</el-button>
                        <el-button @click="update(scope.row, 'deny')">拒绝</el-button>
                    </template>
                </template>
            </el-table-column>
        </el-table>
    </div>
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
    },
    methods: {
        async update(request, action) {
            console.log(action)
            let newRquest = {}
            Object.assign(newRquest, request)
            newRquest.Status = action
            await this.$store.dispatch('updateRquest', newRquest)
        },
    }
}
</script>