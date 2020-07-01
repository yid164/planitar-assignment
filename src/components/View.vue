<template>
    <div>
        Hello
    </div>
</template>

<script>
    export default {
        name: "View",
        data(){
            return{
                articleName: this.$route.params.name
            }
        }
    }
</script>

<style scoped>

</style>
