import { createApp } from 'vue'
import App from './App.vue'
import installRouter from '@/router'
import installStore from '@/store'
import installI18n from '@/lang'

import 'virtual:uno.css'
import '@unocss/reset/normalize.css'

// antd-design-vue
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/reset.css'

const app = createApp(App)

installRouter(app)
installStore(app)
installI18n(app)

app.use(Antd)

app.mount('#app')
