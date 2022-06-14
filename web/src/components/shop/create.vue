<template>
    <div class="modal">
        <div class="modal-header">
            <div class="title">创建商铺</div>
        </div>
        <div class="modal-body">
            <el-form label-position="top" label-width="80px">
                <el-form-item label="名称">
                    <el-input v-model="shop.name"></el-input>
                </el-form-item>
                <el-form-item label="地址">
                    <el-input v-model="shop.address"></el-input>
                </el-form-item>
                    <el-form-item label="描述信息">
                    <el-input v-model="shop.description"></el-input>
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
            shop: {
                name: '',
                address: '',
                description: ''
            }
        }
    },
    methods: {
        async create() {
            this.$store.dispatch('createShop', this.shop)
        },
        cancel() {
            console.log('cancel')
        }
    }
}
</script>