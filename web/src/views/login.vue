<template>
    <el-form :model="form" label-position="top" label-width="120px" class="form">
        <el-form-item label="用户名">
            <el-input v-model="this.user.name" />
        </el-form-item>
        
        <el-form-item label="密码">
            <el-input v-model="this.user.password" />
        </el-form-item>

        <template v-if="action!=='login'">
            <el-form-item label="邮箱">
                <el-input v-model="this.user.email" />
            </el-form-item>
            <el-form-item>
                <el-radio label="1" size="large" value="1" v-model="this.user.auth">店铺管理员</el-radio>
                <el-radio label="2" size="large" value="0" v-model="this.user.auth">普通用户</el-radio>
            </el-form-item>
        </template>

        <el-form-item v-if="action==='login'" class="form-item">
            <el-button type="primary" @click="login">登录</el-button>
            <el-button @click="cancel">取消</el-button>
            <el-link type="primary" @click="action='register'">未注册账号，进行注册</el-link>
        </el-form-item>

        <el-form-item v-if="action!=='login'" class="form-item">
            <el-button type="primary" @click="register">注册</el-button>
            <el-button @click="cancel">取消</el-button>
            <el-link type="primary" @click="action='login'">已有账号，进行登录</el-link>
        </el-form-item>
    </el-form>
</template>

<script>
export default {
    data: function() {
        return {
            user: {
                name: '',
                password: '',
                email: '',
                auth: "1",
            },
            action: 'login'
        }
    },
    
    methods: {
        login: async function() {
            let that = this
            this.$store.dispatch('login', {
                name: this.user.name,
                password: this.user.password
            }).then(function(){
                that.$router.push("/index")
            }, function(err) {
                alert('登录失败，请检查并重新登录')
            })
        },

        register: async function() {
            this.user.auth = parseInt(this.user.auth)
            this.$store.dispatch('register', this.user).then(function(resp) {
                alert('注册成功，请进行登录')
            }, function(err) {
                alert('注册失败，请检查并重新注册')
            })
        },
        
        cancel() {
            this.action = 'login'
        }
    }
}
</script>


<style lang="less" scoped>
.form {
    width: 40%;
    height: 100%;
    margin: auto;
}

.form-item {
    a  {
        margin-left: 10px;
        margin-top: 0;
        font-size: 12px;
    }
}
</style>