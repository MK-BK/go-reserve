<template>
    <div>
        <div class="view-header">用户设置</div>
        <template v-if="triggerAvatar">
            <el-upload class="upload-demo" drag :action=domain multiple ref="upload" :http-request=uploadFile>
                <i class="el-icon-upload"></i>
                <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
            </el-upload>
        </template>
       <el-form :label-position="labelPosition" label-width="80px" :model="formLabelAlign">
            <el-form-item label="头像">
                <el-avatar v-if="!triggerAvatar" shape="square" :size="100" :fit="fit" :src="url"></el-avatar>
                <el-button @click="triggerAvatar=!triggerAvatar">更新头像</el-button>
            </el-form-item>
            <el-form-item label="显示名称">
                <el-input v-model="name"></el-input>
            </el-form-item>
            <el-form-item label="邮箱设置">
                <el-input v-model="email"></el-input>
            </el-form-item>
        </el-form>
        <div class="footer">
            <el-button type="primary" @click="update">更新</el-button>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            user: {
                name: '',
                email: ''
            },
            triggerAvatar: false,
            upload: {},
        }
    },
    methods: {
        async uploadFile() {
            await this.$store.dispatch('uploadAvatar', )
        },

        async update() {
            console.log('update')
        }
    }
}
</script>

<style lang="less" scoped>
.footer {
    display: flex;
    justify-content: flex-end;
}
</style>