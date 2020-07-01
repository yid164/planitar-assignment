<template>
    <div id="home">
        <h1>Hello</h1>
        <div>{{all_info}}</div>
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
                const url = 'articles/'
                this.axios.get(url, {
                    headers: {
                        'Content-Type': '*',
                    }
                })
                    .then((result)=>{
                        this.all_info = result.data
                        console.log(result.data)
                    })
            }
        },

        created() {
            this.fetchInfo()
        }


    }
</script>

<style scoped>

</style>
