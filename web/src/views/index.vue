<template>
    <el-container>
        <el-header>
            <el-badge :value="12" class="item">
                <el-button>comments</el-button>
            </el-badge>
            <el-avatar>admin</el-avatar>
            <el-dropdown trigger="click">
                <span class="el-dropdown-link">
                    <span>admin</span>
                    <el-icon class="el-icon--right"><arrow-down /></el-icon>
                </span>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item>设置</el-dropdown-item>
                        <el-dropdown-item>退出</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </el-header>
        <el-container>
            <el-menu active-text-color="#ffd04b"  background-color="##e5e6eb" default-active="1" text-color="#fff">
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
                    <router-link to="/audit_logs">操作日志</router-link>
                </el-menu-item>
                <el-menu-item index="5">
                    <router-link to="/setting">设置</router-link>
                </el-menu-item>
                <el-menu-item index="6">
                    <router-link to="/job">用户订单</router-link>
                </el-menu-item>
                <el-menu-item index="7">
                    <router-link to="/commodity">商品列表</router-link>
                </el-menu-item>
            </el-menu>
            <router-view/>
        </el-container>
    </el-container>
</template>
