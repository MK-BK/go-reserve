<template>
    <div>
         <div class="view-header">
            <div>商铺列表></div>
            <div>{{id}}</div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            id: this.$router.path
        }
    },
    mounted: async function() {
    }
}
</script>