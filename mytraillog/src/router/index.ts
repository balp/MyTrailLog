import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/logs',
    name: 'logs',
    component: () => import(/* webpackChunkName: "logs" */ '../views/LogsView.vue')
  },
  {
    path: '/register',
    name: 'userregister',
    component: () => import(/* webpackChunkName: "UserRegister" */ '../views/UserRegister.vue')
  },
  {
    path: '/signin',
    name: 'signin',
    component: () => import(/* webpackChunkName: "SignIn" */ '../views/UserSignIn.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
