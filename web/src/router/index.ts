import { createRouter, createWebHistory } from 'vue-router'
import routes from './routes'
import { App } from 'vue'

export const router = createRouter({
  history: createWebHistory(),
  routes: routes,
  scrollBehavior: () => ({ top: 0, left: 0 })
})

export default function (app: App) {
  app.use(router)
}
