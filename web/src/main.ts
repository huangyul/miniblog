import { createApp } from 'vue'
import App from './App.vue'
import installRouter from '@/router'
import installStore from '@/store'

import 'virtual:uno.css'
import '@unocss/reset/normalize.css'

const app = createApp(App)

installRouter(app)
installStore(app)

app.mount('#app')
