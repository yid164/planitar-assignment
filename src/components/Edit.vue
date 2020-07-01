<template>
    <div>
        <v-col>
            <v-list-item-title>{{articleName}}</v-list-item-title>
            <div v-if="!readyPreview">
                <v-textarea v-model="currentContent"></v-textarea>
            </div>
            <div v-else>
                <div v-if="currentContent === ''">
                    <p>No article with this exact name found， Use Edit button in the header to add it</p>
                </div>
                <div v-else>
                    <html>{{currentContent}}</html>
                    <v-btn @click="closePreview()">Close</v-btn>
                </div>
            </div>
        </v-col>
        <v-row align="center">
            <v-btn @click="preview()">Preview</v-btn>
            <v-btn @click="save()">Save</v-btn>
            <v-btn @click="cancel()">Cancel</v-btn>
        </v-row>
    </div>
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
