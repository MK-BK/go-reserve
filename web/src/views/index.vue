<template>
    <el-container class="body">
        <el-header class="header">
            <div class="title">预约系统</div>
            <div class="view">
                <el-badge :value="12" class="item">
                    <el-button class="share-button" icon="el-icon-share" type="primary"></el-button>
                </el-badge>
                <el-avatar>admin</el-avatar>
                <el-dropdown trigger="click">
                    <span class="el-dropdown-link">
                        <span>admin</span>
                        <el-icon class="el-icon--right"></el-icon>
                    </span>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item @click="setting">设置</el-dropdown-item>
                            <el-dropdown-item @click="logout">退出</el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </div>
        </el-header>
        <el-container class="menu">
            <el-menu active-text-color="blue" background-color="#ffff" default-active="1" text-color="#ffff">
                <el-menu-item index="1">
                    <router-link to="/users">用户列表</router-link>
                </el-menu-item>
                <el-menu-item index="2">
                    <router-link to="/shops">商铺列表</router-link>
                </el-menu-item>
                <el-menu-item index="3">
                    <router-link to="/requests">商铺请求</router-link>
                </el-menu-item>
                <el-menu-item index="4">
                    <router-link to="/job">我的订单</router-link>
                </el-menu-item>
                <el-menu-item index="5">
                    <router-link to="/commodity">商品列表</router-link>
                </el-menu-item>
                <el-menu-item index="6">
                    <router-link to="/audit_logs">操作日志</router-link>
                </el-menu-item>
                <el-menu-item index="7">
                    <router-link to="/setting">设置</router-link>
                </el-menu-item>
            </el-menu>
            <div class="router-view">
                <router-view/>
            </div>
        </el-container>
    </el-container>
</template>

<script>
export default {
    data() {
        return {
            menus: {},
        }
    },
    mounted() {
        if (this.$store.state.user.Auth) {

        }
    },
    methods: {
        logout() {
            this.$router.push('/')
        },
        setting() {
            this.$router.push('/setting')
        },
        getMenues() {
            debugger
            if (this.$store.user.Auth) {

            }
        }
    }
}
</script>

<style lang="less" scoped>
.body {
    border: 1px solid;
    height: 100%;
}

.header {
    display: flex;
    align-items: center;
    border-bottom: 1px solid;
    justify-content: space-between;

    > .title {
        font-size: 20px;
    }

    > .view {
        display: flex;
        align-items: center;

        > div {
            margin-left: 20px;
        }

        > span {
            margin-left: 20px;
        }
    }
}

.menu {
    .is-active {
        background-color: rgb(17, 226, 233);
    }

    a {
        font-size: 16px;
        color: rgb(13, 79, 165);
        text-decoration: none;
        margin-top: 0;
    }
}

.router-view {
    width: 100%;
    padding: 10px 20px;
}
</style>