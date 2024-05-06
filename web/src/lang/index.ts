import { createI18n } from 'vue-i18n'
import { App } from 'vue'
import zh_CN from './zh-CN.json'
import en_GB from './en.json'
import ja from './ja.json'
import zh_HK from './zh-HK.json'

const navlang = navigator.language
const locallang = (navlang == 'zh_CN' || navlang == 'en_GB') ? navlang : 'zh_CN'
const lang = localStorage.getItem('language') || locallang || 'zh_CN'
localStorage.setItem('language', lang)

const messages = {
  zh: {
    ...zh_CN,
  },
  en: {
    ...en_GB
  },
  ja: {
    ...ja
  },
  zh_HK: {
    ...zh_HK
  }
}

export type MessageKey = keyof typeof messages

const i18n = createI18n({
  locale: lang, // 设置当前语言类型
  legacy: false, // 如果要支持compositionAPI，此项必须设置为false;
  globalInjection: true, // 全局注册$t方法
  messages,
})

export const changeLanguage = (value: MessageKey) => {
  localStorage.setItem('language', value)
  i18n.global.locale.value = value
}

export default function (app: App) {
  app.use(i18n)
}
