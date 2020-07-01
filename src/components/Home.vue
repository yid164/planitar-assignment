<template>
    <div id="home">
        <div v-for="article in all_info" :key="article.name">
            <button @click="jumpTo(article.name)">
                {{article.name}}
            </button>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Home",
        data(){
            return {
                all_info: []
            }
        },
        methods:{
            fetchInfo:function () {
                const url = '/articles/'
                this.axios.get(url, {
                    headers: {
                        'Content-Type': '*',
                    }
                })
                    .then((result)=>{
                        this.all_info = result.data
                        console.log(result.data)
                    })
            },
            jumpTo:function (name) {
                this.$router.push('/'+name);
            }
        },

        created() {
            this.fetchInfo()
        }


    }
</script>

<style scoped>

</style>
