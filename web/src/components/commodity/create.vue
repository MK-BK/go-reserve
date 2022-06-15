<template>
     <div class="dialog">
        <div class="modal-header">
            <div class="title">创建商品</div>
        </div>
        <div class="dialog-content">
            <el-form label-position="top" label-width="80px">
                <el-form-item label="名称">
                    <el-input v-model="name"></el-input>
                </el-form-item>
                    <el-form-item label="描述信息">
                    <el-input v-model="description"></el-input>
                </el-form-item>
            </el-form>
        </div>
        <div class="modal-footer">
            <el-button @click="create" type="primary">创建</el-button>
            <el-button @click="cancel">取消</el-button>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            name: '',
            description: ''
        }
    },
    methods: {
        create() {
            this.$store.dispatch('createCommodity', {
                name: this.name,
                description: this.description,
                shopID: 1,
            })
            this.cancel()
        }, 

        cancel() {
            this.$router.back()
        }
    }
}
</script>