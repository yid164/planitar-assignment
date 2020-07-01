<template>
    <div>
        <header>
            {{articleName}}
        </header>
        <form v-if="!readyPreview">
            <textarea v-model="currentContent"></textarea>
            <button @click="preview()">preview</button>
        </form>
        <div v-else>
            <div v-if="currentContent === '' ">
                <p>No article with this exact name found， Use Edit button in the header to add it</p>
            </div>
            <div v-else>
                <v-html>{{currentContent}}</v-html>
            </div>
            <button @click="closePreview()">Close</button>
        </div>

        <button @click="save()">Save</button>
        <button @click="cancel()">Cancel</button>
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
