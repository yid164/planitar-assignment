import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import router from './router'
import vuetify from './plugins/vuetify';


Vue.config.productionTip = false;

Vue.prototype.axios = axios;

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
