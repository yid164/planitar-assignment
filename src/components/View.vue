<template>

    <v-container>
        <v-card
                class="mx-auto"
                outlined
        >
            <v-list-item three-line>
                <v-list-item-content>
                    <div class="overline mb-4">Wiki</div>
                        <v-list-item-title class="headline mb-1">{{articleName}}</v-list-item-title>
                        <v-btn @click="edit" x-small>Edit</v-btn>
                    <v-list-item-subtitle v-if="articleContent === '' ">
                        No article with this exact name found， Use Edit button in the header to add it
                    </v-list-item-subtitle>
                    <v-list-item-subtitle v-else>
                        <html>
                        {{articleContent}}
                        </html>
                    </v-list-item-subtitle>

                </v-list-item-content>

            </v-list-item>

            <v-card-actions>
                <v-btn @click="backToAll()">Back To All</v-btn>
            </v-card-actions>
        </v-card>
    </v-container>
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
