<template>
    <!-- <div class="modal">
        <div class="modal-header">
            <div class="title">更新用户</div>
        </div>
        <div class="modal-body">
            <el-form label-position="top" label-width="80px">
                <el-form-item label="名称">
                    <el-input v-model="user.name"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="user.password"></el-input>
                </el-form-item>
                    <el-form-item label="描述信息">
                    <el-input v-model="user.description"></el-input>
                </el-form-item>
            </el-form>
        </div>
        <div class="modal-footer">
            <el-button @click="create" type="primary">更新</el-button>
            <el-button @click="cancel">取消</el-button>
        </div>
    </div> -->

    <div class="dialog">
        <div class="modal-header">
            <div class="title">更新用户</div>
        </div>
        <div class="dialog-content">
            <el-form label-position="top" label-width="80px">
                <el-form-item label="名称">
                    <el-input v-model="user.name"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="user.password"></el-input>
                </el-form-item>
                    <el-form-item label="描述信息">
                    <el-input v-model="user.description"></el-input>
                </el-form-item>
            </el-form>
        </div>
        <div class="modal-footer">
            <el-button @click="create" type="primary">更新</el-button>
            <el-button @click="cancel">取消</el-button>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            user: {
                name: '',
                password: '',
                description: ''
            }
        }
    },
    methods: {
        async create() {
            this.$store.dispatch('createShop', this.shop)
        },
        cancel() {
            this.$router.back()
        }
    }
}
</script>

<style lang="less" scoped>
.dialog {
    position: relative;
    height: 100%;
    border: 1px solid red;
    width: 50%;
    margin: auto;

    .dialog-content {
        width: 50%;
        position: fixed;
        box-sizing: border-box;
        padding: 20px;
        // width: 80%;
        min-height: 140px;
        left: 50%;
        top: 50%;
        transform: translate(-50%, -50%);
        border-radius: 5px;
        background: #fff;
        z-index: 50002;
        .title {
            font-size: 16px;
            font-weight: 600;
            line-height: 30px;
        }
        .text {
            font-size: 14px;
            line-height: 30px;
            color: #555;
        }
        .btn-group {
            display: flex;
            position: absolute;
            right: 0;
            bottom: 10px;
            .btn {
                padding: 10px 20px;
                font-size: 14px;
                &:last-child {
                    color: #76D49B;
                }
            }
        }
    }
    .mask {
        position: fixed;
        top: 0;
        left: 0;
        bottom: 0;
        right: 0;
        z-index: 50001;
        background: rgba(0,0,0,.5);
    }
}
</style>