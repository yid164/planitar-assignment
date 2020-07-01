import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from "@/components/Home";
import Edit from "@/components/Edit";
import View from "@/components/View";

Vue.use(VueRouter)

  const routes = [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/edit/:name',
      name: Edit,
      component: Edit
    },
    {
      path: '/:name',
      name: 'View',
      component: View
    }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
