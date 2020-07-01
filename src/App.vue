<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <v-toolbar-title>MDI-WIKI</v-toolbar-title>
      </div>

      <v-spacer></v-spacer>
      <v-toolbar-title>Welcome</v-toolbar-title>
    </v-app-bar>

    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
// import Home from '@/components/Home';

export default {
  name: 'App',

  components: {
    //Home,
  },

  data: () => ({
    //
  }),
};
</script>
