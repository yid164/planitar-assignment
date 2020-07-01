<template>
    <div>
        Hello
    </div>
</template>

<script>
    export default {
        name: "Edit",
        data() {
            return{
                articleName: this.$route.params.name
            }
        }
    }
</script>

<style scoped>
</style>
