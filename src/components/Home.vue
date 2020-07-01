<template>
    <v-card id="home"
            class="mx-auto"
            color="white"
            width="500px">
        <v-list>
            <v-list-item-group v-model="all_info.name" mandatory color="indigo">
                <v-list-item
                        v-for="article in all_info"
                        :key="article.name"
                >
                    <v-list-item-content>
                        <v-btn text color="black accent-4" @click="jumpTo(article.name)">{{article.name}}</v-btn>
                    </v-list-item-content>
                </v-list-item>
            </v-list-item-group>
        </v-list>

    </v-card>
</template>

<script>
    export default {
        name: "Home",
        data(){
            return {
                all_info: [],
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
