<template>
    <v-container>
        <v-list>
            <v-list-item-group v-model="model" mandatory color="indigo">
                <v-list-item
                        v-for="article in all_info"
                        :key="article.name"
                >

                    <v-list-item-content center>
                        <v-list-item-title v-text="article.name" @click="jumpTo(article.name)"></v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list-item-group>
        </v-list>

    </v-container>
</template>

<script>
    export default {
        name: "Home",
        data(){
            return {
                all_info: [],
                model: 1,
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

        mounted() {
            this.fetchInfo()
        }


    }
</script>

<style scoped>

</style>
