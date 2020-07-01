<template>
    <div>
        <header>
            {{articleName}}
            <button @click="edit()">Edit</button>
        </header>
        <div>
            <div v-if="articleContent === '' ">
                <p>No article with this exact name found， Use Edit button in the header to add it</p>
            </div>
            <div v-else>
                <v-html>{{articleContent}}</v-html>
            </div>
        </div>

        <button @click="backToAll()">Back To All</button>

    </div>
</template>

<script>
    export default {
        name: "View",
        data(){
            return{
                articleName: this.$route.params.name,
                articleContent: ""
            }
        },
        methods:{
            fetchContent:function () {
                const url = '/articles/'+this.articleName
                this.axios.get(url, {
                    headers: {
                        'Content-Type': '*',
                    }
                })
                    .then((result)=>{
                        this.articleContent = result.data.content
                        console.log(result.data)
                    })
            },

            backToAll:function () {
                this.$router.push('/');
            },

            edit:function () {
                this.$router.push('/edit/'+this.articleName);
            }
        },
        mounted() {
            this.fetchContent();
        }
    }
</script>

<style scoped>

</style>
