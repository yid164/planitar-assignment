<template>
    <v-card id="edit"
            class="mx-auto"
            color="white"
            width="500px">
        <v-card-text>
            <p class="display-3 text--primary">
                {{articleName}}
            </p>
            <div v-if="!readyPreview">
                <v-textarea v-model="currentContent"></v-textarea>
            </div>
            <div v-else>
                <div v-if="currentContent === ''" class="text--primary">
                    <p>No article with this exact name found， Use Edit button in the header to add it</p>
                </div>
                <div v-else>
                    <span v-html="currentContent"></span>
                </div>
                <v-btn text color="red accent-4" @click="closePreview()">Close</v-btn>
            </div>

        </v-card-text>
        <v-card-actions>
            <v-btn text color="deep-purple accent-4" @click="preview()">Preview</v-btn>
            <v-btn text color="green accent-4" @click="save()">Save</v-btn>
            <v-btn text color="red accent-4" @click="cancel()">Cancel</v-btn>
        </v-card-actions>
    </v-card>

</template>

<script>
    export default {
        name: "Edit",
        data() {
            return{
                articleName: this.$route.params.name,
                currentContent: '',
                readyPreview: false
            }
        },
        methods:{
            save:function () {
                const url = '/articles/'+this.articleName;
                this.axios.put(url, {content: this.currentContent})
                    .then()
                .catch(err=>{
                    console.log(err)
                });
                this.$router.back();
            },

            cancel:function () {

                this.$router.back();

            },

            fetchContent:function () {
                const url = '/articles/'+this.articleName
                this.axios.get(url, {
                    headers: {
                        'Content-Type': '*',
                    }
                })
                    .then((result)=>{
                        this.currentContent = result.data.content
                        console.log(result.data)
                    })
            },

            preview: function () {
                this.readyPreview = true;
            },

            closePreview: function () {
                this.readyPreview = false
            },

        },
        mounted() {
            this.fetchContent();
        }
    }
</script>

<style scoped>
</style>
