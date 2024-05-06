import { type RouteRecordRaw } from 'vue-router'
import Home from '@/views/home/index.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    component: Home
  },
  {
    path: '/user/:id',
    component: () => import('@/views/user/index.vue')
  }
]
export default routes
