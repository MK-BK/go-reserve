<template>
    <div>
        <div class="view-header">商铺请求列表</div>
        <el-table :data="this.requests" style="width: 100%">
            <el-table-column prop="ID" label="ID" />
            <el-table-column prop="Name" label="Name" />
            <el-table-column prop="Email" label="Email"/>
            <el-table-column prop="Status" label="Status" />
            <el-table-column fixed="right" label="操作" width="100">
                <el-dropdown trigger="click">
                    <span class="el-dropdown-link">
                        <el-icon class="el-icon--right">
                            <el-icon class="el-icon-more"></el-icon>
                        </el-icon>
                    </span>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item @click="approve">允许</el-dropdown-item>
                            <el-dropdown-item @click="deny">拒绝</el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
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
        async approve(request) {
            console.log("approve")
        },

        async deny(request) {
            console.log("deny")
        }
    }
}
</script>
