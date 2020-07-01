<template>
    <v-card id="view"
         class="mx-auto"
         color="white"
         width="500px">

        <v-card-text>
            <p class="display-3 text--primary">
                {{articleName}}
                <a @click="edit" class="text-button sm1">Edit</a>
            </p>
            <div v-if="articleContent === ''" class="text--primary">
                <p>No article with this exact name found， Use Edit button in the header to add it</p>
            </div>
            <div v-else>
                <span v-html="articleContent"></span>
            </div>
        </v-card-text>
        <v-card-actions>
            <v-btn text color="deep-purple accent-4" @click="backToAll">
                Back To All
            </v-btn>
        </v-card-actions>
    </v-card>
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
