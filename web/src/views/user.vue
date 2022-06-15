<template>
    <div>
        <div class="view-header">用户列表</div>
        <el-table :data="users" style="width: 100%">
            <el-table-column prop="name" label="名称" />
            <el-table-column prop="Email" label="邮件" />
            <el-table-column label="用户类型">
                <template #default="scope">
                    <div>{{getAuth(scope.row.Auth)}}</div>
                </template>
            </el-table-column>
            <el-table-column label="状态">
                <template #default="scope">
                    <div>{{getStatus(scope.row.Enable)}}</div>
                </template>
            </el-table-column>
            <el-table-column prop="CreatedAt" label="创建时间" />
            <el-table-column fixed="right" label="操作" width="100">
                <template #default="scope">
                    <el-dropdown trigger="click">
                        <span class="el-dropdown-link">
                            <el-icon class="el-icon--right"><el-icon class="el-icon-more"></el-icon></el-icon>
                        </span>
                        <el-button @click="editUser(scope.row)">编辑</el-button>
                        <!-- <template #dropdown"> -->
                            <!-- <el-dropdown-menu>
                                <el-dropdown-item @click="editUser(scope.row)">编辑</el-dropdown-item>
                                <el-dropdown-item>删除</el-dropdown-item>
                            </el-dropdown-menu> -->
                        <!-- </template> -->
                    </el-dropdown>
                </template>
            </el-table-column>
        </el-table>
        <router-view></router-view>
    </div>
</template>

<script>
export default {
    data: function() {
        return {
            users: []
        }
    },
    
    mounted: async function() {
        this.users = await this.$store.dispatch('listUsers')
    },

    methods: {  
        getAuth: function(auth) {
            switch (auth) {
                case 0: 
                    return "管理员"
                case 1: 
                    return "商铺管理员"
                default: 
                    return "普通用户"
            }
        },
        getStatus(enable) {
            if (enable) {
                return '禁用'
            }
            return '正常'
        },
        editUser(user) {
            this.$router.push('/users/1/edit')
        }
    }
}
</script>